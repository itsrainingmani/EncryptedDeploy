package main

import (
	"io/ioutil"
	"fmt"
	"EncryptedDeploy/filecrypto"
)

func main() {
	b, err := ioutil.ReadFile("cred.txt")
	if err != nil {
		fmt.Print(err)
	}

	passBytes, err := ioutil.ReadFile("pass.txt")
	if err != nil {
		fmt.Print(err)
	}

	ciphertext, err := filecrypto.Seal(passBytes, b)

	fmt.Println(ciphertext)

	ioutil.WriteFile("encrypted-credentials.txt", ciphertext, 0644)
	//fmt.Println(b)
	//str := string(b)
	//fmt.Println(str)
}
