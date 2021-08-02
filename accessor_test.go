package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetAllTickets(t *testing.T) {
	got, err := GetAllTickets()()
	assert.Nil(t, err)
	assert.NotNil(t, got.Tickets)
}
