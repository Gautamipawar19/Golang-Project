package main

import (
	"fmt"
	"strings"
)

const conferenceTickets uint = 50

var conferenceName = "Go Conference"
var remainingTickets uint = conferenceTickets
var bookings []string

func main() {
	greetUsers()

	for {
		firstName, lastName, email, userTickets := getUserInput()

		isValidName, isValidEmail, isValidTicketNumber := validUserInput(firstName, lastName, email, userTickets)

		if isValidName && isValidEmail && isValidTicketNumber {

			bookTicket(userTickets, firstName, lastName, email)
			firstNames := getFirstNames()
			fmt.Printf("The first Name of bookings are : %v \n\n", firstNames)

			if remainingTickets == 0 {
				fmt.Println("Conference is fully booked. Thank you!")
				break
			}

		} else {
			if !isValidName {
				fmt.Println(" insert valid name")
			}
			if !isValidEmail {
				fmt.Println("Email is invalid")
			}
			if !isValidTicketNumber {
				fmt.Println("Ticket number invalid or not available")
			}
			fmt.Println("Please try again.")
		}
	}
}

func greetUsers() {
	fmt.Printf("Welcome to %v Ticket Booking System!\n", conferenceName)
	fmt.Printf(" TotalTickets: %v |  AvailableTickets: %v\n", conferenceTickets, remainingTickets)
}

func validUserInput(firstName, lastName, email string, userTickets uint) (bool, bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@") && strings.Contains(email, ".")
	isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets
	return isValidName, isValidEmail, isValidTicketNumber
}

func getUserInput() (string, string, string, uint) {
	var firstName, lastName, email string
	var userTickets uint

	fmt.Print("Enter First Name: ")
	fmt.Scan(&firstName)

	fmt.Print("Enter Last Name: ")
	fmt.Scan(&lastName)

	fmt.Print("Enter Email: ")
	fmt.Scan(&email)

	fmt.Print("Enter number of tickets: ")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

func bookTicket(userTickets uint, firstName, lastName, email string) {
	remainingTickets -= userTickets
	booking := firstName + " " + lastName
	bookings = append(bookings, booking)

	fmt.Printf("\n Booking is confirmed for %v %v!\n", firstName, lastName)
	fmt.Printf("  Tickets will be sent to %v\n", email)
	fmt.Printf("  Tickets booked: %v | Remaining: %v\n\n", userTickets, remainingTickets)
}

func getFirstNames() []string {
	var firstNames []string
	for _, booking := range bookings {
		names := strings.Fields(booking)
		firstNames = append(firstNames, names[0])
	}
	return firstNames
}
