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

	var bookings [100]string

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

	remainingTickets -= userTickets
	bookings[0] = firstName + " " + lastName

	fmt.Printf("bookings array: %s\n", bookings)
	fmt.Printf("first element in bookings array: %s\n", bookings[0])
	fmt.Printf("array type: %T\n", bookings)
	fmt.Printf("array length: %d\n", len(bookings))

	fmt.Printf("\n%s %s booked %d tickets. A confirmation will be sent to the email address '%s'.\n\n",
		firstName, lastName, userTickets, email)

	fmt.Printf("There are now %d tickets remaining for %s.\n\n\n", remainingTickets, eventName)
}
