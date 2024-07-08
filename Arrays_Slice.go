package main

import "fmt"

func main() {

	var arr1 = [10]int{1, 2, 3, 4, 5, 5, 5, 7, 8, 9}
	arr2 := [20]string{"aaa", "bbb", "ccc"}

	for i := 0; i < len(arr1); i++ {
		fmt.Print(arr1[i])

		if i == len(arr1)-1 {
			break
		}
		fmt.Print(" - ")
	}
	fmt.Print("\n")
	for i := 0; i < len(arr2); i++ {
		fmt.Print(arr2[i] + " ")
	}

	var slice1 = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0, 0}
	slice2 := []string{"qqq", "rrr", "ttt", "ppp"}

	slice2 = append(slice2, "zzz")

	for i := 0; i < len(slice1); i++ {
		fmt.Print(slice1[i])
	}

	for i := 0; i < len(slice2); i++ {
		fmt.Print(slice2[i])
	}

	var Id = map[int]string{1: "Suraj", 2: "Rohan"}
	fmt.Println(Id)
}
