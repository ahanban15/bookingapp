package main 

import "fmt"

func main() {
	// fmt.Println("Hello world on new line!")
	// fmt.Print("Hello world!")
	// fmt.Print("Hello world! on same line!")

	const conferenceName = "Go conference"
	var conferenceTickets = 50 
	var remainingTickets = 50
	// fmt.Print("Welcome to our conference ticket booking application!")
	fmt.Print("Welcome to ", conferenceName, " Ticket Booking Application!")
	fmt.Print("Get your tickets here to attend!")

	fmt.Print("We have a total of ", conferenceTickets, " tickets avaialiable and ", remainingTickets, " tickets are remaining unsold")



	
}