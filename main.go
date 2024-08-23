package main

import (
	"fmt"
	"strconv"	
	"booking-app/helper"
)

const conferenceTickets uint = 50
const conferenceName = "Go Conference"
var remainingTickets uint = 50
var bookings = make([]map[string]string, 0)

func main() {
	greetUsers()

	for {
	firstName, lastName, email, userTickets :=	getUserInput ()
	isValidName, isValidEmail, isValidTicketNumber := helper.ValidateUserInput(firstName, lastName, email, userTickets, remainingTickets)

		if isValidName && isValidEmail && isValidTicketNumber {
		

			bookTicket(remainingTickets, userTickets, bookings, firstName, lastName, email)


			firstNames := getFirstNames()
			fmt.Println("The first names of bookings are: ", firstNames)

			if !isValidName {
				fmt.Println("Your name is invalid, try again")

			}
			if !isValidEmail {
				fmt.Println("Your email doesn't contain @ sign")
			}
			if !isValidTicketNumber {
				fmt.Println("Your ticket number is invalid, try again")
			}
		}
	}
}

func greetUsers() {
	fmt.Println("Welcome to Go Conference", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still availble.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

}

func getFirstNames() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		firstNames = append(firstNames, booking["firstName"])

	}
	return firstNames
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


func bookTicket(remainingTickets uint, userTickets uint, bookings []map[string]string, firstName string, lastName string, email string ) {
	remainingTickets = remainingTickets - userTickets


	// create a map for a user
	var userData = make(map[string]string)
	userData["firstName"] = firstName
	userData["lastName"] = lastName
	userData["email"] = email
	userData["numberOfTickets"] = strconv.FormatUint(uint64(userTickets), 10)

	bookings = append(bookings, userData)
	fmt.Printf("List of bookings is %v\n", bookings)

	fmt.Printf("These are all our bookings: %v\n", bookings)
	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v Remaining tickets %v\n", remainingTickets, conferenceName)

}