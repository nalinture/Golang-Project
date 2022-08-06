package main

import (
	"fmt"
	"strings"
)

func main() {
	// '<var name> := <value>' can be used instaed of 'var <var name> = <value>'. := can't be used for const variable.
	ConferenceName := "Golang Conference"
	//constant variable
	const totalTickets = 50
	// regular way of assigning variable
	var remainingTickets = 50

	//var bookings [50]string // array format - with index
	var bookings []string
	//bookings := []string{} //slices - without index
	//var firstNames []string

	fmt.Println()
	fmt.Printf("Welcome to %v !\n", ConferenceName)
	fmt.Printf("We have %v of %v tickets available to book.\n", remainingTickets, totalTickets)
	fmt.Println()

	// another way of declaring and calling a variable
	var prerequisite string
	prerequisite = "Perrequisite is 'an Intrest to learn'"
	fmt.Println(prerequisite)
	fmt.Println()

	for {

		var firstName string
		var lastName string
		var email string
		var ticketcount uint

		//user inputs
		fmt.Println("Enter your firstName: ")
		fmt.Scan(&firstName)

		fmt.Println("Enter your lastName: ")
		fmt.Scan(&lastName)

		fmt.Println("Enter your email address: ")
		fmt.Scan(&email)

		fmt.Println("How many tickets you need? ")
		fmt.Scan(&ticketcount)

		isValidName := len(firstName) > 2 && len(lastName) > 1
		isValidEmail := strings.Contains(email, "@")
		isValidTC := ticketcount > 0 && ticketcount <= uint(remainingTickets)

		if isValidName && isValidEmail && isValidTC {
			//Logics
			remainingTickets = remainingTickets - int(ticketcount)
			//bookings[0] = firstName + " " + lastName - array format. Its a fixed size. If we dont know about the size while coding use append() which is called slices.
			bookings = append(bookings, firstName+" "+lastName)
			//slices format - append() is a build-in function to store n number of values.

			/*fmt.Printf("The whole slices : %v\n", bookings)
			  fmt.Printf("The first index : %v\n", bookings[0])
			  fmt.Printf("The slices type : %T\n", bookings)
			  fmt.Printf("The slices length : %v\n", len(bookings))*/

			fmt.Printf("We have %v ticket(s) more.\n", remainingTickets)
			fmt.Printf("Thank you %v %v for booking %v tickets. A Confirmation email has been sent to %v\n", firstName, lastName, ticketcount, email)

			firstNames := []string{}
			for _, booking := range bookings {
				var names = strings.Fields(booking)
				firstNames = append(firstNames, names[0])

			}
			//firstNames = append(firstNames, firstName)

			fmt.Printf("these are the firstNames of bookings: %v\n", firstNames)
			//fmt.Printf("These are all our bookings: %v\n", bookings)

			if remainingTickets == 0 {
				fmt.Println("Tickets sold out!")
				break
			}
		} else {

			if !isValidName {
				fmt.Println("firstName or lastName is too short")
			}

			if !isValidEmail {
				fmt.Println("This is not a right email format.")
			}

			if !isValidTC {
				fmt.Printf("Please enter a valid ticketcount. Currently we have only %v tickets.\n", remainingTickets)
			}

		}

	}

}
