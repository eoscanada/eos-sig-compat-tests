package main

import (
	"encoding/hex"
	"fmt"
	"time"

	eos "github.com/eoscanada/eos-go"
	"github.com/eoscanada/eos-go/ecc"
)

// Generated with `openssl rand -hex 32`
var defaultMessage = "89529cb031c69eccc92f3e8492393a8688bd3d071d7346677b6ff59d314d5121"

// Generated with cleos wallet create_key
var privKeyToImport = "5JFhynQnFBYNTPDA9TiKeE7TmujNYaExcbZi9bsRUjhVxwZF4Mt"
var pubKeyToImport = "EOS5jSQLpKBHLaMtuzkftnYE6bCMA5Jxso8f22uZyKj6cDEp32eSj"

func main() {
	fmt.Println("EOS signature compatibility tests from eos-go")
	fmt.Println("-------------------------------")

	digestBytes, err := hex.DecodeString(defaultMessage)
	if err != nil {
		panic("Error decoding hex message")
	}

	fmt.Printf("Importing existing private key to wallet : %s\n", privKeyToImport)
	keyBag := eos.NewKeyBag()
	err = keyBag.ImportPrivateKey(privKeyToImport)
	if err != nil {
		panic("Unable to import private key")
	}

	fmt.Printf("Signing message digest (%s) with imported key\n", defaultMessage)
	sig, err := keyBag.Keys[0].Sign(digestBytes)
	if err != nil {
		panic("Error signing")
	}
	fmt.Printf("Signature : %s\n", sig)

	pub, err := sig.PublicKey(digestBytes)
	fmt.Printf("Recovered key from signature : %s\n", pub)

	fmt.Println("-------------------------------")

	fmt.Printf("eos-go : Generating fresh keypair at %s\n", time.Now())
	freshKeypair, err := ecc.NewRandomPrivateKey()
	if err != nil {
		panic("Error creating fresh keypair")
	}
	fmt.Printf("Public Key : %s\nPrivate Key: %s\n", freshKeypair.PublicKey(), freshKeypair)

	fmt.Printf("Signing message digest (%s) with fresh key\n", defaultMessage)
	sig, err = keyBag.Keys[0].Sign(digestBytes)
	if err != nil {
		panic("Error signing")
	}
	fmt.Printf("Signature : %s\n", sig)

	pub, err = sig.PublicKey(digestBytes)
	fmt.Printf("Recovered key from signature : %s\n", pub)

	fmt.Println("-------------------------------")

}
