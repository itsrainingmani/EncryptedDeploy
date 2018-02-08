package main

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"os"
	"time"

	ps "github.com/gorillalabs/go-powershell"
	"github.com/gorillalabs/go-powershell/backend"
	"github.com/tsmanikandan/EncryptedDeploy/crypto"
)

func generatePassword() []byte {
	key := make([]byte, 32)

	_, err := rand.Read(key)
	if err != nil {
		fmt.Println(err.Error())
	}
	// fmt.Println(key)
	return key
}

func main() {

	fmt.Println("Welcome to the Encryption and Password Generation Utility")

	fmt.Println("\nReading the credentials file...")
	b, err := ioutil.ReadFile("credentials.txt")
	if err != nil {
		fmt.Println(err.Error())
		time.Sleep(5000 * time.Millisecond)
		os.Exit(1)
	}

	passBytes := generatePassword()

	// fmt.Println("Encrypting Credentials...")
	ciphertext, err := crypto.Seal(passBytes, b)
	//fmt.Println(ciphertext)

	fmt.Println("\nGenerating Encrypted Credentials...")
	ioutil.WriteFile("encrypted.txt", ciphertext, 0644)

	keyAsString := base64.StdEncoding.EncodeToString(passBytes)

	fmt.Println("\nGenerating Key file...")
	file, err := os.Create("key.txt")
	if err != nil {
		fmt.Println(err.Error())
		time.Sleep(5000 * time.Millisecond)
		os.Exit(1)
	}
	defer file.Close()

	fmt.Fprintf(file, keyAsString)

	time.Sleep(2000 * time.Millisecond)

	back := &backend.Local{}

	// start a local powershell process
	shell, err := ps.New(back)
	if err != nil {
		panic(err)
	}
	defer shell.Exit()

	fmt.Println("\nGenerating packaged data...")
	// ... and interact with it
	stdout, stderr, err := shell.Execute(".\\go-bindata.exe -o data.go cisco/... encrypted.txt")
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println(stdout, stderr)


	fmt.Println("\nDone!")
	time.Sleep(5000 * time.Millisecond)
}
