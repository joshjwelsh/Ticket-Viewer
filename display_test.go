package main

import (
	"io"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Valid values
func TestSelectOnePrompt(t *testing.T) {
	var tests = []struct {
		want  int
		input string
		size  []Ticket
	}{
		{
			10,
			"10\n",
			make([]Ticket, 20),
		}, {
			15,
			"15\n",
			make([]Ticket, 20),
		}, {
			25,
			"25\n",
			make([]Ticket, 25),
		},
		{
			1,
			"1",
			make([]Ticket, 1),
		}, {
			100,
			"100\n",
			make([]Ticket, 100),
		}, {
			-10,
			"-10\n",
			make([]Ticket, 10),
		},
	}

	for _, test := range tests {

		r := strings.NewReader(test.input)
		reader := CreateDevice(r)

		got, _ := selectOnePrompt(reader, test.size)

		assert.NotNilf(t, got, "selectOnePrompt(io.Reader, []Ticket) failed to return a int value.\nTicket size %v input bytes %v expected %v but got nil.", string(test.input[:]), len(test.size), test.want)
		if got != test.want {
			t.Errorf("selectOnePrompt(io.Reader, []Ticket) returned an unexpected value.\nTicket size %v input bytes %v expected %v but got %v.", len(test.size), string(test.input[:]), test.want, got)
		}
	}
}
func TestSelectOnePrompt_Errors(t *testing.T) {
	var tests = []struct {
		want  bool
		input string
		size  []Ticket
	}{
		{
			false,
			"String\n",
			make([]Ticket, 20),
		}, {
			false,
			"-1\n",
			make([]Ticket, 10),
		}, {
			false,
			"sentence\n",
			make([]Ticket, 3),
		}, {
			false,
			"10",
			make([]Ticket, 4),
		}, {
			false,
			"4000",
			make([]Ticket, 1),
		},
	}
	// test helper
	check := func(err error) bool {
		if err == nil {
			return false
		}
		return true
	}

	for _, test := range tests {

		r := strings.NewReader(test.input)
		reader := CreateDevice(r)

		_, ok := selectOnePrompt(reader, test.size)

		if test.want != check(ok) {
			t.Errorf("selectOnePrompt(io.Reader, []Ticket) expected an error and did not get one or got an error when one was not expected. Ticket size %v input %v", len(test.size), test.input)
		}
	}
}

func TestDisplay_Template(t *testing.T) {
	test := make([]Ticket, 100)
	page := NewPage(test)()
	t.Run("Should return nil", func(t *testing.T) {
		err := display(page)
		assert.Nil(t, err)
	})

}

func TestCont(t *testing.T) {
	tests := []struct {
		Input string
	}{
		{"100\n"},
		{"BIGDOG\n"},
		{"walrus\n"},
	}
	for _, test := range tests {
		reader := MockStdin(test.Input)
		device := CreateDevice(reader)
		_, ok := cont(device)
		assert.Nil(t, ok, "Cont(ReadDevice) expected no nil but got an error.")
	}

}

func TestMenu(t *testing.T) {
	testMenu := CreateMainMenu()
	assert.NotPanics(t, func() {
		menu(testMenu)
	})
}

func TestClear(t *testing.T) {
	ok := clear()
	assert.True(t, ok)

}

// Helper
func MockStdin(s string) io.Reader {
	reader := strings.NewReader(s)
	return reader
}
