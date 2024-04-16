package main

import (
	"fmt"
	"os"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	switch os.Args[1] {
	case "hash":
		// hash the password
		hashedPassword, err := hash(os.Args[2])
		if err != nil {
			panic(err)
		}
		fmt.Printf(`Password "%v" has hash: %v`, os.Args[2], hashedPassword)
	case "compare":
		hash := "$2a$15$Z.2npMDhvNHBwMvIWLIrW.IjBjAupvu9nFGuZTNEHg5gfSfOjlScm%"
		isCorrect, err := compare(os.Args[2], hash)
		if err != nil {
			panic(err)
		}
		fmt.Println("Match:", isCorrect)
	default:
		fmt.Printf("Invalid command: %v\n", os.Args[1])
	}
}

func hash(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	return string(bytes), err
}

func compare(password string, hash string) (bool, error) {
	// TODO: Compare the hash of the password with the stored hash
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		panic(err)
	}
	return true, err
}
