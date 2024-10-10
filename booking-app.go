package main

import (
	"fmt"
)

type UserData struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets uint8
}

const conferenceName = "Go Conference"

var remainingTickets uint8 = 50
var bookings = []UserData{}

func main() {

	greetUsers()

	for {
		firstName, lastName, email, numberOfTickets := getUserInput()
		isNameValid, isEmailValid, isNumOfTicketsValid := validateUserInput(firstName, lastName, numberOfTickets, email, remainingTickets)

		if isNameValid && isEmailValid && isNumOfTicketsValid {
			bookTicket(firstName, lastName, numberOfTickets, email)
			if remainingTickets == 0 {
				fmt.Println("We're out of tickets")
				break
			}
		} else {
			printErrors(isEmailValid, isNameValid, isNumOfTicketsValid)
		}

	}

}

func greetUsers() {
	fmt.Printf("Welcome to %v \n", conferenceName)
	fmt.Printf("We got %v total tickets and %v tickets are available\n", remainingTickets, remainingTickets)
}

func getUserInput() (string, string, string, uint8) {
	var firstName string
	var lastName string
	var email string
	var numberOfTickets uint8

	fmt.Println("Enter your first Name")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last Name")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email")
	fmt.Scan(&email)

	fmt.Println("Enter how many tickets you want")
	fmt.Scan(&numberOfTickets)

	return firstName, lastName, email, numberOfTickets
}

func bookTicket(firstName string, lastName string, numberOfTickets uint8, email string) {
	remainingTickets = remainingTickets - numberOfTickets
	bookings = append(bookings, UserData{firstName: firstName, lastName: lastName, email: email, numberOfTickets: numberOfTickets})
	fmt.Printf("Thank you for your purchase of %v tickets\n", numberOfTickets)
	fmt.Printf("You will receive a confirmation to %v shortly\n", email)
}

func printErrors(isEmailValid bool, isNameValid bool, isNumOfTicketsValid bool) {
	if !isEmailValid {
		fmt.Println("You have entered invalid email")
	}

	if !isNameValid {
		fmt.Println("You have entered invalid name")
	}

	if !isNumOfTicketsValid {
		fmt.Println("You have entered invalid ticket amount")
	}

}
