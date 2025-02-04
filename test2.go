package main

import "fmt"

func main() {

	var name []string
	count := 5
	count1 := count
	for i := 0; i <= count; i++ {
		fmt.Println(count1)
		var in string
		fmt.Scan(&in)
		name = append(name, in)

		count1 = count1 - 1

	}

	for _, i := range name {
		fmt.Println(i)
	}

}
