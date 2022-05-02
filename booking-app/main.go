package main

import (
	"fmt"
)

func main() {

	//array func
	arraysFunc()
	var conferenceName = "Go Conference"
	const conferenceTickets = 50
	var remainingTickets = conferenceTickets

	fmt.Printf("conferenceTickets is %T, remainingTickets is %T, conferenceName is %T\n", conferenceTickets, remainingTickets, conferenceName)
	fmt.Println("Welcome to" + conferenceName + "booking application")
	fmt.Println("We have total of ", conferenceTickets, "tickets and ", remainingTickets, "are still availavle")
	fmt.Println("Get your tickets here to attend\n")

	//v fo strug ?
	fmt.Printf("Wlecome to %v ", conferenceName)

	var username string
	//ask user for their name
	var userTickets int
	var firstName string
	var lastName string
	var email string

	userTickets = 2
	username = "Tom"
	fmt.Println(username)
	fmt.Printf("User %v booked %v tickets \n", username, userTickets)

	fmt.Println("Enter our first name:")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last name:")	
	fmt.Scan(&lastName)

	fmt.Println("Enter your email address:")
	fmt.Scan(&email)

	remainingTickets = remainingTickets - userTickets
	fmt.Printf("%v tickets remainig for %v \n", remainingTickets, username)
	
}

func arraysFunc(){
	var bookings = [50] string {"Pesho", "Nicole", "Peter"}
	bookings[0] = "Pesho"
	

}
