package main

import (
	"crypto/rand"
	"fmt"
	"learn_ether/crypto"
)

func main(){
	key := make([]byte, 32)
	rand.Read(key)

	fmt.Println("hello world", len(key))
	privKey, _ := crypto.ToECDSA(key)
	fmt.Println("privKey ", privKey)
}