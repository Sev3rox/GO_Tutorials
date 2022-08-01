package main

import (
	"booking-app/helpers"
	"fmt"
	"time"
	//"sync"
)

var conferenceName = "Go Conference"

const conftickets int = 50

var remainingtickets uint = 50
var bookings = make([]UserData, 0) //bookings := []string{}

type UserData struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets uint
}

//var wg = sync.WaitGroup{}

func main() {
	//fmt.Printf("Types: %T, %T, %T\n", conferenceName, conftickets, remainingtickets)

	greetUsers()

	for {
		firstName, lastName, email, userTickets := getUserInput()

		isValidName, isValidEmail, isValidTicketNumber := helpers.ValidateUserInput(firstName, lastName, email, userTickets, remainingtickets)

		if isValidName && isValidEmail && isValidTicketNumber {

			bookTicket(userTickets, firstName, lastName, email)
			go sendTicket(userTickets, firstName, lastName, email)

			firstnames := getFirstNames()
			fmt.Printf("First Names: %v\n", firstnames)

			if remainingtickets == 0 {
				fmt.Print("End")
				break
			}
		} else {
			if !isValidName {
				fmt.Println("first name or last name you entered is too short")
			}
			if !isValidEmail {
				fmt.Println("email address you entered doesn't contain @ sign")
			}
			if !isValidTicketNumber {
				fmt.Println("number of tickets you entered is invalid")
			}

			//continue
		}
	}
	//wg.Wait()
}

func greetUsers() {
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available.\n", conftickets, remainingtickets)
	fmt.Println("Get your tickets here to attend")
}

func getFirstNames() []string {
	firstnames := []string{}
	for _, booking := range bookings {
		firstnames = append(firstnames, booking.firstName)
	}
	return firstnames
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

func bookTicket(userTickets uint, firstName string, lastName string, email string) {
	remainingtickets = remainingtickets - userTickets

	var userData = UserData{
		firstName:       firstName,
		lastName:        lastName,
		email:           email,
		numberOfTickets: userTickets,
	}

	bookings = append(bookings, userData)
	fmt.Printf("Bookings: %v\n", bookings)
	//fmt.Printf("slice 1st: %v\n", bookings[0])
	//fmt.Printf("slice: %v\n", bookings)
	//fmt.Printf("slice type: %T\n", bookings)
	//fmt.Printf("slice length: %v\n", len(bookings[0]))

	fmt.Printf("User %v Booked %v tickets\n", userTickets, userTickets)
	fmt.Printf("%v tickets remaining for %v\n", remainingtickets, conferenceName)
}

func sendTicket(userTickets uint, firstName string, lastName string, email string) {
	time.Sleep(2 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)
	fmt.Println("#################")
	fmt.Printf("Sending ticket:\n %v \nto email address %v\n", ticket, email)
	fmt.Println("#################")
	//wg.Done()
}
