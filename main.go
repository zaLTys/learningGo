package main

import "fmt"

func main() {

	conferenceName : = "Go Conference"
	const conferenceTickets = 50
	var remainingTickets = 50

	fmt.Printf("confName is %T , confTickets is %T", conferenceName, conferenceTickets)
	fmt.Printf("Welcome to %v\n", conferenceName)
	fmt.Printf("Total %v, tickets remaining: %v\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here ")

	var userName string
	var userTickets int

	userName = "Povilas"
	userTickets = 5

	fmt.Printf("User %v has %v tickets\n", userName, userTickets)

}
