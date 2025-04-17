package types

import (
	"github.com/cosmos/cosmos-sdk/types/bech32"
)

// ComputeSonrAddr computes the Sonr address from a public key
func ComputeSonrAddr(pk []byte) (string, error) {
	sonrAddr, err := bech32.ConvertAndEncode("idx", pk)
	if err != nil {
		return "", err
	}
	return sonrAddr, nil
}

// ComputeBitcoinAddr computes the Bitcoin address from a public key
func ComputeBitcoinAddr(pk []byte) (string, error) {
	btcAddr, err := bech32.ConvertAndEncode("bc", pk)
	if err != nil {
		return "", err
	}
	return btcAddr, nil
}
