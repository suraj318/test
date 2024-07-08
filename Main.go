package main

import "fmt"

func main() {
	var confName = "Go Lang"
	const confTicket = 50
	var remainingTicket = 50
	var userName string
	var userTicket int

	fmt.Println("Welcome to", confName)
	fmt.Println("---------------------------------------")
	fmt.Println("Total Tickets are : ", confTicket)
	fmt.Println("Remaining Ticket :", remainingTicket)

	fmt.Printf("\nThere are %v total ticket and %v are remaining", confTicket, remainingTicket)
	fmt.Printf("\nThere are %T total ticket and %T are remaining\n", confTicket, remainingTicket)
	//fmt.Print(confTicket + remainingTicket)

	userName = "Tom"
	userTicket = 20
	fmt.Println(userName, " = = = = = ", userTicket)

}
