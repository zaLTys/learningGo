package main

import (
	"booking-app/shared"
	"fmt"
	"time"
)

const conferenceTickets = 50

var conferenceName = "Go Conference"
var remainingTickets uint = 50
var bookings = make([]UserData, 0)

type UserData struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets uint
}

func main() {
	greetUsers()

	fmt.Println("Pointer example")
	fmt.Println(remainingTickets)
	fmt.Println(&remainingTickets)

	for remainingTickets > 0 && len(bookings) < 50 {

		firstName, lastName, email, userTickets := collectUserInputs()

		var isValidName, isValidEmail, isValidUserTickets = shared.ValidateInputs(firstName, lastName, email, userTickets, remainingTickets)

		if isValidName && isValidEmail && isValidUserTickets {
			bookTicket(firstName, lastName, email, userTickets)
			go sendTicket(userTickets, firstName, lastName, email)

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
		firstNames = append(firstNames, booking.firstName)
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

	var userData = UserData{
		firstName,
		lastName,
		email,
		userTickets}

	bookings = append(bookings, userData)

	fmt.Printf("Thans %v %v for reserving %v tickets on email %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("Total %v, tickets remaining: %v\n", conferenceTickets, remainingTickets)
}

func sendTicket(numberOfTickets uint, firstName string, lastName string, emailAddress string) {
	time.Sleep(10 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v", numberOfTickets, firstName, lastName)
	fmt.Println("##################")
	fmt.Printf("Sent ticket: \n %v \n to email %v \n \n ", ticket, emailAddress)
	fmt.Println("##################")
}
