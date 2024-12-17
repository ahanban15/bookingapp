package main 

import "fmt"

func main() {
	// fmt.Println("Hello world on new line!")
	// fmt.Print("Hello world!")
	// fmt.Print("Hello world! on same line!")

	const conferenceName = "Go conference"
	var conferenceTickets = 50 
	
	//syntactic sugar
	remainingTickets := 50
	// fmt.Print("Welcome to our conference ticket booking application!")
	fmt.Println("Welcome to ", conferenceName, " Ticket Booking Application!")
	fmt.Println("Get your tickets here to attend!")

	// fmt.Print("We have a total of ", conferenceTickets, " tickets avaialiable and ", remainingTickets, " tickets are remaining unsold")

	//formatted output
	fmt.Printf("We have a total of %v tickets avaialiable and %v tickets are remaining unsold\n", conferenceTickets, remainingTickets)

	// placeholder to get the type of variables
	fmt.Printf("conferenceTickets is %T, conferenceName is %T\n", conferenceTickets, conferenceName)



	
}