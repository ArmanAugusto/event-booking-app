package main

import "fmt"

func main() {
	const conferenceTickets uint = 100
	var eventName string = "Developer Conference"
	var remainingTickets uint = 100

	fmt.Printf("\n\nWelcome to our %s Booking Application\n\n", eventName)
	fmt.Printf("We have %d available tickets remaining out of %d.\n", remainingTickets, conferenceTickets)
	fmt.Printf("Get your tickets here to attend.")
	fmt.Printf("\n\n")

	var userName string
	var userTickets uint

	userName = "dev23"
	userTickets = 2
	fmt.Printf("User %s booked %d tickets.\n\n\n", userName, userTickets)
}
