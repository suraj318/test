package main

import "fmt"

func main() {

	for1 := []int{1, 2, 2, 3, 4, 5, 6}
	var for2 []int

	fmt.Print("Enter Number : ")
	for i := 1; i < 5; i++ {
		var in int
		fmt.Scanln(&in)
		for2 = append(for2, in)
		for2 = append(for2, i)
	}

	for _, i := range for1 {
		fmt.Print(i)
	}
	fmt.Println()

	for _, j := range for2 {
		fmt.Print(j)
	}

}
