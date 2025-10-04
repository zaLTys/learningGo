package main

import "fmt"

func main() {

	conferenceName := "Go Conference"
	const conferenceTickets = 50
	var remainingTickets uint = 50

	fmt.Printf("confName is %T , confTickets is %T", conferenceName, conferenceTickets)
	fmt.Printf("Welcome to %v\n", conferenceName)
	fmt.Printf("Total %v, tickets remaining: %v\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here ")

	var firstName string
	var lastName string
	var email string
	var userTickets uint

	fmt.Println("Pointer example")
	fmt.Println(remainingTickets)
	fmt.Println(&remainingTickets)

	fmt.Println("Enter your first name:")
	fmt.Scanln(&firstName)
	fmt.Println("Enter your last name:")
	fmt.Scanln(&lastName)
	fmt.Println("Enter your email:")
	fmt.Scanln(&email)
	fmt.Println("Enter number of tickets:")
	fmt.Scanln(&userTickets)

	var bookings []string
	bookings = append(bookings, firstName+"  "+lastName)

	remainingTickets = remainingTickets - userTickets

	fmt.Printf("Whole array %v\n", bookings)

	fmt.Printf("Thans %v %v for reserving %v tickets on email %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("Total %v, tickets remaining: %v\n", conferenceTickets, remainingTickets)
}
