package main

import (
	"strings"
)

func validateUserInput(firstName string, lastName string, emailID string, userTickets uint) (bool, bool, bool) {

//function export
// func ValidateUserInput(firstName string, lastName string, emailID string, userTickets uint, RemainingTickets uint) (bool, bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(emailID, "@")
	isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets

	return isValidName, isValidEmail, isValidTicketNumber
}