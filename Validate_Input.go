package main

import (
	"fmt"
	"strings"
)

func main() {

	var name string
	fmt.Print("Enter Name : ")
	fmt.Scan(&name)
	var email string
	fmt.Print("Enter Email : ")
	fmt.Scan(&email)

	fmt.Println(Validate(name, email))
	//fmt.Println(isEmailValid)
}

func Validate(name string, email string) (bool, bool) {
	isNameValid := len(name) >= 2 && len(name) < 10
	isEmailValid := strings.Contains(email, ".com") && strings.Contains(email, "@")

	return isNameValid, isEmailValid

}
