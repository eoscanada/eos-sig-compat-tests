package main

import (
	"encoding/hex"
	"fmt"
	"os"
	"time"

	eos "github.com/eoscanada/eos-go"
	"github.com/eoscanada/eos-go/ecc"
)

// Generated with `openssl rand -hex 32`
var defaultMessage = "89529cb031c69eccc92f3e8492393a8688bd3d071d7346677b6ff59d314d5121"

// Generated with cleos wallet create_key
var privKeyToImport = "5JFhynQnFBYNTPDA9TiKeE7TmujNYaExcbZi9bsRUjhVxwZF4Mt"
var pubKeyToImport = "EOS5jSQLpKBHLaMtuzkftnYE6bCMA5Jxso8f22uZyKj6cDEp32eSj"

func errorCheck(prefix string, err error) {
	if err != nil {
		fmt.Printf("ERROR: %s: %s\n", prefix, err)
		os.Exit(1)
	}
}

func main() {
	fmt.Println("EOS signature compatibility tests from eos-go")
	fmt.Println("-------------------------------")

	digestBytes, err := hex.DecodeString(defaultMessage)
	errorCheck("Error decoding hex message", err)

	fmt.Printf("Importing existing private key to wallet : %s\n", privKeyToImport)
	keyBag := eos.NewKeyBag()
	err = keyBag.ImportPrivateKey(privKeyToImport)
	errorCheck("Unable to import private key", err)

	fmt.Printf("Signing message digest (%s) with imported key\n", defaultMessage)
	sig, err := keyBag.Keys[0].Sign(digestBytes)
	errorCheck("Error signing", err)
	fmt.Printf("Signature : %s\n", sig)

	pub, err := sig.PublicKey(digestBytes)
	fmt.Printf("Recovered key from signature : %s\n", pub)

	fmt.Println("-------------------------------")

	fmt.Printf("eos-go : Generating fresh keypair at %s\n", time.Now())
	freshKeypair, err := ecc.NewRandomPrivateKey()
	errorCheck("Error creating fresh keypair", err)
	fmt.Printf("Public Key : %s\nPrivate Key: %s\n", freshKeypair.PublicKey(), freshKeypair)

	fmt.Printf("Signing message digest (%s) with fresh key\n", defaultMessage)
	sig, err = freshKeypair.Sign(digestBytes)
	errorCheck("Error signing", err)
	fmt.Printf("Signature : %s\n", sig)

	pub, err = sig.PublicKey(digestBytes)
	errorCheck("Unable to recover pub key", err)
	fmt.Printf("Recovered key from signature : %s\n", pub)

	fmt.Println("-------------------------------")
	fmt.Println("-------------------------------")

	////// The following were generated using the KEOSD test suite
	////// We are validating that we recover the same public key from those sigs

	// ------------------------------------
	// Creating fresh key in the new wallet
	// New pub key : EOS5gWrScGTTMyieGGhFDAmrVtDCp3UYzwdE7VLoZQnFSiGcezE3H
	// Exported fresh private key : 5Kk2STsBpo6UkY5Uw8BQ1YeFjp2BGLiBEsC5h4TYYiRDb7y5BTR
	// Signing message digest (89529cb031c69eccc92f3e8492393a8688bd3d071d7346677b6ff59d314d5121) with fresh key
	// Signature : SIG_K1_KhAncPv4QcGwE2gTrwhPZyQLy23AoNTKPiuN1puiT1F2xrdz7zTWEFvTJnADGgAnrBHpe2YsRNQZJnhkKShCo1FRMfUec1
	// ------------------------------------

	sig, err = ecc.NewSignature("SIG_K1_KhAncPv4QcGwE2gTrwhPZyQLy23AoNTKPiuN1puiT1F2xrdz7zTWEFvTJnADGgAnrBHpe2YsRNQZJnhkKShCo1FRMfUec1")
	errorCheck("Unable to reconstruct sig", err)
	pub, err = sig.PublicKey(digestBytes)
	errorCheck("Unable to recover pub key", err)
	if pub.String() != "EOS5gWrScGTTMyieGGhFDAmrVtDCp3UYzwdE7VLoZQnFSiGcezE3H" {
		panic("Unable to recover same pub key")
	} else {
		fmt.Printf("KEOSD-generated Sig %s, recovered same key as expected %s\n", sig, pub)
	}

	// ------------------------------------
	// Creating fresh key in the new wallet
	// New pub key : EOS8DRz4tLoEPZ49hDZNdvra8aZTgKZpFEhGVnEqWmAF5TEfnwwnd
	// Exported fresh private key : 5JHQ3iwGHrGhW8pmDNc6SsjqrATKhfQFanLubLKoYpyFkP37Xc8
	// Signing message digest (89529cb031c69eccc92f3e8492393a8688bd3d071d7346677b6ff59d314d5121) with fresh key
	// Signature : SIG_K1_KhF1Vur6J9cNFTibbSXBKtkmBWmsCV5mNbPYws2En77AnV1ad3bpzWWp5U1rKkZ7oJRx29LBBLPVHnUsFN8AYxcGGLLzVC
	// ------------------------------------

	sig, err = ecc.NewSignature("SIG_K1_KhF1Vur6J9cNFTibbSXBKtkmBWmsCV5mNbPYws2En77AnV1ad3bpzWWp5U1rKkZ7oJRx29LBBLPVHnUsFN8AYxcGGLLzVC")
	errorCheck("Unable to reconstruct sig", err)
	pub, err = sig.PublicKey(digestBytes)
	errorCheck("Unable to recover pub key", err)
	if pub.String() != "EOS8DRz4tLoEPZ49hDZNdvra8aZTgKZpFEhGVnEqWmAF5TEfnwwnd" {
		panic("Unable to recover same pub key")
	} else {
		fmt.Printf("KEOSD-generated Sig %s, recovered same key as expected %s\n", sig, pub)
	}

	// ------------------------------------
	// Creating fresh key in the new wallet
	// New pub key : EOS5vppWGhMUPZ7TvktEgXjeK7cfdnLnAgVK1KEAQfoeVfakAoYR4
	// Exported fresh private key : 5K6kC8rZTq86DLbAJji8m2c6dcAaTFGgbaVmXhFNRtzeRzHduxC
	// Signing message digest (89529cb031c69eccc92f3e8492393a8688bd3d071d7346677b6ff59d314d5121) with fresh key
	// Signature : SIG_K1_KBttGR2KMmdht6SsugkLQnuNUyRbbFrMLqZG7AwpFtrBLnCN1jU4ZmMAAE6pfsgFmcYxxgMxDGjH9nWue6Ak9voaXQyV2z
	// ------------------------------------

	sig, err = ecc.NewSignature("SIG_K1_KBttGR2KMmdht6SsugkLQnuNUyRbbFrMLqZG7AwpFtrBLnCN1jU4ZmMAAE6pfsgFmcYxxgMxDGjH9nWue6Ak9voaXQyV2z")
	errorCheck("Unable to reconstruct sig", err)
	pub, err = sig.PublicKey(digestBytes)
	errorCheck("Unable to recover pub key", err)
	if pub.String() != "EOS5vppWGhMUPZ7TvktEgXjeK7cfdnLnAgVK1KEAQfoeVfakAoYR4" {
		panic("Unable to recover same pub key")
	} else {
		fmt.Printf("KEOSD-generated Sig %s, recovered same key as expected %s\n", sig, pub)
	}

}
