package main 

import "fmt"

func main() {
	// fmt.Println("Hello world on new line!")
	// fmt.Print("Hello world!")
	// fmt.Print("Hello world! on same line!")

	const conferenceName = "Go conference"
	var conferenceTickets = 50 
	
	//syntactic sugar
	remainingTickets := uint(conferenceTickets)
	// fmt.Print("Welcome to our conference ticket booking application!")
	fmt.Println("Welcome to ", conferenceName, " Ticket Booking Application!")
	fmt.Println("Get your tickets here to attend!")

	// fmt.Print("We have a total of ", conferenceTickets, " tickets avaialiable and ", remainingTickets, " tickets are remaining unsold")

	//formatted output
	fmt.Printf("We have a total of %v tickets avaialiable and %v tickets are remaining unsold\n", conferenceTickets, remainingTickets)

	// placeholder to get the type of variables
	fmt.Printf("conferenceTickets is %T, conferenceName is %T\n", conferenceTickets, conferenceName)

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

	fmt.Printf("User %v %v booked %v tickets with his ID %v\n", firstName, lastName, userTickets, emailID) 

	remainingTickets = remainingTickets - userTickets
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
}