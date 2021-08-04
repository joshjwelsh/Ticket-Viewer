package main

import (
	"log"
	"os"
)

func main() {

	auth := NewAuth()
	accessor := Accessor{}
	err := accessor.GetAllTickets(auth)()
	if err != nil {
		log.Fatalf("Main failed on Get: %v", err)

	}

	display := NewDisplay(os.Stdin, *accessor.Data)

	if err := display.CLI(); err != nil {
		log.Fatalf("Main failed on Display: %v", err)
	}

}
