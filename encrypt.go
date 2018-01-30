package main

import (
	"io/ioutil"
	"fmt"
	"EncryptedDeploy/filecrypto"
	"encoding/base64"
	"os"
	"log"
	"crypto/rand"
)

func generatePassword() {
	key := make([]byte, 32)

	_, err := rand.Read(key)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(key)
	keyAsString := base64.StdEncoding.EncodeToString(key)

	file, err := os.Create("pass.txt")
	if err != nil {
		log.Fatal("Cannot create file", err)
	}
	defer file.Close()

	fmt.Fprintf(file, keyAsString)
}

func main() {

	generatePassword()

	b, err := ioutil.ReadFile("cred.txt")
	if err != nil {
		fmt.Print(err)
	}

	passBytes, err := ioutil.ReadFile("pass.txt")
	if err != nil {
		fmt.Print(err)
	}

	ciphertext, err := filecrypto.Seal(passBytes, b)
	//fmt.Println(ciphertext)
	ioutil.WriteFile("encrypted.txt", ciphertext, 0644)
}
