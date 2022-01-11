package main

import (
	"fmt"
)

func main() {
	var Name = "TBA"
	const totalTickets int = 200
	var remainingTickets uint = uint(totalTickets)
	var bookings = []string{}
	var fname string
	var lname string
	var email string
	var userTickets uint

	fmt.Printf("\nWelcome to %v !!\n", Name)
	fmt.Printf("\nTotal Tickets : %v\n", totalTickets)
	fmt.Printf("Remaining Tickets : %v\n", remainingTickets)
	// Asking user for their details
	fmt.Print("\n\nEnter Your First Name : ")
	fmt.Scan(&fname)
	fmt.Print("\nEnter Your Last Name : ")
	fmt.Scan(&lname)
	fmt.Print("\nEnter Your Email : ")
	fmt.Scan(&email)
	fmt.Print("\nEnter number of tickets you want to book : ")
	fmt.Scan(&userTickets)
	remainingTickets -= userTickets
	bookings = append(bookings, fname+" "+lname)

	fmt.Printf("Thank you %v %v for booking tickets from our Application.\n", fname, lname)
	fmt.Printf("You will receive a confirmation mail at %v ", email)

}
