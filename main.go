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



	fmt.Printf("conferenceTickets is %T, remainingTickets is %T", conferenceTickets, remainingTickets)
	fmt.Println("Welcome to", conferenceName, "booking application")
	fmt.Println("We have total of", conferenceTickets, "tickets and", remainingTickets, "are still available.")															
	fmt.Println("Get your tickets here to attend")
	
	

	for{
	
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
	
	remainingTickets = remainingTickets - userTickets
	 bookings = append(bookings,  firstName + " " + lastName)


	fmt.Printf("These are all our bookings: %v\n", bookings)


	
	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v Remaining tickets %v\n", remainingTickets, conferenceName)


	firstNames := []string{}
	for index, booking := range bookings {

		var names = strings.Fields(booking)
		var firstName = names[0]
	}
	fmt.Printf("These are all our bookings: %v\n ", bookings)
}
}
