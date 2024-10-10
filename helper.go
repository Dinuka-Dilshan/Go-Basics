package main

import (
	"strings"
)

func validateUserInput(firstName string, lastName string, numberOfTickets uint8, email string, remainingTickets uint8) (bool, bool, bool) {
	isNameValid := len(firstName) > 2 && len(lastName) > 2
	isEmailValid := strings.Contains(email, "@")
	isNumOfTicketsValid := numberOfTickets > 0 && numberOfTickets <= remainingTickets

	return isNameValid, isEmailValid, isNumOfTicketsValid
}
