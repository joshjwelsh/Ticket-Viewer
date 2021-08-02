package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateDevice(t *testing.T) {
	reader := MockStdin("test")
	got := CreateDevice(reader)
	assert.NotNil(t, got)
	assert.NotNil(t, got.Reader)

}

func TestGetInput(t *testing.T) {
	reader := MockStdin("test")
	got := CreateDevice(reader)
	err := got.GetInput()
	assert.Nil(t, err)
}
