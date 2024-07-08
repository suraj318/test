package main

import (
	"fmt"
	"strconv"
)

func main() {

	fmt.Println("This is MAP...")
	fname := "Suraj"
	lname := "Mandhane"
	email := "surmandhane@gmail.com"
	age := 24

	var Data = make(map[string]string)

	Data["fname"] = fname
	Data["lname"] = lname
	Data["mail"] = email
	Data["age"] = strconv.FormatInt(int64(age), 10)

	fmt.Println(Data)
	// for i := range Data {
	// 	fmt.Println(i)
	// }

}
