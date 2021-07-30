package main

import (
	"flag"
	"fmt"
	"log"
)

// Endpoints
const (
	get_all_tickets = "tickets"
)

func main() {
	textPtr := flag.String("text", "", "Text to parse.")
	fmt.Println(textPtr)

	tickets, err := Get()
	if err != nil {
		log.Fatalf("Main failed on Get: %v", err)

	}

	for idx, ticket := range tickets.Tickets {
		fmt.Printf("%v) %v\n", idx, ticket.Subject)
	}

}

func Get() (Ticket, error) {
	GetAllTickets, err := GetAllTickets()()
	if err != nil {
		return Ticket{}, fmt.Errorf("Get all Tickets call failed: %v", err)
	}
	return GetAllTickets, nil
}
