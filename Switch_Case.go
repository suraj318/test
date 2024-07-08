package main

import "fmt"

func main() {

	fmt.Println("---------------------Welcome to SWITCH CASE----------------")

	city := "Pune"

	switch city {
	case "Raipur":
		fmt.Println("Welcome to Raipur")
	case "Rampur":
		fmt.Println("Welcome to Rampur")
	case "Nagar":
		fmt.Println("Welcome to Nagar")
	case "Pune":
		fmt.Println("Welcome to Pune")

	default:
		fmt.Print("Default Statement")
	}

}
