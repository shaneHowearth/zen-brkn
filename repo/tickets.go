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
	d.TicketIdx["ID"] = make(map[string][]*Ticket)
	d.TicketIdx["URL"] = make(map[string][]*Ticket)
	d.TicketIdx["ExternalID"] = make(map[string][]*Ticket)
	d.TicketIdx["CreatedAt"] = make(map[string][]*Ticket)
	d.TicketIdx["Type"] = make(map[string][]*Ticket)
	d.TicketIdx["Subject"] = make(map[string][]*Ticket)
	d.TicketIdx["Description"] = make(map[string][]*Ticket)
	d.TicketIdx["Priority"] = make(map[string][]*Ticket)
	d.TicketIdx["Status"] = make(map[string][]*Ticket)
	d.TicketIdx["SubmitterID"] = make(map[string][]*Ticket)
	d.TicketIdx["AssigneeID"] = make(map[string][]*Ticket)
	d.TicketIdx["OrganizationID"] = make(map[string][]*Ticket)
	d.TicketIdx["Tags"] = make(map[string][]*Ticket)
	d.TicketIdx["HasIncidents"] = make(map[string][]*Ticket)
	d.TicketIdx["DueAt"] = make(map[string][]*Ticket)
	d.TicketIdx["Via"] = make(map[string][]*Ticket)
	for i := range d.Tickets {
		d.TicketIdx["ID"][strings.ToLower(d.Tickets[i].ID)] = append(d.TicketIdx["ID"][strings.ToLower(d.Tickets[i].ID)], d.Tickets[i])
		d.TicketIdx["URL"][strings.ToLower(d.Tickets[i].URL)] = append(d.TicketIdx["URL"][strings.ToLower(d.Tickets[i].URL)], d.Tickets[i])
		d.TicketIdx["ExternalID"][strings.ToLower(d.Tickets[i].ExternalID)] = append(d.TicketIdx["ExternalID"][strings.ToLower(d.Tickets[i].ExternalID)], d.Tickets[i])
		d.TicketIdx["CreatedAt"][strings.ToLower(d.Tickets[i].CreatedAt)] = append(d.TicketIdx["CreatedAt"][strings.ToLower(d.Tickets[i].CreatedAt)], d.Tickets[i])
		d.TicketIdx["Type"][strings.ToLower(d.Tickets[i].Type)] = append(d.TicketIdx["Type"][strings.ToLower(d.Tickets[i].Type)], d.Tickets[i])
		d.TicketIdx["Subject"][strings.ToLower(d.Tickets[i].Subject)] = append(d.TicketIdx["Subject"][strings.ToLower(d.Tickets[i].Subject)], d.Tickets[i])
		d.TicketIdx["Description"][strings.ToLower(d.Tickets[i].Description)] = append(d.TicketIdx["Description"][strings.ToLower(d.Tickets[i].Description)], d.Tickets[i])
		d.TicketIdx["Priority"][strings.ToLower(d.Tickets[i].Priority)] = append(d.TicketIdx["Priority"][strings.ToLower(d.Tickets[i].Priority)], d.Tickets[i])
		d.TicketIdx["Status"][strings.ToLower(d.Tickets[i].Status)] = append(d.TicketIdx["Status"][strings.ToLower(d.Tickets[i].Status)], d.Tickets[i])
		d.TicketIdx["SubmitterID"][fmt.Sprintf("%d", d.Tickets[i].SubmitterID)] = append(d.TicketIdx["SubmitterID"][fmt.Sprintf("%d", d.Tickets[i].SubmitterID)], d.Tickets[i])
		d.TicketIdx["AssigneeID"][fmt.Sprintf("%d", d.Tickets[i].AssigneeID)] = append(d.TicketIdx["AssigneeID"][fmt.Sprintf("%d", d.Tickets[i].AssigneeID)], d.Tickets[i])
		d.TicketIdx["OrganizationID"][fmt.Sprintf("%d", d.Tickets[i].OrganizationID)] = append(d.TicketIdx["OrganizationID"][fmt.Sprintf("%d", d.Tickets[i].OrganizationID)], d.Tickets[i])
		for _, tag := range d.Tickets[i].Tags {
			d.TicketIdx["Tags"][strings.ToLower(tag)] = append(d.TicketIdx["Tags"][strings.ToLower(tag)], d.Tickets[i])
		}
		d.TicketIdx["HasIncidents"][strconv.FormatBool(d.Tickets[i].HasIncidents)] = append(d.TicketIdx["HasIncidents"][strconv.FormatBool(d.Tickets[i].HasIncidents)], d.Tickets[i])
		d.TicketIdx["DueAt"][strings.ToLower(d.Tickets[i].DueAt)] = append(d.TicketIdx["DueAt"][strings.ToLower(d.Tickets[i].DueAt)], d.Tickets[i])
		d.TicketIdx["Via"][strings.ToLower(d.Tickets[i].Via)] = append(d.TicketIdx["Via"][strings.ToLower(d.Tickets[i].Via)], d.Tickets[i])
	}
}
