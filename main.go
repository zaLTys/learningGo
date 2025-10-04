package main

import (
	"fmt"
	"strings"
)

const conferenceTickets = 50

var conferenceName = "Go Conference"
var remainingTickets uint = 50
var bookings []string

func main() {
	greetUsers()

	fmt.Println("Pointer example")
	fmt.Println(remainingTickets)
	fmt.Println(&remainingTickets)

	for remainingTickets > 0 && len(bookings) < 50 {

		firstName, lastName, email, userTickets := collectUserInputs()

		var isValidName, isValidEmail, isValidUserTickets = validateInputs(firstName, lastName, email, userTickets)

		if isValidName && isValidEmail && isValidUserTickets {
			bookTicket(firstName, lastName, email, userTickets)

			fmt.Printf("Bookers %v\n", getFirstNames())

			noTicketsRemaining := remainingTickets == 0
			if noTicketsRemaining {
				fmt.Println("all booked, thanks!")
				break
			}
		} else {
			fmt.Printf("Your input data is invalid")
		}

		city := "London"
		switch city {
		case "Singapore":
			// sg booking
		case "London", "New York":
		default:
			fmt.Println("Unrecognized city")
		}
	}
}

func getFirstNames() []string {
	var firstNames []string
	for _, booking := range bookings {
		names := strings.Fields(booking)
		firstNames = append(firstNames, names[0])
	}
	return firstNames
}

func greetUsers() {
	fmt.Println("Welcome")
	fmt.Printf("confName is %T , confTickets is %T", conferenceName, conferenceTickets)
	fmt.Printf("Welcome to %v\n", conferenceName)
	fmt.Printf("Total %v, tickets remaining: %v\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here ")
}

func collectUserInputs() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint

	fmt.Println("Enter your first name:")
	fmt.Scanln(&firstName)
	fmt.Println("Enter your last name:")
	fmt.Scanln(&lastName)
	fmt.Println("Enter your email:")
	fmt.Scanln(&email)
	fmt.Println("Enter number of tickets:")
	fmt.Scanln(&userTickets)

	return firstName, lastName, email, userTickets
}

func bookTicket(firstName string, lastName string, email string, userTickets uint) {
	remainingTickets = remainingTickets - userTickets
	bookings = append(bookings, firstName+"  "+lastName)

	fmt.Printf("Thans %v %v for reserving %v tickets on email %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("Total %v, tickets remaining: %v\n", conferenceTickets, remainingTickets)
}
