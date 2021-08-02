package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {
	textPtr := flag.String("text", "", "Text to parse.")
	fmt.Println(textPtr)

	tickets, err := Get()
	if err != nil {
		log.Fatalf("Main failed on Get: %v", err)

	}
	if err := Display(os.Stdin, tickets.Tickets); err != nil {
		log.Fatalf("Main failed on Display: %v", err)
	}

}

func Get() (*TicketResponse, error) {
	GetAllTickets, err := GetAllTickets()()
	if err != nil {
		return &TicketResponse{}, fmt.Errorf("Get all Tickets call failed: %v", err)
	}
	return GetAllTickets, nil
}
