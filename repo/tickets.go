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

// ToDTO - TODO
func (t *Ticket) ToDTO() map[string][]string {
	m := map[string][]string{}
	m["_id"] = []string{t.ID}
	m["url"] = []string{t.URL}
	m["external_id"] = []string{t.ExternalID}
	m["created_at"] = []string{t.CreatedAt}
	m["type"] = []string{t.Type}
	m["subject"] = []string{t.Subject}
	m["description"] = []string{t.Description}
	m["priority"] = []string{t.Priority}
	m["status"] = []string{t.Status}
	m["submitter_id"] = []string{fmt.Sprintf("%d", t.SubmitterID)}
	m["assignee_id"] = []string{fmt.Sprintf("%d", t.AssigneeID)}
	m["organization_id"] = []string{fmt.Sprintf("%d", t.OrganizationID)}
	m["tags"] = t.Tags
	m["has_incidents"] = []string{strconv.FormatBool(t.HasIncidents)}
	m["due_at"] = []string{t.DueAt}
	m["via"] = []string{t.Via}

	return m
}

// CreateIndex -
func (t *Ticket) CreateIndex(in interface{}, name string) map[string]map[string][]item {
	d := in.([]*Ticket)
	// TicketIndexes -
	// func (d *Data) TicketIndexes() {
	// map[fieldname]map[fieldvalue][]*Ticket
	m := make(map[string]map[string][]item)
	m["_id"] = make(map[string][]item)
	m["url"] = make(map[string][]item)
	m["external_id"] = make(map[string][]item)
	m["created_at"] = make(map[string][]item)
	m["type"] = make(map[string][]item)
	m["subject"] = make(map[string][]item)
	m["description"] = make(map[string][]item)
	m["priority"] = make(map[string][]item)
	m["status"] = make(map[string][]item)
	m["submitter_id"] = make(map[string][]item)
	m["assignee_id"] = make(map[string][]item)
	m["organization_id"] = make(map[string][]item)
	m["tags"] = make(map[string][]item)
	m["has_incidents"] = make(map[string][]item)
	m["due_at"] = make(map[string][]item)
	m["via"] = make(map[string][]item)
	for i := range d {
		m["_id"][strings.ToLower(d[i].ID)] = append(m["_id"][strings.ToLower(d[i].ID)], d[i])
		m["url"][strings.ToLower(d[i].URL)] = append(m["url"][strings.ToLower(d[i].URL)], d[i])
		m["external_id"][strings.ToLower(d[i].ExternalID)] = append(m["external_id"][strings.ToLower(d[i].ExternalID)], d[i])
		m["created_at"][strings.ToLower(d[i].CreatedAt)] = append(m["created_at"][strings.ToLower(d[i].CreatedAt)], d[i])
		m["type"][strings.ToLower(d[i].Type)] = append(m["type"][strings.ToLower(d[i].Type)], d[i])
		m["subject"][strings.ToLower(d[i].Subject)] = append(m["subject"][strings.ToLower(d[i].Subject)], d[i])
		m["description"][strings.ToLower(d[i].Description)] = append(m["description"][strings.ToLower(d[i].Description)], d[i])
		m["priority"][strings.ToLower(d[i].Priority)] = append(m["priority"][strings.ToLower(d[i].Priority)], d[i])
		m["status"][strings.ToLower(d[i].Status)] = append(m["status"][strings.ToLower(d[i].Status)], d[i])
		m["submitter_id"][fmt.Sprintf("%d", d[i].SubmitterID)] = append(m["submitter_id"][fmt.Sprintf("%d", d[i].SubmitterID)], d[i])
		m["assignee_id"][fmt.Sprintf("%d", d[i].AssigneeID)] = append(m["assignee_id"][fmt.Sprintf("%d", d[i].AssigneeID)], d[i])
		m["organization_id"][fmt.Sprintf("%d", d[i].OrganizationID)] = append(m["organization_id"][fmt.Sprintf("%d", d[i].OrganizationID)], d[i])
		for _, tag := range d[i].Tags {
			m["tags"][strings.ToLower(tag)] = append(m["tags"][strings.ToLower(tag)], d[i])
		}
		m["has_incidents"][strconv.FormatBool(d[i].HasIncidents)] = append(m["has_incidents"][strconv.FormatBool(d[i].HasIncidents)], d[i])
		m["due_at"][strings.ToLower(d[i].DueAt)] = append(m["due_at"][strings.ToLower(d[i].DueAt)], d[i])
		m["via"][strings.ToLower(d[i].Via)] = append(m["via"][strings.ToLower(d[i].Via)], d[i])
	}
	return m
}
