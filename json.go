package main

import (
	"time"
)

// ------------------------Json data representation in structs ----------------------------------------------------------
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
