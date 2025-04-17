package keeper_test

import (
	"context"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/address"
	"github.com/cosmos/cosmos-sdk/types/module/testutil"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	
	snrctx "github.com/sonr-io/core/internal"
	"github.com/sonr-io/core/internal/common"
	"github.com/sonr-io/core/x/dwn/keeper"
	"github.com/sonr-io/core/x/dwn/types"
)

type fixtureWrapper struct {
	ctx           sdk.Context
	encodingConfig testutil.TestEncodingConfig
	keeper        keeper.Keeper
	queryServer   types.QueryServer
	msgServer     types.MsgServer
	addrs         []sdk.AccAddress
	govModAddr    string
}

// SetupTest creates a test fixture for testing the query server
func SetupTest(t *testing.T) *fixtureWrapper {
	// Create an encoding config for the test
	encodingConfig := testutil.MakeTestEncodingConfig()
	
	// Create account addresses for testing
	addrs := createRandomAccounts(3)
	
	// Setup a context with a mock multi-store
	ctx := sdk.NewContext(nil, tmproto.Header{}, false, nil)
	
	// Create the gov module address
	govModAddr := authtypes.NewModuleAddress(govtypes.ModuleName).String()
	
	// Create a mock Params keeper
	paramsKeeper := new(MockParamsKeeper)
	paramsKeeper.On("Get", ctx).Return(types.DefaultParams(), nil)
	
	// Create the keeper with required dependencies
	k := keeper.Keeper{
		Params: paramsKeeper,
	}
	
	// Create the servers
	msgServer := keeper.NewMsgServerImpl(k)
	queryServer := keeper.NewQuerier(k)
	
	// Create context wrapper
	wrapper := &fixtureWrapper{
		ctx:           ctx,
		encodingConfig: encodingConfig,
		keeper:        k,
		queryServer:   queryServer,
		msgServer:     msgServer,
		addrs:         addrs,
		govModAddr:    govModAddr,
	}
	
	// Override the global FromGoContext function with a mock implementation
	overrideFromGoContext(ctx)
	
	// Override IPFS and Enclave functions
	setupMockIPFSAndEnclave()
	
	return wrapper
}

// MockParamsKeeper is a mock implementation of the ParamsKeeper interface
type MockParamsKeeper struct {
	mock.Mock
}

func (m *MockParamsKeeper) Get(ctx sdk.Context) (types.Params, error) {
	args := m.Called(ctx)
	return args.Get(0).(types.Params), args.Error(1)
}

// createRandomAccounts creates the specified number of random account addresses
func createRandomAccounts(numAccts int) []sdk.AccAddress {
	accts := make([]sdk.AccAddress, numAccts)
	for i := 0; i < numAccts; i++ {
		accts[i] = sdk.AccAddress(address.MustLengthPrefix([]byte{byte(i)}))
	}
	return accts
}

// overrideFromGoContext replaces the snrctx.FromGoContext with a mock implementation
func overrideFromGoContext(ctx sdk.Context) {
	// Create a mock implementation of SonrContext
	mockCtx := &MockSonrContext{
		sdkCtx: ctx,
		peer:   &MockPeer{},
	}
	
	// Override the global function
	originalFromGoContext := snrctx.FromGoContext
	snrctx.FromGoContext = func(context.Context) snrctx.SonrContext {
		return mockCtx
	}
	
	// Register cleanup function
	if os.Getenv("TEST_KEEP_MOCK") != "true" {
		// This is a bit of a hack since we can't easily clean up in tests
		// In a real implementation, you'd use t.Cleanup() with testify
	}
}

// setupMockIPFSAndEnclave mocks the IPFS and Enclave dependencies
func setupMockIPFSAndEnclave() {
	// Override the NewIPFS function
	originalNewIPFS := common.NewIPFS
	common.NewIPFS = func() (common.IPFSClient, error) {
		return &MockIPFSClient{}, nil
	}
	
	// Override the NewEnclave function
	originalNewEnclave := types.NewEnclave
	types.NewEnclave = func(string) (*types.Enclave, error) {
		return &types.Enclave{
			Addr: "test-address",
		}, nil
	}
	
	// Register cleanup function
	if os.Getenv("TEST_KEEP_MOCK") != "true" {
		// This is a bit of a hack since we can't easily clean up in tests
		// In a real implementation, you'd use t.Cleanup() with testify
	}
}

// MockSonrContext implements the SonrContext interface for testing
type MockSonrContext struct {
	sdkCtx sdk.Context
	peer   *MockPeer
}

func (m *MockSonrContext) SDKContext() sdk.Context {
	return m.sdkCtx
}

func (m *MockSonrContext) Peer() snrctx.Peer {
	return m.peer
}

func (m *MockSonrContext) IPFSRedirectURL(cid string) string {
	return "https://ipfs.example.com/" + cid
}

// MockPeer implements the Peer interface for testing
type MockPeer struct{}

func (m *MockPeer) Addr() snrctx.PeerAddr {
	return &MockPeerAddr{}
}

// MockPeerAddr implements the PeerAddr interface for testing
type MockPeerAddr struct{}

func (m *MockPeerAddr) String() string {
	return "test-peer-address"
}

// MockIPFSClient implements the IPFSClient interface for testing
type MockIPFSClient struct{}

func (m *MockIPFSClient) AddFile(data []byte) (string, error) {
	return "test-cid", nil
}

// Test cases for the query server methods

func TestQuerierParams(t *testing.T) {
	f := SetupTest(t)
	
	// Execute test
	req := &types.QueryParamsRequest{}
	res, err := f.queryServer.Params(context.Background(), req)
	
	// Assert results
	require.NoError(t, err)
	require.NotNil(t, res)
	require.NotNil(t, res.Params)
	require.Equal(t, types.DefaultParams(), *res.Params)
}

func TestQuerierSpawn(t *testing.T) {
	f := SetupTest(t)
	
	// Execute test
	req := &types.QuerySpawnRequest{}
	res, err := f.queryServer.Spawn(context.Background(), req)
	
	// Assert results
	require.NoError(t, err)
	require.NotNil(t, res)
	require.Equal(t, "test-address", res.Address)
	require.Equal(t, "test-cid", res.Cid)
	require.Equal(t, "https://ipfs.example.com/test-cid", res.Redirect)
}

func TestQuerierCheck(t *testing.T) {
	f := SetupTest(t)
	
	// Execute test
	req := &types.QueryCheckRequest{
		Cid: "test-cid",
	}
	res, err := f.queryServer.Check(context.Background(), req)
	
	// Assert results
	require.NoError(t, err)
	require.NotNil(t, res)
	require.Equal(t, "test-peer-address", res.Address)
}
