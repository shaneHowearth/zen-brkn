package repo

import (
	"fmt"
	"strconv"
	"strings"
)

// Ticket -
type Ticket struct {
	ID             string   `json:"_id"`
	URL            string   `json:"url"`
	ExternalID     string   `json:"external_id"`
	CreatedAt      string   `json:"created_at"`
	Type           string   `json:"type"`
	Subject        string   `json:"subject"`
	Description    string   `json:"description"`
	Priority       string   `json:"priority"`
	Status         string   `json:"status"`
	SubmitterID    int      `json:"submitter_id"`
	AssigneeID     int      `json:"assignee_id"`
	OrganizationID int      `json:"organization_id"`
	Tags           []string `json:"tags"`
	HasIncidents   bool     `json:"has_incidents"`
	DueAt          string   `json:"due_at"`
	Via            string   `json:"via"`
}

// TicketIndexes -
func (d *Data) TicketIndexes() {
	// map[fieldname]map[fieldvalue][]*Ticket
	d.TicketIdx = make(map[string]map[string][]*Ticket)
	d.TicketIdx["_id"] = make(map[string][]*Ticket)
	d.TicketIdx["url"] = make(map[string][]*Ticket)
	d.TicketIdx["external_id"] = make(map[string][]*Ticket)
	d.TicketIdx["created_at"] = make(map[string][]*Ticket)
	d.TicketIdx["type"] = make(map[string][]*Ticket)
	d.TicketIdx["subject"] = make(map[string][]*Ticket)
	d.TicketIdx["description"] = make(map[string][]*Ticket)
	d.TicketIdx["priority"] = make(map[string][]*Ticket)
	d.TicketIdx["status"] = make(map[string][]*Ticket)
	d.TicketIdx["submitter_id"] = make(map[string][]*Ticket)
	d.TicketIdx["assignee_id"] = make(map[string][]*Ticket)
	d.TicketIdx["organization_id"] = make(map[string][]*Ticket)
	d.TicketIdx["tags"] = make(map[string][]*Ticket)
	d.TicketIdx["has_incidents"] = make(map[string][]*Ticket)
	d.TicketIdx["due_at"] = make(map[string][]*Ticket)
	d.TicketIdx["via"] = make(map[string][]*Ticket)
	d.Terms["tickets"] = map[string]struct{}{
		"_id":             struct{}{},
		"url":             struct{}{},
		"external_id":     struct{}{},
		"created_at":      struct{}{},
		"type":            struct{}{},
		"subject":         struct{}{},
		"description":     struct{}{},
		"priority":        struct{}{},
		"status":          struct{}{},
		"submitter_id":    struct{}{},
		"assignee_id":     struct{}{},
		"organization_id": struct{}{},
		"tags":            struct{}{},
		"has_incidents":   struct{}{},
		"due_at":          struct{}{},
		"via":             struct{}{},
	}
	for i := range d.Tickets {
		d.TicketIdx["_id"][strings.ToLower(d.Tickets[i].ID)] = append(d.TicketIdx["_id"][strings.ToLower(d.Tickets[i].ID)], d.Tickets[i])
		d.TicketIdx["url"][strings.ToLower(d.Tickets[i].URL)] = append(d.TicketIdx["url"][strings.ToLower(d.Tickets[i].URL)], d.Tickets[i])
		d.TicketIdx["external_id"][strings.ToLower(d.Tickets[i].ExternalID)] = append(d.TicketIdx["external_id"][strings.ToLower(d.Tickets[i].ExternalID)], d.Tickets[i])
		d.TicketIdx["created_at"][strings.ToLower(d.Tickets[i].CreatedAt)] = append(d.TicketIdx["created_at"][strings.ToLower(d.Tickets[i].CreatedAt)], d.Tickets[i])
		d.TicketIdx["type"][strings.ToLower(d.Tickets[i].Type)] = append(d.TicketIdx["type"][strings.ToLower(d.Tickets[i].Type)], d.Tickets[i])
		d.TicketIdx["subject"][strings.ToLower(d.Tickets[i].Subject)] = append(d.TicketIdx["subject"][strings.ToLower(d.Tickets[i].Subject)], d.Tickets[i])
		d.TicketIdx["description"][strings.ToLower(d.Tickets[i].Description)] = append(d.TicketIdx["description"][strings.ToLower(d.Tickets[i].Description)], d.Tickets[i])
		d.TicketIdx["priority"][strings.ToLower(d.Tickets[i].Priority)] = append(d.TicketIdx["priority"][strings.ToLower(d.Tickets[i].Priority)], d.Tickets[i])
		d.TicketIdx["status"][strings.ToLower(d.Tickets[i].Status)] = append(d.TicketIdx["status"][strings.ToLower(d.Tickets[i].Status)], d.Tickets[i])
		d.TicketIdx["submitter_id"][fmt.Sprintf("%d", d.Tickets[i].SubmitterID)] = append(d.TicketIdx["submitter_id"][fmt.Sprintf("%d", d.Tickets[i].SubmitterID)], d.Tickets[i])
		d.TicketIdx["assignee_id"][fmt.Sprintf("%d", d.Tickets[i].AssigneeID)] = append(d.TicketIdx["assignee_id"][fmt.Sprintf("%d", d.Tickets[i].AssigneeID)], d.Tickets[i])
		d.TicketIdx["organization_id"][fmt.Sprintf("%d", d.Tickets[i].OrganizationID)] = append(d.TicketIdx["organization_id"][fmt.Sprintf("%d", d.Tickets[i].OrganizationID)], d.Tickets[i])
		for _, tag := range d.Tickets[i].Tags {
			d.TicketIdx["tags"][strings.ToLower(tag)] = append(d.TicketIdx["tags"][strings.ToLower(tag)], d.Tickets[i])
		}
		d.TicketIdx["has_incidents"][strconv.FormatBool(d.Tickets[i].HasIncidents)] = append(d.TicketIdx["has_incidents"][strconv.FormatBool(d.Tickets[i].HasIncidents)], d.Tickets[i])
		d.TicketIdx["due_at"][strings.ToLower(d.Tickets[i].DueAt)] = append(d.TicketIdx["due_at"][strings.ToLower(d.Tickets[i].DueAt)], d.Tickets[i])
		d.TicketIdx["via"][strings.ToLower(d.Tickets[i].Via)] = append(d.TicketIdx["via"][strings.ToLower(d.Tickets[i].Via)], d.Tickets[i])
	}
}
