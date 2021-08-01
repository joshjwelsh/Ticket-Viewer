package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func GetAllTickets() func() (*TicketResponse, error) {
	const endpoint string = "tickets"
	const method string = http.MethodGet

	return func() (*TicketResponse, error) {
		var ticketList TicketResponse
		resp, err := Login(method, endpoint)
		if err != nil {
			return &ticketList, fmt.Errorf("Login failed: %v", err)
		}
		// getHeader(resp)
		bytesBody, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return &ticketList, fmt.Errorf("ReadAll failed on resp.Body: %v ", err)
		}
		if err := json.Unmarshal(bytesBody, &ticketList); err != nil {
			return &ticketList, fmt.Errorf("Unable to unmarshal to struct: %v ", err)
		}
		defer resp.Body.Close()
		return &ticketList, nil

	}

}

func getHeader(resp *http.Response) {
	for key, val := range resp.Header {
		fmt.Printf("%v - %v\n", key, val)
	}

}
