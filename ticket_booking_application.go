package main

import (
	"fmt"
	"os"
)

func bookTicket(remaining uint) uint {
	var bookings = []string{} // creating slice for storing name of users who has booked tickets from our application
	var fname string          // first name of user
	var lname string          // last name of user
	var email string          // email of user
	var userTickets uint      // number of tickets user want to book

	// Asking user for their details
	fmt.Print("\n\nEnter Your First Name : ")
	fmt.Scan(&fname)
	fmt.Print("Enter Your Last Name : ")
	fmt.Scan(&lname)
	fmt.Print("Enter Your Email : ")
	fmt.Scan(&email)
	fmt.Print("Enter number of tickets you want to book : ")
	fmt.Scan(&userTickets)
	remaining -= userTickets
	bookings = append(bookings, fname+" "+lname)

	fmt.Printf("\n\nThank you %v %v for booking tickets from our Application.\n", fname, lname)
	fmt.Printf("You will receive a confirmation mail at %v ", email)
	fmt.Print("\n\nThere are all our bookings : ")
	var i int
	for i = 0; i > len(bookings); i++ {
		fmt.Printf("%v\n", bookings[i])
	}

	return remaining
}
func seatAvailbility(total uint, remaining uint) {
	fmt.Printf("\nTotal Tickets : %v\n", total)
	fmt.Printf("Remaining Tickets : %v\n", remaining)
}

func main() {
	var appName = "TBA"
	const totalTickets int = 200                   // total tickets
	var remainingTickets uint = uint(totalTickets) // number of tickets remaining

	for remainingTickets > 0 {
		fmt.Printf("\nWelcome to %v !!\n", appName)
		fmt.Println("1. Check seat availbility")
		fmt.Println("2. Book Ticket")
		fmt.Println("3. Exit")
		var choice int
		fmt.Printf("\nEnter your choice : ") //asking user for their choice
		fmt.Scan(&choice)
		switch choice {
		case 1:
			seatAvailbility(uint(totalTickets), remainingTickets)

		case 2:
			remainingTickets = bookTicket(remainingTickets)
		case 3:
			os.Exit(0)

		}
	}
}
