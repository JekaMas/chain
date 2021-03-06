package chain

import (
	"golang.org/x/crypto/ed25519"
	"testing"
)

func TestBlockProcessing(t *testing.T) {
	pubkey, _, err := ed25519.GenerateKey(nil)
	if err != nil {
		t.Fatal(err)
	}

	validatorAddr, err := PubKeyToAddress(pubkey)
	if err != nil {
		t.Fatal(err)
	}

	nd := &Node{
		validators: []ed25519.PublicKey{pubkey},
	}
	nd.state = map[string]uint64{
		"one":         200,
		"two":         50,
		validatorAddr: 50,
	}

	err = nd.AddBlock(Block{
		BlockNum: 1,
		Transactions: []Transaction{
			{
				From:   "one",
				To:     "two",
				Fee:    10,
				Amount: 100,
			},
		},
	})
	t.Log(nd.blocks)
	if err != nil {
		t.Fatal(err)
	}
	if nd.state["one"] != 90 {
		t.Error()
	}
	if nd.state["two"] != 150 {
		t.Error()
	}
	if nd.state[validatorAddr] != 60 {
		t.Error()
	}
}
