package main
import "fmt"

func main() {
	var conferenceName = "Go Conference"
	const totalTickets =50
	var availableTickets uint=totalTickets
    fmt.Printf("Welcome to %v booking application\n",conferenceName)
	fmt.Println("Total tickets:",totalTickets,"available tickets:",availableTickets)


	var firstName string
	var lastName string
	var email string
	var userTickets uint

	fmt.Println("Enter your first name")
	fmt.Scan(&firstName)
	fmt.Println("Enter your last name")
	fmt.Scan(&lastName)
	fmt.Println("Enter your email")
	fmt.Scan(&email)
	fmt.Println("Enter number of tickets")
	fmt.Scan(&userTickets)

	fmt.Printf("Dear user %v %v, you have booked %v tickets. You will receive a confirmation email at %v\n",firstName,lastName,userTickets,email)
	
}