#!/usr/bin/make -f

PACKAGES_SIMTEST=$(shell go list ./... | grep '/simulation')
VERSION := $(shell echo $(shell git describe --tags) | sed 's/^v//')
COMMIT := $(shell git log -1 --format='%H')
LEDGER_ENABLED ?= true
SDK_PACK := $(shell go list -m github.com/cosmos/cosmos-sdk | sed  's/ /\\@/g')
BINDIR ?= $(GOPATH)/bin
SIMAPP = ./app

# Fetch from env
VERSION ?= $(shell echo $(shell git describe --tags) | sed 's/^v//')
COMMIT ?= $(shell git log -1 --format='%H')
OS ?= $(shell uname -s)
ROOT ?= $(shell git rev-parse --show-toplevel)

# for dockerized protobuf tools
DOCKER := $(shell which docker)
HTTPS_GIT := github.com/sonr-io/core.git

export GO111MODULE = on

# process build tags

build_tags = netgo
ifeq ($(LEDGER_ENABLED),true)
  ifeq ($(OS),Windows_NT)
    GCCEXE = $(shell where gcc.exe 2> NUL)
    ifeq ($(GCCEXE),)
      $(error gcc.exe not installed for ledger support, please install or set LEDGER_ENABLED=false)
    else
      build_tags += ledger
    endif
  else
    UNAME_S = $(shell uname -s)
    ifeq ($(UNAME_S),OpenBSD)
      $(warning OpenBSD detected, disabling ledger support (https://github.com/cosmos/cosmos-sdk/issues/1988))
    else
      GCC = $(shell command -v gcc 2> /dev/null)
      ifeq ($(GCC),)
        $(error gcc not installed for ledger support, please install or set LEDGER_ENABLED=false)
      else
        build_tags += ledger
      endif
    endif
  endif
endif

ifeq ($(WITH_CLEVELDB),yes)
  build_tags += gcc
endif
build_tags += $(BUILD_TAGS)
build_tags := $(strip $(build_tags))

whitespace :=
empty = $(whitespace) $(whitespace)
comma := ,
build_tags_comma_sep := $(subst $(empty),$(comma),$(build_tags))

# process linker flags
ldflags = -X github.com/cosmos/cosmos-sdk/version.Name=sonr \
		  -X github.com/cosmos/cosmos-sdk/version.AppName=snrd \
		  -X github.com/cosmos/cosmos-sdk/version.Version=$(VERSION) \
		  -X github.com/cosmos/cosmos-sdk/version.Commit=$(COMMIT) \
		  -X "github.com/cosmos/cosmos-sdk/version.BuildTags=$(build_tags_comma_sep)"

ifeq ($(WITH_CLEVELDB),yes)
  ldflags += -X github.com/cosmos/cosmos-sdk/types.DBBackend=cleveldb
endif
ifeq ($(LINK_STATICALLY),true)
	ldflags += -linkmode=external -extldflags "-Wl,-z,muldefs -static"
endif
ldflags += $(LDFLAGS)
ldflags := $(strip $(ldflags))

BUILD_FLAGS := -tags "$(build_tags_comma_sep)" -ldflags '$(ldflags)' -trimpath

# The below include contains the tools and runsim targets.
all: install lint test

build: go.sum
ifeq ($(OS),Windows_NT)
	$(error wasmd server not supported. Use "make build-windows-client" for client)
	exit 1
else
	@gum spin --show-error --title "[BUILD] Running go build..." -- go build -mod=readonly $(BUILD_FLAGS) -o bin/snrd ./cmd/snrd
	@gum log --level info --time kitchen "[BUILD] Completed go build successfully."
endif

build-windows-client: go.sum
	@gum spin --show-error --title "[BUILD] Building Windows client..." -- GOOS=windows GOARCH=amd64 go build -mod=readonly $(BUILD_FLAGS) -o bin/snrd.exe ./cmd/snrd
	@gum log --level info --time kitchen "[BUILD] Completed Windows client build successfully."

build-contract-tests-hooks:
ifeq ($(OS),Windows_NT)
	@gum spin --show-error --title "[BUILD] Building contract tests..." -- go build -mod=readonly $(BUILD_FLAGS) -o bin/contract_tests.exe ./cmd/contract_tests
	@gum log --level info --time kitchen "[BUILD] Completed contract tests build successfully."
else
	@gum spin --show-error --title "[BUILD] Building contract tests..." -- go build -mod=readonly $(BUILD_FLAGS) -o bin/contract_tests ./cmd/contract_tests
	@gum log --level info --time kitchen "[BUILD] Completed contract tests build successfully."
endif

install: go.sum
	@gum spin --show-error --title "[INSTALL] Installing snrd..." -- go install -mod=readonly $(BUILD_FLAGS) ./cmd/snrd
	@gum log --level info --time kitchen "[INSTALL] Successfully installed snrd."

########################################
### Tools & dependencies
########################################

go-mod-cache: go.sum
	@gum spin --show-error --title "[DEPS] Downloading go modules to local cache..." -- go mod download
	@gum log --level info --time kitchen "[DEPS] Successfully downloaded go modules."

go.sum: go.mod
	@gum spin --show-error --title "[DEPS] Verifying dependencies..." -- go mod verify
	@gum log --level info --time kitchen "[DEPS] Dependencies verified successfully."

draw-deps:
	@# requires brew install graphviz or apt-get install graphviz
	@gum spin --show-error --title "[DEPS] Installing goviz..." -- go install github.com/RobotsAndPencils/goviz@latest
	@gum spin --show-error --title "[DEPS] Generating dependency graph..." -- goviz -i . -d 2 | dot -Tpng -o dependency-graph.png
	@gum log --level info --time kitchen "[DEPS] Successfully generated dependency graph: dependency-graph.png"

clean:
	@gum spin --show-error --title "[CLEAN] Cleaning project..." -- rm -rf .aider* static .out hway.db snapcraft-local.yaml bin/ build
	@gum log --level info --time kitchen "[CLEAN] Project cleaned successfully."

distclean: clean
	@gum spin --show-error --title "[CLEAN] Deep cleaning project..." -- rm -rf vendor/
	@gum log --level info --time kitchen "[CLEAN] Project deep cleaned successfully."

########################################
### Testing
########################################

test: test-unit
test-all: test-race test-cover test-system

test-unit:
	@gum spin --show-error --title "[TEST] Running unit tests..." -- VERSION=$(VERSION) go test -mod=readonly -tags='ledger test_ledger_mock' ./...
	@gum log --level info --time kitchen "[TEST] Unit tests completed successfully."

test-race:
	@gum spin --show-error --title "[TEST] Running race tests..." -- VERSION=$(VERSION) go test -mod=readonly -race -tags='ledger test_ledger_mock' ./...
	@gum log --level info --time kitchen "[TEST] Race tests completed successfully."

test-cover:
	@gum spin --show-error --title "[TEST] Running coverage tests..." -- go test -mod=readonly -timeout 30m -race -coverprofile=coverage.txt -covermode=atomic -tags='ledger test_ledger_mock' ./...
	@gum log --level info --time kitchen "[TEST] Coverage tests completed successfully."

benchmark:
	@gum spin --show-error --title "[BENCH] Running benchmarks..." -- go test -mod=readonly -bench=. ./...
	@gum log --level info --time kitchen "[BENCH] Benchmarks completed successfully."

test-sim-import-export: runsim
	@gum log --level info --time kitchen "[SIM] Running application import/export simulation. This may take several minutes..."
	@gum spin --show-error --title "[SIM] Running import/export simulation..." -- $(BINDIR)/runsim -Jobs=4 -SimAppPkg=$(SIMAPP) -ExitOnFail 50 5 TestAppImportExport
	@gum log --level info --time kitchen "[SIM] Import/export simulation completed successfully."

test-sim-multi-seed-short: runsim
	@gum log --level info --time kitchen "[SIM] Running short multi-seed application simulation. This may take awhile!"
	@gum spin --show-error --title "[SIM] Running multi-seed simulation..." -- $(BINDIR)/runsim -Jobs=4 -SimAppPkg=$(SIMAPP) -ExitOnFail 50 5 TestFullAppSimulation
	@gum log --level info --time kitchen "[SIM] Multi-seed simulation completed successfully."

test-sim-deterministic: runsim
	@gum log --level info --time kitchen "[SIM] Running application deterministic simulation. This may take awhile!"
	@gum spin --show-error --title "[SIM] Running deterministic simulation..." -- $(BINDIR)/runsim -Jobs=4 -SimAppPkg=$(SIMAPP) -ExitOnFail 1 1 TestAppStateDeterminism
	@gum log --level info --time kitchen "[SIM] Deterministic simulation completed successfully."

test-system: install
	@gum spin --show-error --title "[TEST] Running system tests..." -- $(MAKE) -C tests/system/ test
	@gum log --level info --time kitchen "[TEST] System tests completed successfully."

###############################################################################
###                                Linting                                  ###
###############################################################################

format-tools:
	@gum spin --show-error --title "[LINT] Installing formatting tools..." -- go install mvdan.cc/gofumpt@v0.4.0 github.com/client9/misspell/cmd/misspell@v0.3.4 github.com/daixiang0/gci@v0.11.2
	@gum log --level info --time kitchen "[LINT] Formatting tools installed successfully."

lint: format-tools
	@gum spin --show-error --title "[LINT] Running golangci-lint..." -- golangci-lint run --tests=false
	@gum spin --show-error --title "[LINT] Checking code format..." -- find . -name '*.go' -type f -not -path "./vendor*" -not -path "./tests/system/vendor*" -not -path "*.git*" -not -path "*_test.go" | xargs gofumpt -d
	@gum log --level info --time kitchen "[LINT] Linting completed successfully."

format: format-tools
	@gum spin --show-error --title "[FORMAT] Formatting code with gofumpt..." -- find . -name '*.go' -type f -not -path "./vendor*" -not -path "./tests/system/vendor*" -not -path "*.git*" -not -path "./client/lcd/statik/statik.go" | xargs gofumpt -w
	@gum spin --show-error --title "[FORMAT] Fixing spelling errors..." -- find . -name '*.go' -type f -not -path "./vendor*" -not -path "./tests/system/vendor*" -not -path "*.git*" -not -path "./client/lcd/statik/statik.go" | xargs misspell -w
	@gum spin --show-error --title "[FORMAT] Running gci..." -- find . -name '*.go' -type f -not -path "./vendor*" -not -path "./tests/system/vendor*" -not -path "*.git*" -not -path "./client/lcd/statik/statik.go" | xargs gci write --skip-generated -s standard -s default -s "prefix(cosmossdk.io)" -s "prefix(github.com/cosmos/cosmos-sdk)" -s "prefix(github.com/CosmWasm/wasmd)" --custom-order
	@gum log --level info --time kitchen "[FORMAT] Code formatting completed successfully."

mod-tidy:
	@gum spin --show-error --title "[DEPS] Running go mod tidy..." -- go mod tidy
	@gum log --level info --time kitchen "[DEPS] go mod tidy completed successfully."

.PHONY: format-tools lint format mod-tidy


###############################################################################
###                                Protobuf                                 ###
###############################################################################
protoVer=0.15.1
protoImageName=ghcr.io/cosmos/proto-builder:$(protoVer)
protoImage=$(DOCKER) run --rm -v $(CURDIR):/workspace --workdir /workspace $(protoImageName)

proto-gen:
	@gum log --level info --time kitchen "[PROTO] Generating Protobuf files"
	@gum spin --show-error --title "[PROTO] Installing protoc-gen-go-cosmos-orm..." -- go install cosmossdk.io/orm/cmd/protoc-gen-go-cosmos-orm@latest
	@gum spin --show-error --title "[PROTO] Running protocgen.sh..." -- $(protoImage) sh ./scripts/protocgen.sh
	@gum spin --show-error --title "[PROTO] Running stub-gen..." -- spawn stub-gen
	@gum log --level info --time kitchen "[PROTO] Protobuf generation completed successfully."

proto-format:
	@gum log --level info --time kitchen "[PROTO] Formatting Protobuf files"
	@gum spin --show-error --title "[PROTO] Running clang-format..." -- $(protoImage) find ./ -name "*.proto" -exec clang-format -i {} \;
	@gum log --level info --time kitchen "[PROTO] Protobuf formatting completed successfully."

proto-lint:
	@gum spin --show-error --title "[PROTO] Running buf lint..." -- $(protoImage) buf lint --error-format=json
	@gum log --level info --time kitchen "[PROTO] Protobuf linting completed successfully."

proto-check-breaking:
	@gum spin --show-error --title "[PROTO] Checking for breaking changes..." -- $(protoImage) buf breaking --against $(HTTPS_GIT)#branch=master
	@gum log --level info --time kitchen "[PROTO] Protobuf breaking changes check completed successfully."

.PHONY: all install install-debug \
	go-mod-cache draw-deps clean build format \
	test test-all test-build test-cover test-unit test-race \
	test-sim-import-export build-windows-client \
	test-system

## --- Testnet Utilities ---
get-localic:
	@gum log --level info --time kitchen "[TESTNET] Installing local-interchain"
	@gum spin --show-error --title "[TESTNET] Cloning interchaintest..." -- git clone --branch v8.7.0 https://github.com/strangelove-ventures/interchaintest.git interchaintest-downloader
	@gum spin --show-error --title "[TESTNET] Installing local-ic..." -- cd interchaintest-downloader/local-interchain && make install
	@gum log --level info --time kitchen "[TESTNET] ✅ local-interchain installed $(shell which local-ic)"

is-localic-installed:
ifeq (,$(shell which local-ic))
	@gum spin --show-error --title "[TESTNET] local-ic not found, installing..." -- make get-localic
	@gum log --level info --time kitchen "[TESTNET] local-ic installed successfully."
endif

get-heighliner:
	@gum spin --show-error --title "[TESTNET] Cloning heighliner..." -- git clone https://github.com/strangelove-ventures/heighliner.git
	@gum spin --show-error --title "[TESTNET] Installing heighliner..." -- cd heighliner && go install
	@gum log --level info --time kitchen "[TESTNET] heighliner installed successfully."

local-image:
ifeq (,$(shell which heighliner))
	@gum log --level error --time kitchen "[TESTNET] 'heighliner' binary not found. Consider running 'make get-heighliner'"
else
	@gum spin --show-error --title "[TESTNET] Building local image..." -- heighliner build -c snrd --local -f chains.yaml
	@gum log --level info --time kitchen "[TESTNET] Local image built successfully."
endif

.PHONY: get-heighliner local-image is-localic-installed

###############################################################################
###                                     e2e                                 ###
###############################################################################

ictest-basic:
	@gum log --level info --time kitchen "[E2E] Running basic interchain tests"
	@gum spin --show-error --title "[E2E] Running TestBasicChain..." -- cd interchaintest && go test -race -v -run TestBasicChain .
	@gum log --level info --time kitchen "[E2E] Basic interchain tests completed successfully."

ictest-ibc:
	@gum log --level info --time kitchen "[E2E] Running IBC interchain tests"
	@gum spin --show-error --title "[E2E] Running TestIBC..." -- cd interchaintest && go test -race -v -run TestIBC .
	@gum log --level info --time kitchen "[E2E] IBC interchain tests completed successfully."

ictest-wasm:
	@gum log --level info --time kitchen "[E2E] Running cosmwasm interchain tests"
	@gum spin --show-error --title "[E2E] Running TestCosmWasmIntegration..." -- cd interchaintest && go test -race -v -run TestCosmWasmIntegration .
	@gum log --level info --time kitchen "[E2E] Cosmwasm interchain tests completed successfully."

ictest-packetforward:
	@gum log --level info --time kitchen "[E2E] Running packet forward middleware interchain tests"
	@gum spin --show-error --title "[E2E] Running TestPacketForwardMiddleware..." -- cd interchaintest && go test -race -v -run TestPacketForwardMiddleware .
	@gum log --level info --time kitchen "[E2E] Packet forward middleware tests completed successfully."

ictest-poa:
	@gum log --level info --time kitchen "[E2E] Running proof of authority interchain tests"
	@gum spin --show-error --title "[E2E] Running TestPOA..." -- cd interchaintest && go test -race -v -run TestPOA .
	@gum log --level info --time kitchen "[E2E] Proof of authority tests completed successfully."

ictest-tokenfactory:
	@gum log --level info --time kitchen "[E2E] Running token factory interchain tests"
	@gum spin --show-error --title "[E2E] Running TestTokenFactory..." -- cd interchaintest && go test -race -v -run TestTokenFactory .
	@gum log --level info --time kitchen "[E2E] Token factory tests completed successfully."

###############################################################################
###                                    testnet                              ###
###############################################################################

setup-ipfs:
	@gum spin --show-error --title "[IPFS] Setting up IPFS..." -- ./scripts/ipfs_config.sh
	@gum log --level info --time kitchen "[IPFS] IPFS setup completed successfully."

setup-testnet: mod-tidy is-localic-installed install local-image set-testnet-configs setup-testnet-keys
	@gum log --level info --time kitchen "[TESTNET] Testnet setup completed successfully."

# Run this before testnet keys are added
# chainid-1 is used in the testnet.json
set-testnet-configs:
	@gum spin --show-error --title "[TESTNET] Setting client chain-id..." -- snrd config set client chain-id sonr-testnet-1
	@gum spin --show-error --title "[TESTNET] Setting keyring-backend..." -- snrd config set client keyring-backend test
	@gum spin --show-error --title "[TESTNET] Setting output format..." -- snrd config set client output text
	@gum log --level info --time kitchen "[TESTNET] Client configs set successfully."

# import keys from testnet.json into test keyring
setup-testnet-keys:
	@gum spin --show-error --title "[TESTNET] Setting up acc0 key..." -- -`echo "decorate bright ozone fork gallery riot bus exhaust worth way bone indoor calm squirrel merry zero scheme cotton until shop any excess stage laundry" | snrd keys add acc0 --recover`
	@gum spin --show-error --title "[TESTNET] Setting up acc1 key..." -- -`echo "wealth flavor believe regret funny network recall kiss grape useless pepper cram hint member few certain unveil rather brick bargain curious require crowd raise" | snrd keys add acc1 --recover`
	@gum log --level info --time kitchen "[TESTNET] Testnet keys setup completed successfully."

# default testnet is with IBC
testnet: setup-testnet
	@gum spin --show-error --title "[TESTNET] Starting IBC testnet..." -- spawn local-ic start ibc-testnet
	@gum log --level info --time kitchen "[TESTNET] IBC testnet started successfully."

testnet-basic: setup-testnet
	@gum spin --show-error --title "[TESTNET] Starting basic testnet..." -- spawn local-ic start testnet
	@gum log --level info --time kitchen "[TESTNET] Basic testnet started successfully."

sh-testnet: mod-tidy
	@gum spin --show-error --title "[TESTNET] Starting shell testnet..." -- CHAIN_ID="sonr-testnet-1" BLOCK_TIME="1000ms" CLEAN=true sh scripts/test_node.sh
	@gum log --level info --time kitchen "[TESTNET] Shell testnet completed."

.PHONY: setup-testnet set-testnet-configs testnet testnet-basic sh-testnet dop-testnet

###############################################################################
###                                    extra utils                          ###
###############################################################################
status:
	@gum spin --show-error --title "[STATUS] Checking GitHub runs..." -- gh run ls -L 3
	@gum format -- "# Sonr ($OS-$VERSION)" "- ($(COMMIT)) $ROOT" "- $(RELEASE_DATE)"
	@gum log --level info --time kitchen "[STATUS] Status check completed."
	@sleep 3

push-docker:
	@gum spin --show-error --title "[DOCKER] Building Docker image..." -- docker build -t ghcr.io/onsonr/sonr:$(VERSION) .
	@gum spin --show-error --title "[DOCKER] Tagging latest image..." -- docker tag ghcr.io/onsonr/sonr:$(VERSION) ghcr.io/onsonr/sonr:latest
	@gum spin --show-error --title "[DOCKER] Pushing version image..." -- docker push ghcr.io/onsonr/sonr:$(VERSION)
	@gum spin --show-error --title "[DOCKER] Pushing latest image..." -- docker push ghcr.io/onsonr/sonr:latest
	@gum log --level info --time kitchen "[DOCKER] Docker images pushed successfully."

push-buf:
	@cd proto && gum spin --show-error --title "[BUF] Building protobufs..." -- buf build
	@cd proto && gum spin --show-error --title "[BUF] Pushing protobufs..." -- buf push
	@gum log --level info --time kitchen "[BUF] Protobufs pushed successfully."

release:
	@gum spin --show-error --title "[RELEASE] Running cz:bump..." -- devbox run cz:bump
	@gum log --level info --time kitchen "[RELEASE] Release completed successfully."

release-dry:
	@gum spin --show-error --title "[RELEASE] Running dry release..." -- devbox run release:dry
	@gum log --level info --time kitchen "[RELEASE] Dry release completed successfully."

deploy-deps:
	@gum log --level info --time kitchen "[DEPLOY] Installing deploy dependencies"
	@gum spin --show-error --title "[DEPLOY] Installing starship-ci..." -- npm install -g @starship-ci/cli
	@gum spin --show-error --title "[DEPLOY] Installing starship..." -- starship install
	@gum log --level info --time kitchen "[DEPLOY] Deploy dependencies installed successfully."

up:
	@gum log --level info --time kitchen "[DEPLOY] Starting deployment"
	@gum spin --show-error --title "[DEPLOY] Starting starship..." -- starship start --config .github/deploy/config.yml
	@gum log --level info --time kitchen "[DEPLOY] Deployment started successfully."

down:
	@gum log --level info --time kitchen "[DEPLOY] Stopping deployment"
	@gum spin --show-error --title "[DEPLOY] Stopping starship..." -- starship stop --config .github/deploy/config.yml
	@gum log --level info --time kitchen "[DEPLOY] Deployment stopped successfully."

###############################################################################
###                                     help                                ###
###############################################################################

help:
	@gum format -- "# Available targets:" \
		"  install             : Install the binary" \
		"  local-image         : Install the docker image" \
		"  proto-gen           : Generate code from proto files" \
		"  testnet             : Local devnet with IBC" \
		"  sh-testnet          : Shell local devnet" \
		"  ictest-basic        : Basic end-to-end test" \
		"  ictest-ibc          : IBC end-to-end test"

.PHONY: help
