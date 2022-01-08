package main

import "fmt"

func main() {
	const conferenceTickets uint = 100
	var eventName string = "Developer Conference"
	var remainingTickets uint = 100

	fmt.Printf("\n\nWelcome to our %s Booking Application\n\n", eventName)
	fmt.Printf("We have %d available tickets remaining out of %d.\n", remainingTickets, conferenceTickets)
	fmt.Println("Get your tickets here to attend.")
	fmt.Printf("\n")

	var userName string
	var userTickets uint

	fmt.Printf("Enter your username:  ")
	fmt.Scan(&userName)
	fmt.Printf("\nHello %s.\n\n", userName)
	fmt.Printf("Enter amount of tickets you would like to book:  ")
	fmt.Scan(&userTickets)

	fmt.Printf("\nUser %s booked %d tickets.\n\n\n", userName, userTickets)
}
