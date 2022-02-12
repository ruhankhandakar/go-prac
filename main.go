package main

import (
	"booking/helper"
	"fmt"
	"sync"
	"time"
)

const conferenceTickets int = 50
var conferenceName = "Go conference"
// conferenceName := "Go conference"
var remainingTickets uint = 50
var bookings  = make([]UserData, 0)

var wg = sync.WaitGroup{}

type UserData struct {
	firstName string
	lastName string
	email string
	userTickets uint
}

func main() {
	greetUser()

	firstName, lastName, email, userTickets := getUserInput()

	isValidName, isValidEmail, isValidTicketNumber := helper.ValidateUserInput(firstName, lastName, email, userTickets, remainingTickets)
	
	if isValidName && isValidEmail && isValidTicketNumber {
		bookTicket(userTickets, firstName, lastName, email)
		wg.Add(1)
		go sendTicket(userTickets, firstName, lastName, email)
		
		firstNames := getFirstNames()
		fmt.Printf("The first names of bookings are: %v \n", firstNames)

		noTicketsRemaining  := remainingTickets <= 0
		if noTicketsRemaining {
			fmt.Println("Our conference is booked out. Come back next year.")
			// break
		}
	} else {
		fmt.Printf("We only have %v tickets reamaining. So you can't book %v .\n", remainingTickets, userTickets)
		// continue
	}
	// city := "London"

	// switch city {
	// 	case "New York":
	// 		// Execute code for booking new york conference tickets
	// 	case "Singapore", "Hong Kong":
	// 		// 
	// 	case "London", "Berlin":
	// 		// 
	// 	case "Mexico City":
	// 		// 
	// 	default:
	// 		fmt.Println("No valid city selected")
	// }

	wg.Wait()
}

func greetUser(){
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	// fmt.Printf("conferenceTickets: %T, remainingTickets: %T,conferenceName: %T \n", conferenceTickets, remainingTickets, conferenceName)
	fmt.Printf("We  have total of %v tickets and %v are still available\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your ticket to attend")

}

func getFirstNames() []string {
	firstNames := []string{}
	for _, booking := range bookings {		
		firstNames = append(firstNames, booking.firstName)
	}

	return firstNames
}

func getUserInput() (string, string,string, uint){
	var firstName string
	var lastName string
	var email string
	var userTickets uint

	fmt.Printf("Please enter your First name ")
	fmt.Scan(&firstName)
	fmt.Printf("Please enter your last name ")
	fmt.Scan(&lastName)
	fmt.Printf("Please enter your email ")
	fmt.Scan(&email)
	fmt.Printf("Please Number of tickets ")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

func bookTicket(userTickets uint, firstName string, lastName string, email string){
	remainingTickets = remainingTickets - userTickets

	var userData = UserData{
		firstName: firstName,
		lastName: lastName,
		email: email,
		userTickets: userTickets,
	}
	bookings = append(bookings, userData)

	fmt.Printf("List of booking is %v\n", bookings)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v.\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets for %v\n", remainingTickets, conferenceName)
}

func sendTicket(userTickets uint, firstName string, lastName string, email string){
	time.Sleep(10 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)
	fmt.Println("############")
	fmt.Printf("Sending ticket:\n %v to email address %v\n", ticket, email)
	fmt.Println("############")

	wg.Done()
}