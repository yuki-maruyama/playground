package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func main(){
	fmt.Println("Enter your password")
	var s string
	fmt.Scanf("%s", &s)
	hash, err := bcrypt.GenerateFromPassword([]byte(s), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hashed password -> %s\n", string(hash))
	fmt.Println("Verify your password: ")
	fmt.Scanf("%s", &s)
	if err := bcrypt.CompareHashAndPassword(hash, []byte(s)); err != nil{
		panic(err)
	}
	fmt.Println("OK!")
}
