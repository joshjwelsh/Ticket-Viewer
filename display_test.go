package main

import (
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

// type MockTestPaginate struct {
// 	ResultPaginate
// 	Input []string
// 	Size  []Ticket
// }
// type ResultPaginate struct {
// 	start int
// 	end   int
// }

// func TestPaginate(t *testing.T) {

// 	tests := []MockTestPaginate{
// 		{
// 			ResultPaginate: ResultPaginate{
// 				5,
// 				10,
// 			},
// 			Input: []string{"1\n3\n3\n"},
// 			Size:  make([]Ticket, 13),
// 		},
// 		{
// 			ResultPaginate: ResultPaginate{
// 				0,
// 				5,
// 			},
// 			Input: []string{"1\n2\n3\n3\n"},
// 			Size:  make([]Ticket, 200),
// 		},
// 	}
// 	for _, test := range tests {
// 		for _, s := range test.Input {

// 			testDevice := CreateDevice(MockStdin(s))
// 			page := NewPage(test.Size)()

// 			paginate(testDevice, page)
// 			if page.Start != test.start {
// 				t.Errorf("Paginate(ReadDevice, *Page) updated the Page struct incorrectly; expected %v but got %v", test.start, page.Start)
// 			}
// 		}
// 	}
// }

// func MockStdin(s string) io.Reader {
// 	reader := strings.NewReader(s)
// 	return reader
// }
