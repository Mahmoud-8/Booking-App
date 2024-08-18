package main

import (
	"fmt"
	"strings"
)

func main() {

	conferenceName := "Go Conference"
	const conferenceTickets int = 50
	var remainingTickets uint = 50
	bookings := []string{}


	greetUsers()

	fmt.Printf("conferenceTickets is %T, remainingTickets is %T", conferenceTickets, remainingTickets)
	fmt.Println("Welcome to", conferenceName, "booking application")
	fmt.Println("We have total of", conferenceTickets, "tickets and", remainingTickets, "are still available.")
	fmt.Println("Get your tickets here to attend")

	for {

		var firstName string
		var lastName string
		var email string
		var userTickets uint

		fmt.Println("Enter your first name: ")
		fmt.Scan(&firstName)

		fmt.Println("Enter your last name: ")
		fmt.Scan(&lastName)

		fmt.Println("Enter your email address: ")
		fmt.Scan(&email)

		fmt.Println("Enter number of tickets: ")
		fmt.Scan(&userTickets)

		isValidName := len(firstName) >= 2 && len(lastName) >=2
		isValidEmail := strings.Contains(email, "@")
		isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets

		

		if isValidName && isValidEmail && isValidTicketNumber {
			remainingTickets = remainingTickets - userTickets
			bookings = append(bookings, firstName+" "+lastName)

			fmt.Printf("These are all our bookings: %v\n", bookings)

			fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
			fmt.Printf("%v Remaining tickets %v\n", remainingTickets, conferenceName)

			firstNames := []string{}
			for _, booking := range bookings {

				var names = strings.Fields(booking)
				firstNames = append(firstNames, names[0])
			}
			fmt.Printf("The first names of bookings are: %v\n ", firstNames)

			if remainingTickets == 0 {
				fmt.Println("All tickets are sold out")
				break
			}
		} else {
			if !isValidName {
				fmt.Println("Your name is invalid, try again")

			}
			if !isValidEmail {
				fmt.Println("Your email doesn't contain @ sign")
			}
			if !isValidTicketNumber {
				fmt.Println("Your ticket number is invalid, try again")
			}

			fmt.Printf("Your input data is invalid, try again")

		}
	}
}

func greetUsers() {
	fmt.Println("Welcome to Go Conference")
}
