
package main

import "fmt"

func main() {
	fmt.Print("Hello WORLD")

	//creating variable in go
	//anothay that replace var conf_name  = "go conference" 
	//doesnt work with consts
	conf_name  := "go conference"

	//a constant variable whih value cannot change
	const ticketAmount = 50
	var ticket_avail uint = 50
	var firstName string
	var lastName string
	var emailAddress string
	//only positive numbers
	var userTickets uint

	//we want to be avlaieb to sub the amount of ticket taken out ofthe
	ticket_avail = ticket_avail - userTickets
 
	

	fmt.Printf("welcome to the %v hopefully you enjoy your time\n", conf_name)
	fmt.Printf("We have %v tickerts and so far have only %v left \n", ticketAmount, ticket_avail)

	fmt.Println("firstname")
	//poins to the memory address of firstname to "assign" value
	fmt.Scan(&firstName)
	fmt.Println("lastname")
	fmt.Scan(&lastName)
	fmt.Println("emailaddress")
	fmt.Scan(&emailAddress)
	fmt.Println("howmany tckets you want")
	fmt.Scan(&userTickets)
	


}