package main

import "fmt"

func main() {
	conferenceName := "Go Conference"
	const conftickets int = 50
	var remainingtickets uint = 50

	//fmt.Printf("Types: %T, %T, %T\n", conferenceName, conftickets, remainingtickets)

	fmt.Printf("Welcome in %v!\n", conferenceName)
	fmt.Printf("Tickets: %v remaining: %v\n", conftickets, remainingtickets)
	fmt.Print("Get your ticket!\n")

	var bookings []string

	var username string
	var usertickets uint

	fmt.Print("Write Name\n")
	fmt.Scan(&username)

	fmt.Print("How many tickets?\n")
	fmt.Scan(&usertickets)

	remainingtickets = remainingtickets - usertickets
	bookings = append(bookings, username)

	//fmt.Printf("slice 1st: %v\n", bookings[0])
	//fmt.Printf("slice: %v\n", bookings)
	//fmt.Printf("slice type: %T\n", bookings)
	//fmt.Printf("slice length: %v\n", len(bookings[0]))

	fmt.Printf("User %v Booked %v tickets\n", username, usertickets)
	fmt.Printf("%v tickets remaining for %v\n", remainingtickets, conferenceName)
	fmt.Printf("All bookings: %v\n", bookings)
}
