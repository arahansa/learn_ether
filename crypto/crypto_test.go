package crypto

import (
	"crypto/rand"
	"github.com/ethereum/go-ethereum/crypto"
	"testing"
)

func TestToECDSA(t *testing.T) {

	key := make([]byte, 32)
	rand.Read(key)

	privKey, _ := crypto.ToECDSA(key)
	t.Log(privKey)

	if privKey == nil {
		t.Fail()
	}
}
