package main

import (
	"log"
	"math"
	"time"
)

// ------------------------Json data representation ----------------------------------------------------------
type TicketResponse struct {
	Tickets []Ticket `json:"tickets"`
}

type Ticket struct {
	AllowAttachments bool      `json:"allow_attachments,omitempty"`
	AllowChannelBack bool      `json:"allow_channelback,omitempty"`
	AssigneeId       int       `json:"assignee_id,omitempty"`
	BrandId          int       `json:"brand_id,omitempty"`
	CollaboratorIds  []int     `json:"collaborator_ids,omitempty"`
	Collaborators    []string  `json:"collaborators,omitempty"`
	CreatedAt        time.Time `json:"created_at,omitempty"`
	CustomFields     []string  `json:"custom_fields,omitempty"`
	Description      string    `json:"description"`
	DueAt            string    `json:"due_at,omitempty"`
	EmailCCIds       []int     `json:"email_cc_ids,omitempty"`
	ExternalId       string    `json:"external_id,omitempty"`
	FollowerIds      []string  `json:"follower_ids,omitempty"`
	FollowupIds      []string  `json:"followup_ids,omitempty"`
	ForumTopicId     int       `json:"forum_topic_id,omitempty"`
	GroupId          int       `json:"group_id,omitempty"`
	HasIncidents     bool      `json:"has_incidents,omitempty"`
	Id               int       `json:"id,omitempty"`
	IsPublic         bool      `json:"is_public,omitempty"`
	MacroIds         int       `json:"macro_ids,omitempty"`
	OrganizationId   int       `json:"organization_id,omitempty"`
	Priority         string    `json:"priority,omitempty"`
	ProblemId        int       `json:"problem_id,omitempty"`
	RawSubject       string    `json:"raw_subject"`
	Receipient       string    `json:"receipient,omitempty"`
	RequesterId      int       `json:"requester_id"`
	// Suppose to be an object
	SatisfactionRating  SatisfactionRating `json:"satisfaction_rating,omitempty"`
	SharingAgreementIds []int              `json:"sharing_agreement_ids,omitempty"`
	Status              string             `json:"status,omitempty"`
	Subject             string             `json:"subject"`
	SubmitterId         int                `json:"submitter_id,omitempty"`
	Tags                []string           `json:"tags"`
	TicketFormId        int                `json:"ticket_form_id,omitempty"`
	Type                string             `json:"type,omitempty"`
	UpdatedAt           time.Time          `json:"updated_at,omitempty"`
	Url                 string             `json:"url,omitempty"`
	// Suppose to be an object
	// Via                 Via `json:",omitempty"`
	ViaFollowupSourceId int `json:"via_followup_source_id,omitempty"`
}

type Via struct {
	Channel string `json:",omitempty"`
	Source  string `json:",omitempty"`
}

type SatisfactionRating struct {
	AssigneeId  int    `json:"assignee_id"`
	Comment     string `json:",omitempty"`
	CreatedAt   string `json:"created_at,omitempty"`
	GroupId     int    `json:"group_id"`
	Id          int    `json:",omitempty"`
	Reason      string `json:",omitempty"`
	ReasonCode  int    `json:"reason_code,omitempty"`
	RequesterId int    `json:"requester_id"`
	Score       string `json:"score"`
	TicketId    int    `json:"ticket_id"`
	UpdatedAt   string `json:"updated_at,omitempty"`
	Url         string `json:",omitempty"`
}

// -------------------------------------------------------------------------------------------------------------------

// --------------------------------- Represents Page and update struct ---------------------------------------
type Page struct {
	MaxPageSize  int
	CurrentPage  int
	CanGoBack    bool
	CanGoForward bool
	Start        int
	End          int
	CurrentSlice []Ticket
	FullSlice    []Ticket
	All          bool
	Select       int
}

func NewPage(t []Ticket) *Page {
	max := findMaxPage(len(t))
	return &Page{
		MaxPageSize:  max,
		CurrentPage:  1,
		CanGoBack:    false,
		CanGoForward: true,
		Start:        0,
		End:          MAX_PAGE_SIZE,
		CurrentSlice: t[0:MAX_PAGE_SIZE],
		FullSlice:    t,
		All:          true,
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

func updateEnd(end int, size int) int {
	temp := end + MAX_PAGE_SIZE
	if temp > size {
		return temp % size
	} else {
		return MAX_PAGE_SIZE
	}
}

// -------------------------------------- Represents menu ------------------------------------

type Menu struct {
	Main    bool
	ContAll bool
	ViewOne bool
	Opt1    string
	Opt2    string
	Opt3    string
}

func CreateMainMenu() Menu {
	return Menu{
		Main: true,
		Opt1: "1)\tList all Tickets",
		Opt2: "2)\tView a ticket",
		Opt3: "3)\tExit",
	}
}

func CreateViewAllMenu() Menu {
	return Menu{
		ContAll: true,
		Opt1:    "1)\tNext Page",
		Opt2:    "2)\tPrevious Page",
		Opt3:    "3)\tReturn to main menu",
	}
}

func CreateSelectMenu() Menu {
	return Menu{
		ViewOne: true,
		Opt1:    "1)\tSelect another ticket",
		Opt3:    "3)\tReturn to main menu",
	}
}

// -------------------------------------------------------------------------------------------------------------------
