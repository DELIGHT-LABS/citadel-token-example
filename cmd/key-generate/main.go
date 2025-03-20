package main

import (
	"crypto/ed25519"
	"encoding/hex"
	"fmt"
)

func main() {
	pubkey, privkey, err := ed25519.GenerateKey(nil)
	if err != nil {
		panic(err)
	}

	accessPrivKeyHex := hex.EncodeToString(privkey)
	accessPubKeyHex := hex.EncodeToString(pubkey)

	fmt.Println("private key = ", accessPrivKeyHex)
	fmt.Println("public key = ", accessPubKeyHex)
}
