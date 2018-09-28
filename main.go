package main

import (
	"encoding/hex"
	"fmt"
	"os"

	eos "github.com/eoscanada/eos-go"
	"github.com/eoscanada/eos-go/ecc"
)

func errorCheck(prefix string, err error) {
	if err != nil {
		fmt.Printf("ERROR: %s: %s\n", prefix, err)
		os.Exit(1)
	}
}

func importToKeybag(priv, expectedPub string) *eos.KeyBag {
	fmt.Println("=================================================================")
	fmt.Printf("Importing existing private key to wallet : %s\n", priv)
	keyBag := eos.NewKeyBag()
	err := keyBag.ImportPrivateKey(priv)
	errorCheck("Unable to import private key", err)
	if keyBag.Keys[0].PublicKey().String() != expectedPub {
		panic("Imported key does not match")
	}
	return keyBag
}

func signAndCheck(message, keosdSig string, keyBag *eos.KeyBag) {
	fmt.Println("---------------------------------------------")
	expectedPub := keyBag.Keys[0].PublicKey()
	fmt.Printf("Signing message digest (%s) with %s\n", message, expectedPub)
	bytes, err := hex.DecodeString(message)
	errorCheck("Cannot hex decode", err)
	sig, err := keyBag.Keys[0].Sign(bytes)
	errorCheck("Error signing", err)
	fmt.Printf("Signature : %s\n", sig)

	keosdSigObj, err := ecc.NewSignature(keosdSig)
	errorCheck("Cannot parse KEOSD sig", err)
	extractedPubFromSig, err := keosdSigObj.PublicKey(bytes)
	fmt.Printf("Recovered key from KEOSD signature : %s\n", extractedPubFromSig)
	if extractedPubFromSig.String() == expectedPub.String() {
		fmt.Println("Matches? true")
	} else {
		panic("Holy crap !!!!!")
	}
}

func main() {
	fmt.Println("EOS signature compatibility tests from eos-go")
	fmt.Println("-------------------------------")

	var keyBag *eos.KeyBag

	keyBag = importToKeybag("5JFhynQnFBYNTPDA9TiKeE7TmujNYaExcbZi9bsRUjhVxwZF4Mt", "EOS5jSQLpKBHLaMtuzkftnYE6bCMA5Jxso8f22uZyKj6cDEp32eSj")
	signAndCheck("89529cb031c69eccc92f3e8492393a8688bd3d071d7346677b6ff59d314d5121", "SIG_K1_KVp1bPmzswSvbcZCMENXbawKFVXPyYrUeJNZ9ChgWdhxLd5K8WtRmCtFY5cqVFgxjCZH8CwdNkxM3HBZ7EXeJmzcK78mHA", keyBag)
	signAndCheck("0000000000000000000000000000000000000000000000000000000000000000", "SIG_K1_KiDE7eMg8tHm2wnjJgxyXGtxgMX21hF6uXmCdaKHjp6dxjyweKRsNVKa8iV8h23a45rRP65VazY8R1cKsk89rK2ZzgjnJi", keyBag)
	signAndCheck("1111111111111111111111111111111111111111111111111111111111111111", "SIG_K1_K7SZ7Kch5ubBioPd2GpLUCcQAdDJize252xgQ1kYrAwYibGo4qAZHWETdMXo4yanwa7CLzHLM8SCDi9Fga1FMJTXTQ1ztV", keyBag)

	keyBag = importToKeybag("5Kk2STsBpo6UkY5Uw8BQ1YeFjp2BGLiBEsC5h4TYYiRDb7y5BTR", "EOS5gWrScGTTMyieGGhFDAmrVtDCp3UYzwdE7VLoZQnFSiGcezE3H")
	signAndCheck("89529cb031c69eccc92f3e8492393a8688bd3d071d7346677b6ff59d314d5121", "SIG_K1_KhAncPv4QcGwE2gTrwhPZyQLy23AoNTKPiuN1puiT1F2xrdz7zTWEFvTJnADGgAnrBHpe2YsRNQZJnhkKShCo1FRMfUec1", keyBag)
	signAndCheck("0000000000000000000000000000000000000000000000000000000000000000", "SIG_K1_KZQ3LcuJ741hKzGVjDNUPfYYBB5v88EZURFeyD1g3mxSx3MxzFAX5panzaAUyUugoLFRCxSw18qFtomtA6RSd8tRiQtKgs", keyBag)
	signAndCheck("1111111111111111111111111111111111111111111111111111111111111111", "SIG_K1_KfHdgQjHWdHetVcjc32PemVMsvSBxLSf5dvchauFa87QaXBaxadJ64UULasknsXhJLQu98vzEEgU6WbqEqJNebPtuJ7NEP", keyBag)
	signAndCheck("2222222222222222222222222222222222222222222222222222222222222222", "SIG_K1_Kf4omhweDa9fKDvR89qeVHGYfa2KccNdGvV7ng1uzVz7gERNLahRbGS5bZtWPdudUco7WApHXALTNZRVSvYknNoFKVAp8p", keyBag)
	signAndCheck("3333333333333333333333333333333333333333333333333333333333333333", "SIG_K1_KbBD8esKxZJiBtX1hfHGSCiytoNccoU5zKEHEWtWBdriUGwsmcRWkL8Zm1R3Tq7Uj9X7PW6vypfateHNuDKR8Az4maHhHo", keyBag)
	signAndCheck("4444444444444444444444444444444444444444444444444444444444444444", "SIG_K1_K8xFcoZKAvjCKZ5RGyVXKswhxwbztXjchMPRbk4vFc8jZaRpifMsDC57x6zjcPdmGzopsVLYmGu2EhfBzfqPBy3EB73KZq", keyBag)
	signAndCheck("5555555555555555555555555555555555555555555555555555555555555555", "SIG_K1_JwZVdQSqqpXVgdsBK2XtMdU13RTKsKmP16YxckMvMagC3nKpfAMSaz7Bnk4vWFyASsh2FuJoqGH9zXGtvDgXJvJo1qu9hU", keyBag)
	signAndCheck("BBBBBBBBBBBBBBBBBBBBBBBBCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCC", "SIG_K1_Ka1FvtKxHQzY2yUsM3jDJ3epGzrbXtUwS8kZLNu5iWDjbi7besBdHFu8RoR8b5fcJsrnwqBepeppFHfFoRtN9PLwDkyn8X", keyBag)

	keyBag = importToKeybag("5HxQKWDznancXZXm7Gr2guadK7BhK9Zs8ejDhfA9oEBM89ZaAru", "EOS7dwvuZfiNdTbo3aamP8jgq8RD4kzauNkyiQVjxLtAhDHJm9joQ")
	signAndCheck("6cb75bc5a46a7fdb64b92efefca01ed7b060ab5e0d625226e8efbc0980c3ddc1", "SIG_K1_K6jyndumSBa6P8tuKSNmsQrpZKTHZwW2FeZzEqs9sAZ23GXivhoZNGTrvvbsTmgsXmEmNeP1wL8vZUZWomb55Uz6MbHyD6", keyBag)
}
