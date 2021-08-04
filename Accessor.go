package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	endpoint string = "tickets"
	method   string = http.MethodGet
)

type Accessor struct {
	Data *[]Ticket
}

func (a *Accessor) GetAllTickets(auth Auth) func() error {
	type response struct {
		Tickets []Ticket `json:"tickets"`
	}

	r := response{}

	return func() error {
		resp, err := auth.Login(method, endpoint)
		if err != nil {
			return fmt.Errorf("Login failed: %v", err)
		}
		bytesBody, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return fmt.Errorf("ReadAll failed on resp.Body: %v ", err)
		}
		err = json.Unmarshal(bytesBody, &r)
		if err != nil {
			return fmt.Errorf("Unable to unmarshal to struct: %v ", err)
		}

		temp := make([]Ticket, len(r.Tickets))
		copy(temp, r.Tickets)

		a.Data = &temp

		defer resp.Body.Close()

		return nil

	}

}
