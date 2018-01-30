package main

import (
	"crypto/rand"
	"fmt"
	"os"
	"log"
	"encoding/base64"
)

func main() {
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
