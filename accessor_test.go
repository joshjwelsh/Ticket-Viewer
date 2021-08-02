package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetAllTickets(t *testing.T) {
	got, err := GetAllTickets()()
	assert.Nil(t, err, "GetAllTickets expected to return no error but returned an error.")
	assert.NotNil(t, got.Tickets, "GetAllTickets expected to return response object but got nil.")
}
