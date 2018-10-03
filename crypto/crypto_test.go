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

func TestCryptoGenerateKey(t *testing.T){
	privateKey, _ := crypto.GenerateKey()
	if privateKey == nil {
		t.Fail()
	}
	t.Log("privateKey in Test XXX", privateKey)
}

func TestGeneratePubKey(t *testing.T){
	privateKey, _ := crypto.GenerateKey()
	public := privateKey.PublicKey
	t.Log("publicKey :", public)
}

func TestPubkeyToAddress(t *testing.T) {
	privateKey, _ := crypto.GenerateKey()
	public := privateKey.PublicKey

	address := crypto.PubkeyToAddress(public)
	t.Log("address :: ", address)
	t.Log("Hex address :: ", address.Hex())
}



