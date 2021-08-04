package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetAllTickets(t *testing.T) {
	auth := NewAuth()
	accessor := Accessor{}
	err := accessor.GetAllTickets(auth)()
	assert.Nil(t, err, "GetAllTickets expected to return no error but returned an error.")
	assert.NotNil(t, accessor.Data, "GetAllTickets expected to return response object but got nil.")
}
