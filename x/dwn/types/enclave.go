package types

import (
	"encoding/hex"
	"encoding/json"

	"github.com/cosmos/cosmos-sdk/types/bech32"
	"github.com/sonr-io/core/common"
	"github.com/sonr-io/crypto/mpc"
)

type Enclave struct {
	Data      *mpc.EnclaveData
	Addr      string `json:"addr"`
	ChainID   string `json:"chain_id"`
	PubKeyHex string `json:"pub_key_hex"`
}

// NewEnclave creates a new enclave
func NewEnclave(chainID string) (*Enclave, error) {
	data, err := mpc.NewEnclave()
	if err != nil {
		return nil, err
	}
	addr, err := getSonrAddress(data)
	if err != nil {
		return nil, err
	}
	return &Enclave{
		Data:      data.GetData(),
		Addr:      addr,
		ChainID:   chainID,
		PubKeyHex: data.PubKeyHex(),
	}, nil
}

// FileBytes returns the bytes of the enclave file
func (e *Enclave) File() (common.File, error) {
	bz, err := json.Marshal(e)
	if err != nil {
		return nil, err
	}
	return common.NewFile("motr.jsonr", bz), nil
}

// getSonrAddress returns the Sonr address for the given enclave
func getSonrAddress(enclave mpc.Enclave) (string, error) {
	pkhex := enclave.PubKeyHex()
	pkBz, err := hex.DecodeString(pkhex)
	if err != nil {
		return "", err
	}
	return bech32.ConvertAndEncode("idx", pkBz)
}
