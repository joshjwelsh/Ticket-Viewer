package main

import (
	"log"
	"math"
)

// --------------------------------- Represents Page and update struct ---------------------------------------
type Page struct {
	MaxPageSize    int
	CurrentPage    int
	CanGoBack      bool
	CanGoForward   bool
	Start          int
	End            int
	CurrentSlice   []Ticket
	FullSlice      []Ticket
	All            bool
	SelectedTicket Ticket
	Select         int
}

func NewPage(t []Ticket) func() *Page {
	var lastIdx int
	size := len(t)
	max := findMaxPage(size)
	if max < MAX_PAGE_SIZE {
		lastIdx = (max % MAX_PAGE_SIZE)
	} else {
		lastIdx = MAX_PAGE_SIZE
	}
	return func() *Page {
		return &Page{
			MaxPageSize:  max,
			CurrentPage:  1,
			CanGoBack:    false,
			CanGoForward: true,
			Start:        0,
			End:          MAX_PAGE_SIZE,
			CurrentSlice: t[0:lastIdx],
			FullSlice:    t,
			All:          true,
		}

	}

}

func (p *Page) PageForward() {
	if p.CanGoForward == false {
		return
	}
	p.CurrentPage += 1
	p.Start += MAX_PAGE_SIZE
	p.End += updateEnd(p.End, len(p.FullSlice))
	p.CurrentSlice = page(p.FullSlice, p.Start, p.End)
	if p.CurrentPage == p.MaxPageSize {
		p.CanGoForward = false
	}
	if verbose {
		log.Printf("Page Up: start %v end %v\n", p.Start, p.End)
	}

}
func (p *Page) PageBack() {

	if p.CanGoBack == false {
		if p.CurrentPage != 1 {
			p.CanGoBack = true
		} else if p.CurrentPage == 1 {
			return
		}
	}
	p.CurrentPage -= 1
	p.Start -= MAX_PAGE_SIZE
	p.End -= MAX_PAGE_SIZE
	p.CurrentSlice = page(p.FullSlice, p.Start, p.End)
	if p.CurrentPage == 1 {
		p.CanGoBack = false
	}
	if verbose {
		log.Printf("Page down: start %v end %v\n", p.Start, p.End)
	}
}

func findMaxPage(size int) int {
	val := float64(size) / float64(MAX_PAGE_SIZE)
	return int(math.Ceil(val))
}

func page(ticket []Ticket, start int, end int) []Ticket {
	if start < 0 || end < 0 {
		log.Fatalf("Page start %v and end %v error, negative values for slice", start, end)
	}
	if start > end {
		log.Printf("Start %v is greater than end %v", start, end)
		return ticket
	}
	return ticket[start:end]
}

func updateEnd(currentEnd int, sizeOfSlice int) int {
	overflow := currentEnd + MAX_PAGE_SIZE
	if overflow > sizeOfSlice {
		return sizeOfSlice % currentEnd
	} else {
		return MAX_PAGE_SIZE
	}
}
func (p *Page) SelectOne(index int) bool {
	p.Select = index
	last := index - 1
	if last >= 0 && last <= len(p.FullSlice)-1 {
		p.SelectedTicket = p.FullSlice[last]
		return true
	}
	return false

}
