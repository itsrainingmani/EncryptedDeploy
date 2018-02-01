package main

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/tsmanikandan/EncryptedDeploy/crypto"
)

func generatePassword() []byte {
	key := make([]byte, 32)

	_, err := rand.Read(key)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(key)
	return key

}

func main() {

	passBytes := generatePassword()

	b, err := ioutil.ReadFile("cred.txt")
	if err != nil {
		fmt.Println(err.Error())
	}

	ciphertext, err := crypto.Seal(passBytes, b)
	//fmt.Println(ciphertext)
	ioutil.WriteFile("encrypted.txt", ciphertext, 0644)

	keyAsString := base64.StdEncoding.EncodeToString(passBytes)

	file, err := os.Create("passkey.txt")
	if err != nil {
		fmt.Println("Cannot create file", err)
	}
	defer file.Close()

	fmt.Fprintf(file, keyAsString)
}
