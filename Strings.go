package main

import (
	"fmt"
	"strings"
)

func main() {

	var name = "Suraj Mandhane"
	var re = strings.Split(name, " ")

	fmt.Println(name)
	fmt.Println(re[0])
	fmt.Println(re[1])

	re1 := strings.Fields(name)
	fmt.Print(re1)
}
