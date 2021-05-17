package repo

import (
	"fmt"
	"strconv"
	"strings"
)

// Organisation -
type Organisation struct {
	ID            int      `json:"_id"`
	URL           string   `json:"url"`
	ExternalID    string   `json:"external_id"`
	Name          string   `json:"name"`
	DomainNames   []string `json:"domain_names"`
	CreatedAt     string   `json:"created_at"`
	Details       string   `json:"details"`
	SharedTickets bool     `json:"shared_tickets"`
	Tags          []string `json:"tags"`
}

// OrganisationIndexes -
func (d *Data) OrganisationIndexes() {
	// map[fieldname]map[fieldvalue][]*Organisation
	d.OrgIdx = make(map[string]map[string][]*Organisation)
	d.OrgIdx["_id"] = make(map[string][]*Organisation)
	d.OrgIdx["url"] = make(map[string][]*Organisation)
	d.OrgIdx["external_id"] = make(map[string][]*Organisation)
	d.OrgIdx["name"] = make(map[string][]*Organisation)
	d.OrgIdx["domain_names"] = make(map[string][]*Organisation)
	d.OrgIdx["created_at"] = make(map[string][]*Organisation)
	d.OrgIdx["details"] = make(map[string][]*Organisation)
	d.OrgIdx["shared_tickets"] = make(map[string][]*Organisation)
	d.OrgIdx["tags"] = make(map[string][]*Organisation)
	d.Terms["organisations"] = map[string]struct{}{
		"_id":            struct{}{},
		"url":            struct{}{},
		"external_id":    struct{}{},
		"name":           struct{}{},
		"domain_names":   struct{}{},
		"created_at":     struct{}{},
		"details":        struct{}{},
		"shared_tickets": struct{}{},
		"tags":           struct{}{},
	}

	for i := range d.Organisations {
		d.OrgIdx["_id"][fmt.Sprintf("%d", d.Organisations[i].ID)] = append(d.OrgIdx["_id"][fmt.Sprintf("%d", d.Organisations[i].ID)], d.Organisations[i])
		d.OrgIdx["url"][strings.ToLower(d.Organisations[i].URL)] = append(d.OrgIdx["url"][strings.ToLower(d.Organisations[i].URL)], d.Organisations[i])
		d.OrgIdx["external_id"][strings.ToLower(d.Organisations[i].ExternalID)] = append(d.OrgIdx["external_id"][strings.ToLower(d.Organisations[i].ExternalID)], d.Organisations[i])
		d.OrgIdx["name"][strings.ToLower(d.Organisations[i].Name)] = append(d.OrgIdx["name"][strings.ToLower(d.Organisations[i].Name)], d.Organisations[i])
		for _, domainName := range d.Organisations[i].DomainNames {
			d.OrgIdx["domain_names"][strings.ToLower(domainName)] = append(d.OrgIdx["domain_names"][strings.ToLower(domainName)], d.Organisations[i])
		}
		d.OrgIdx["created_at"][strings.ToLower(d.Organisations[i].CreatedAt)] = append(d.OrgIdx["created_at"][strings.ToLower(d.Organisations[i].CreatedAt)], d.Organisations[i])
		d.OrgIdx["details"][strings.ToLower(d.Organisations[i].Details)] = append(d.OrgIdx["details"][strings.ToLower(d.Organisations[i].Details)], d.Organisations[i])
		d.OrgIdx["shared_tickets"][strconv.FormatBool(d.Organisations[i].SharedTickets)] = append(d.OrgIdx["shared_tickets"][strconv.FormatBool(d.Organisations[i].SharedTickets)], d.Organisations[i])
		for _, tag := range d.Organisations[i].Tags {
			d.OrgIdx["tags"][strings.ToLower(tag)] = append(d.OrgIdx["tags"][strings.ToLower(tag)], d.Organisations[i])
		}
	}
}
