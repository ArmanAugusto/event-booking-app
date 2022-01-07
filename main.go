package main

import "fmt"

func main() {
	var eventName = "Developer Conference"
	const conferenceTickets = 100
	var remainingTickets = 100

	fmt.Println("\nWelcome to our", eventName, "Booking Application")
	fmt.Println("\nWe have", remainingTickets, "available tickets remaining out of", conferenceTickets)
	fmt.Println("\nGet your tickets here to attend")
	fmt.Println("")
}
