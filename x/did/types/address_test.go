package types_test

import (
	"testing"

	"github.com/sonr-io/core/x/did/types"
	"github.com/sonr-io/crypto/mpc"
)

func mockPK() ([]byte, error) {
	e, err := mpc.NewEnclave()
	if err != nil {
		return nil, err
	}
	return e.PubKeyBytes(), nil
}

// TestAllAddress tests the Address List
func TestAllAddress(t *testing.T) {
	pk, err := mockPK()
	if err != nil {
		t.Fatal(err)
	}
	for _, c := range types.CoinList() {
		addr, err := types.FormatAddress(c, pk)
		if err != nil {
			t.Fatal("Failed to format address", err, c)
		}
		t.Log(addr)
	}
}
