package main

import "fmt"

func main() {

	var conferenceName = "Go Conference"
	const conferenceTickets = 50
	var remainingTickets = 30


	fmt.Println("Welcome to", conferenceName, "booking application")
	fmt.Println("We have total of", conferenceTickets, "tickets and", remainingTickets, "are still available.")															
	fmt.Println("Get your tickets here to attend")
	
	var userName string
	var userTickets int

	userName = "John"
	userTickets = 5

	fmt.Printf( "User %v booked %v tickets" ,userName, userTickets)

}
