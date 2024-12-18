package main

import (
	"fmt"
	"strconv"  
	// "bookingapp/helper"
)

// package level variables
const conferenceName = "Go conference"

var conferenceTickets int = 50
var remainingTickets uint = 50
// var bookingSlice []string
var bookingSlice = make([]map[string]string, 0)

func main() {
	// fmt.Println("Hello world on new line!")
	// fmt.Print("Hello world!")
	// fmt.Print("Hello world! on same line!")

	// const conferenceName = "Go conference"
	// var conferenceTickets = 50

	//syntactic sugar
	// remainingTickets := uint(conferenceTickets)

	// greetUsers(conferenceName, conferenceTickets, remainingTickets)
	// no need for parameters as we now use package level variables
	greetUsers()

	// fmt.Print("Welcome to our conference ticket booking application!")
	// fmt.Println("Welcome to ", conferenceName, " Ticket Booking Application!")
	// fmt.Println("Get your tickets here to attend!")

	// fmt.Print("We have a total of ", conferenceTickets, " tickets avaialiable and ", remainingTickets, " tickets are remaining unsold")

	//formatted output
	// fmt.Printf("We have a total of %v tickets avaialiable and %v tickets are remaining unsold\n", conferenceTickets, remainingTickets)

	// placeholder to get the type of variables
	// fmt.Printf("conferenceTickets is %T, conferenceName is %T\n", conferenceTickets, conferenceName)

	for {
		// var firstName string
		// var lastName string
		// var emailID string
		// var userTickets uint

		// fmt.Println("Enter your first name:")
		// fmt.Scan(&firstName)

		// fmt.Println("Enter your last name:")
		// fmt.Scan(&lastName)

		// fmt.Println("Enter your emailID:")
		// fmt.Scan(&emailID)

		// fmt.Println("Enter the number of tickets you want:")
		// fmt.Scan(&userTickets)

		// isValidName := len(firstName) >= 2 && len(lastName) >= 2
		// isValidEmail := strings.Contains(emailID, "@")
		// isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets

		firstName, lastName, emailID, userTickets := getUserInput()
		isValidName, isValidEmail, isValidTicketNumber := validateUserInput(firstName, lastName, emailID, userTickets)

		if isValidName && isValidEmail && isValidTicketNumber {
			

			// if userTickets > remainingTickets{
			// 	fmt.Println("%v tickets unavailable \n", userTickets)
			// 	continue
			// }
			// remainingTickets = remainingTickets - userTickets
			// fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)

			//arrays and slices
			// var emptyArray [50]string

			// var users = [50]string{"John", "Peter", "Ahan"}
			// var bookings [50]string

			// bookings[0] = firstName + " " + lastName
			// fmt.Printf("The whole array is: %v\n", bookings)
			// fmt.Printf("The first value is: %v\n", bookings[0])
			// fmt.Printf("The array type is: %T\n", bookings)
			// fmt.Printf("The array length is: %v\n", len(bookings))

			// bookingSlice = append(bookingSlice, firstName+" "+lastName)

			// fmt.Printf("The whole slice is: %v\n", bookingSlice)
			// fmt.Printf("The first value is: %v\n", bookingSlice[0])
			// fmt.Printf("The slice type is: %T\n", bookingSlice)
			// fmt.Printf("The slice length is: %v\n", len(bookingSlice))

			// firstNames := []string{}
			// for _, booking := range bookingSlice {
			// 	var names = strings.Fields(booking)
			// 	var firstName = names[0]

			// 	firstNames = append(firstNames, firstName)
			// }\]

			// bookTicket(userTickets, firstName, lastName, emailID) 
			bookTicket(userTickets, firstName, lastName, emailID)

			// firstNames := getFirstNames(bookingSlice)
			firstNames := getFirstNames()

			fmt.Printf("The whole firstnames slice is: %v\n", firstNames)

			if remainingTickets <= 0 {
				fmt.Println("Come back next year")
				break
			}
		} else {
			if !isValidName {
				fmt.Println("The entered firstName or lastName is TOO SHort")
			}
			if !isValidEmail {
				fmt.Println("The entered email address is in invalid format")
			}
			if !isValidTicketNumber {
				fmt.Println("The entered number of tickets is invalid")
			}
		}
	}
}

// func greetUsers(conferenceName string, conferenceTickets int, remainingTickets uint) {
func greetUsers() {
	fmt.Println("Welcome to ", conferenceName, " Ticket Booking Application!")
	fmt.Println("Get your tickets here to attend!")
	fmt.Printf("We have a total of %v tickets avaialiable and %v tickets are remaining unsold\n", conferenceTickets, remainingTickets)
}

func getFirstNames() []string {
	firstNames := []string{}
	for _, booking := range bookingSlice {
		// var names = strings.Fields(booking)
		// var firstName = names[0]

		firstNames = append(firstNames, booking["firstName"])
	}

	return firstNames
}

// func validateUserInput(firstName string, lastName string, emailID string, userTickets uint) (bool, bool, bool) {
// 	isValidName := len(firstName) >= 2 && len(lastName) >= 2
// 	isValidEmail := strings.Contains(emailID, "@")
// 	isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets

// 	return isValidName, isValidEmail, isValidTicketNumber
// }

func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var emailID string
	var userTickets uint

	fmt.Println("Enter your first name:")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last name:")
	fmt.Scan(&lastName)

	fmt.Println("Enter your emailID:")
	fmt.Scan(&emailID)

	fmt.Println("Enter the number of tickets you want:")
	fmt.Scan(&userTickets)

	return firstName, lastName, emailID, userTickets
}

func bookTicket(userTickets uint, firstName string, lastName string, emailID string) {
	remainingTickets = remainingTickets - userTickets


	//create a map for userData

	var userData = make(map[string]string)
	userData["firstName"] = firstName
	userData["lastName"] = lastName
	userData["emailID"] = emailID
	userData["numberOfTickets"] = strconv.FormatUint(uint64(userTickets), 10)

	// bookingSlice = append(bookingSlice, firstName+" "+lastName)

	bookingSlice = append(bookingSlice, userData)

	fmt.Printf("List of bookings is %v\n", bookingSlice)

	fmt.Printf("User %v %v booked %v tickets with his ID %v\n", firstName, lastName, userTickets, emailID)

	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
}
