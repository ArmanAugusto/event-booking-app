package main

import (
	"fmt"
	"sync"
	"time"
)

// package level variables
const conferenceTickets uint = 100

var eventName string = "Developer Conference"
var remainingTickets uint = 100
var bookings = make([]UserData, 0)

type UserData struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets uint
}

//synchronizing
var wg = sync.WaitGroup{}

func main() {

	greetUsers()

	for {

		firstName, lastName, email, userTickets := getUserInput()
		isValidName, isValidEmail, isValidTicketAmount := validateUserInput(firstName, lastName, email, userTickets, remainingTickets)

		if isValidName && isValidEmail && isValidTicketAmount {
			bookTickets(userTickets, firstName, lastName, email)

			wg.Add(1)
			// go routine (concurrency)
			go sendTicket(userTickets, firstName, lastName, email)

			firstNames := getFirstNames()
			fmt.Printf("\nBookings (First Names):  %s\n\n", firstNames)

			if remainingTickets == 0 {
				// end program
				fmt.Println("Conference Tickets are sold out.")
				fmt.Println()
				fmt.Println()
				break
			}
		} else {
			if !isValidName {
				fmt.Println("First name and last name needs to have more than one letter.")
			}
			if !isValidEmail {
				fmt.Println("Invalid email entry.")
			}
			if !isValidTicketAmount {
				fmt.Println("Invalid ticket amount entered.")
			}
		}
		wg.Wait()
	}
}

func greetUsers() {
	fmt.Printf("\nWelcome to our %s Booking Application\n\n", eventName)
	fmt.Printf("We have %d available tickets remaining out of %d.\n", remainingTickets, conferenceTickets)
	fmt.Println("Get your tickets here to attend.")
	fmt.Printf("\n")
}

func getFirstNames() []string {
	// Slice of users' first namess
	firstNames := []string{}

	// places user first names from names variable into firstNames slice
	for _, booking := range bookings {
		firstNames = append(firstNames, booking.firstName)
	}
	return firstNames
}

func getUserInput() (string, string, string, uint) {
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

	return firstName, lastName, email, userTickets
}

func bookTickets(userTickets uint, firstName string, lastName string, email string) {
	remainingTickets = remainingTickets - userTickets

	// create a userData object for a user
	var userData = UserData{
		firstName:       firstName,
		lastName:        lastName,
		email:           email,
		numberOfTickets: userTickets,
	}

	bookings = append(bookings, userData)
	fmt.Printf("List of bookings is %v\n", bookings)

	fmt.Printf("\n%s %s booked %d tickets. A confirmation will be sent to the email address '%s'.\n\n",
		firstName, lastName, userTickets, email)

	fmt.Printf("There are now %d tickets remaining for %s.\n\n\n", remainingTickets, eventName)
}

func sendTicket(userTickets uint, firstName string, lastName string, email string) {
	time.Sleep(20 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %s %s", userTickets, firstName, lastName)
	fmt.Println("################")
	fmt.Printf("Sending ticket:\n%v \nto email address %s", ticket, email)
	fmt.Println("\n################")
	wg.Done()
}
