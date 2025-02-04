package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Functions")
	name := "Raghav Jeje"
	age := 99
	greet(name, age)

}

func greet(name string, age int) {

	fmt.Println("This is GREET() Function")
	fmt.Printf("Hello %v, The Age is %v", name, age)

	fName := []string{}
	var a = strings.Fields(name)
	fName = append(fName, a[1])

	fmt.Println(fName)
}
