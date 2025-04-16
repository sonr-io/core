package types

import "github.com/sonr-io/core/common"

func NewIPFSClient() (common.IPFS, error) {
	return common.NewIPFS()
}
