package main

import (
	"fmt"
	"strings"
)

func main() {
	const conferenceTickets uint = 100
	var eventName string = "Developer Conference"
	var remainingTickets uint = 100
	bookings := []string{}

	fmt.Printf("\n\nWelcome to our %s Booking Application\n\n", eventName)
	fmt.Printf("We have %d available tickets remaining out of %d.\n", remainingTickets, conferenceTickets)
	fmt.Println("Get your tickets here to attend.")
	fmt.Printf("\n")

	for {
		var firstName string
		var lastName string
		var email string
		var userTickets uint

		fmt.Println("Enter the following information:")
		fmt.Print("\nFirst Name:  ")
		fmt.Scan(&firstName)
		fmt.Print("Last Name:  ")
		fmt.Scan(&lastName)
		fmt.Print("Email:  ")
		fmt.Scan(&email)
		fmt.Printf("\nHello %s.\n\n", firstName)
		fmt.Printf("Enter amount of tickets you would like to book:  ")
		fmt.Scan(&userTickets)

		if userTickets > remainingTickets {
			fmt.Printf("We only have %d tickets available.", remainingTickets)
			continue
		} else {

			remainingTickets = remainingTickets - userTickets
			bookings = append(bookings, firstName+" "+lastName)

			fmt.Printf("\n%s %s booked %d tickets. A confirmation will be sent to the email address '%s'.\n\n",
				firstName, lastName, userTickets, email)

			fmt.Printf("There are now %d tickets remaining for %s.\n\n\n", remainingTickets, eventName)

			// Slice of users' first namess
			firstNames := []string{}

			// places user first names from names variable into firstNames slice
			for _, booking := range bookings {
				var names = strings.Fields(booking)
				firstNames = append(firstNames, names[0])
			}
			fmt.Printf("Bookings (First Names):  %v\n\n\n", firstNames)

			if remainingTickets == 0 {
				// end program
				fmt.Println("Conference Tickets are sold out.")
				break
			}
		}
	}
}
