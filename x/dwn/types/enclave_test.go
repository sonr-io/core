package types_test

import (
	"testing"

	"github.com/sonr-io/core/x/dwn/types"
	"github.com/sonr-io/crypto/mpc"
)

// TestEnclave tests the enclave
func TestEnclave(t *testing.T) {
	enc, err := types.NewEnclave("test")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(enc.Addr)
	t.Log(enc.PubKeyHex)
}

// TestEnclaveFile tests the enclave file
func TestEnclaveFile(t *testing.T) {
	enc, err := types.NewEnclave("test")
	if err != nil {
		t.Fatal(err)
	}
	file, err := enc.File()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(file.Name())
}

// TestEnclaveIPFSAdd tests the enclave file add
func TestEnclaveIPFSAdd(t *testing.T) {
	enc, err := types.NewEnclave("test")
	if err != nil {
		t.Fatal(err)
	}
	file, err := enc.File()
	if err != nil {
		t.Fatal(err)
	}
	client, err := types.NewIPFSClient()
	if err != nil {
		t.Fatal(err)
	}
	cid, err := client.AddFile(file)
	if err != nil {
		panic(err)
		// t.Fatal(err)
	}
	if cid == "" {
		t.Fatal("cid is empty")
	}
	t.Log(cid)
}

func TestEnclaveIPFSGet(t *testing.T) {
	enc, err := types.NewEnclave("test")
	if err != nil {
		t.Fatal(err)
	}
	client, err := types.NewIPFSClient()
	if err != nil {
		t.Fatal(err)
	}
	file, err := enc.File()
	if err != nil {
		t.Fatal(err)
	}
	cid, err := client.AddFile(file)
	if err != nil {
		panic(err)
		// t.Fatal(err)
	}
	if cid == "" {
		t.Fatal("cid is empty")
	}
	t.Log(cid)
	got, err := client.Get(cid)
	if err != nil {
		t.Fatal(err)
	}
	if got == nil {
		t.Fatal("got is nil")
	}
	e, err := mpc.ImportEnclave(mpc.WithEnclaveJSON(got))
	if err != nil {
		t.Fatal(err)
	}
	t.Log(e.PubKeyHex())
}
