package main

import (
	"fmt"
	"log"
	"os"

	"github.com/ethereum/go-ethereum/accounts/keystore"
)

func main() {
	fmt.Println("Enter the password to encrypt the account")
	var password string
	fmt.Scanln(&password)
	ks := keystore.NewKeyStore("./wallet", keystore.StandardScryptN, keystore.StandardScryptP)
	account, err := ks.NewAccount(password)
	if err != nil {
		log.Fatal(err)
	}

	// write the password to a 'passwords.txt' file
	// this is not a secure way to store passwords
	f := "passwords.txt"
	file, err := os.OpenFile(f, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	if _, err := file.WriteString(password + "\n"); err != nil {
		log.Fatal(err)
	}

	fmt.Println(account.Address.Hex())
}
