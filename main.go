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

	greetUsers(conferenceName, conferenceTickets, remainingTickets)

	for {

	firstName, lastName, email, userTickets :=	getUserInput ()
	isValidName, isValidEmail, isValidTicketNumber := validateUserInput(firstName, lastName, email, userTickets, remainingTickets)

		if isValidName && isValidEmail && isValidTicketNumber {
		

			 bookTicket(remainingTickets , userTickets , bookings , firstName , lastName , email , conferenceName )


			firstNames := getFirstNames(bookings)
			fmt.Printf("The first names of bookings are: %v\n ", firstNames)

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

func greetUsers(confName string, confTickets int, remaingingTickets uint) {
	fmt.Println("Welcome to Go Conference", confName)
	fmt.Printf("We have total of %v tickets and %v are still availble.\n", confTickets, remaingingTickets)
	fmt.Println("Get your tickets here to attend")

}

func getFirstNames(bookings []string) []string {
	firstNames := []string{}
	for _, booking := range bookings {

		var names = strings.Fields(booking)
		firstNames = append(firstNames, names[0])
	}
	return firstNames
}

func validateUserInput(firstName string, lastName string, email string, userTickets uint, remainingTickets uint) (bool, bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets

	return isValidName, isValidEmail, isValidTicketNumber

}

func getUserInput() (string, string, string, uint) {
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
		return firstName, lastName, email, userTickets

}


func bookTicket(remainingTickets uint, userTickets uint, bookings []string, firstName string, lastName string, email string, conferenceName string ) {
	remainingTickets = remainingTickets - userTickets
	bookings = append(bookings, firstName +" "+lastName)

	fmt.Printf("These are all our bookings: %v\n", bookings)
	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v Remaining tickets %v\n", remainingTickets, conferenceName)

}