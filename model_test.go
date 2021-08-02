package main

import (
	"testing"
)

func TestCreateMainMenu(t *testing.T) {
	got := CreateMainMenu()
	if got.Main == false {
		t.Errorf("CreateMainMenu() returned a menu object with Menu.Main == %v", got.Main)
	}
	if got.ContAll == true {
		t.Errorf("CreateMainMenu() returned a menu object with Menu.ContAll == %v", got.ContAll)
	}
	if got.ViewOne == true {
		t.Errorf("CreateMainMenu() returned a menu object with Menu.ViewOne == %v", got.ViewOne)
	}
}

func TestCreateViewAllMenu(t *testing.T) {
	got := CreateViewAllMenu()
	if got.Main == true {
		t.Errorf("CreateViewAllMenu() returned a menu object with Menu.Main == %v", got.Main)
	}
	if got.ContAll == false {
		t.Errorf("CreateViewAllMenu() returned a menu object with Menu.ContAll == %v", got.ContAll)
	}
	if got.ViewOne == true {
		t.Errorf("CreateViewAllMenu() returned a menu object with Menu.ViewOne == %v", got.ViewOne)
	}
}

func TestCreateSelectMenu(t *testing.T) {
	got := CreateSelectMenu()
	if got.Main == true {
		t.Errorf("CreateSelectMenu() returned a menu object with Menu.Main == %v", got.Main)
	}
	if got.ContAll == true {
		t.Errorf("CreateSelectMenu() returned a menu object with Menu.ContAll == %v", got.ContAll)
	}
	if got.ViewOne == false {
		t.Errorf("CreateSelectMenu() returned a menu object with Menu.ViewOne == %v", got.ViewOne)
	}

}

func TestCreateNewPage(t *testing.T) {
	want := Page{
		MaxPageSize:  1,
		CurrentPage:  1,
		CanGoBack:    false,
		CanGoForward: true,
		Start:        0,
		End:          1,
		FullSlice:    make([]Ticket, 1),
		All:          true,
	}
	ticket := make([]Ticket, 1)
	got := NewPage(ticket)()
	if got.MaxPageSize != want.MaxPageSize {
		t.Errorf("NewPage([]Ticket) returned a MaxPageSize of %v but expected %v.", got.MaxPageSize, got.MaxPageSize)
	}
	if got.CurrentPage != 1 {
		t.Errorf("NewPage([]Ticket) returned a CurrentPage value of %v but expected %v.", got.CurrentPage, want.CurrentPage)
	}
	if got.CanGoBack != false {
		t.Errorf("NewPage([]Ticket) returned a CanGoBack value of %v but expected %v.", got.CanGoBack, want.CanGoBack)
	}
	if got.CanGoForward != true {
		t.Errorf("NewPage([]Ticket) returned a CanGoForward value of %v but expected %v.", got.CanGoForward, want.CanGoForward)
	}
	if got.Start != 0 {
		t.Errorf("NewPage([]Ticket) returned a Start value of %v but expected %v.", got.Start, want.Start)
	}
	if got.All != true {
		t.Errorf("NewPage([]Ticket) returned a All value of %v but expected %v.", got.All, want.All)
	}
}

func TestPageForward(t *testing.T) {
	var tests = []struct {
		startCheck int
		endCheck   int
		iterate    int
	}{
		{
			10,
			15,
			2,
		},
	}
	got := NewPage(make([]Ticket, 20))()
	for _, test := range tests {
		for i := 0; i < test.iterate; i++ {
			got.PageForward()
		}
		if got.Start != test.startCheck {
			t.Errorf("PageForward() incorrect page update. Got Start %v but expected %v", got.Start, test.startCheck)
		}
		if got.End != test.endCheck {
			t.Errorf("PageBack() incorrect page update. Got End %v but expected %v", got.End, test.endCheck)
		}

	}

}

func TestPageBack(t *testing.T) {
	var tests = []struct {
		startCheck int
		endCheck   int
		goForward  int
		goBack     int
	}{
		{
			40,
			45,
			10,
			2,
		},
	}
	got := NewPage(make([]Ticket, 100))()
	for _, test := range tests {
		for i := 0; i < test.goForward; i++ {
			got.PageForward()
		}
		for j := 0; j < test.goBack; j++ {
			got.PageBack()
		}
		if got.Start != test.startCheck {
			t.Errorf("PageBack() incorrect page update. Got Start %v but expected %v", got.Start, test.startCheck)
		}
		if got.End != test.endCheck {
			t.Errorf("PageBack() incorrect page update. Got End %v but expected %v", got.End, test.endCheck)
		}

	}

}

func TestFindMaxPage(t *testing.T) {
	var tests = []struct {
		inputSize int
		want      int
	}{
		{
			14,
			3,
		},
	}
	for _, test := range tests {
		got := findMaxPage(test.inputSize)
		if got != test.want {
			t.Errorf("FindMaxPage(size int) returned incorrect value. Expected %v but got %v with input size %v", test.want, got, test.inputSize)
		}
	}

}

func TestPage(t *testing.T) {
	var tests = []struct {
		input []Ticket
		start int
		end   int
		want  int
	}{
		{
			make([]Ticket, 200),
			15,
			20,
			5,
		},
		{
			make([]Ticket, 18),
			15,
			18,
			3,
		},
	}
	for _, test := range tests {
		got := page(test.input, test.start, test.end)
		if len(got) != test.want {
			t.Errorf("page( []Ticket , int , int ) failed to return a []Ticket with correct size. Expected % v but got %v with Start %v and End %v", test.want, len(got), test.start, test.end)
		}
	}
}

func TestUpdateEnd(t *testing.T) {
	tests := []struct {
		currentEnd int
		size       int
		want       int
	}{
		{
			45,
			100,
			50,
		},
		{
			15,
			17,
			17,
		},
		{
			20,
			24,
			24,
		},
		{
			80,
			86,
			85,
		},
	}
	for _, test := range tests {
		curr := test.currentEnd
		result := updateEnd(curr, test.size)
		got := result + curr
		if got != test.want {
			t.Errorf("updateEnd(int, int) returned an incorrect value. Expected update to return %v but got %v with size %v and starting end value of %v", test.want, got, test.size, test.currentEnd)
		}
	}
}

func TestSelectOne(t *testing.T) {
	tests := []struct {
		PageSize []Ticket
		Index    int
		Want     bool
	}{
		{
			make([]Ticket, 100),
			10,
			true,
		}, {
			make([]Ticket, 100),
			1010,
			false,
		}, {
			make([]Ticket, 10),
			-10,
			false,
		}, {
			make([]Ticket, 50),
			25,
			true,
		},
	}
	for _, test := range tests {
		page := NewPage(test.PageSize)()
		got := page.SelectOne(test.Index)
		if got != test.Want {
			t.Errorf("(*Page)SelectOne(int) was expected to return %v but returned %v with input size %v", test.Want, got, len(test.PageSize))
		}
	}
}
