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
	d.OrgIdx["ID"] = make(map[string][]*Organisation)
	d.OrgIdx["URL"] = make(map[string][]*Organisation)
	d.OrgIdx["ExternalID"] = make(map[string][]*Organisation)
	d.OrgIdx["Name"] = make(map[string][]*Organisation)
	d.OrgIdx["DomainNames"] = make(map[string][]*Organisation)
	d.OrgIdx["CreatedAt"] = make(map[string][]*Organisation)
	d.OrgIdx["Details"] = make(map[string][]*Organisation)
	d.OrgIdx["SharedTickets"] = make(map[string][]*Organisation)
	d.OrgIdx["Tags"] = make(map[string][]*Organisation)

	for i := range d.Organisations {
		d.OrgIdx["ID"][fmt.Sprintf("%d", d.Organisations[i].ID)] = append(d.OrgIdx["ID"][fmt.Sprintf("%d", d.Organisations[i].ID)], d.Organisations[i])
		d.OrgIdx["URL"][strings.ToLower(d.Organisations[i].URL)] = append(d.OrgIdx["URL"][strings.ToLower(d.Organisations[i].URL)], d.Organisations[i])
		d.OrgIdx["ExternalID"][strings.ToLower(d.Organisations[i].ExternalID)] = append(d.OrgIdx["ExternalID"][strings.ToLower(d.Organisations[i].ExternalID)], d.Organisations[i])
		d.OrgIdx["Name"][strings.ToLower(d.Organisations[i].Name)] = append(d.OrgIdx["Name"][strings.ToLower(d.Organisations[i].Name)], d.Organisations[i])
		for _, domainName := range d.Organisations[i].DomainNames {
			d.OrgIdx["DomainNames"][strings.ToLower(domainName)] = append(d.OrgIdx["DomainNames"][strings.ToLower(domainName)], d.Organisations[i])
		}
		d.OrgIdx["CreatedAt"][strings.ToLower(d.Organisations[i].CreatedAt)] = append(d.OrgIdx["CreatedAt"][strings.ToLower(d.Organisations[i].CreatedAt)], d.Organisations[i])
		d.OrgIdx["Details"][strings.ToLower(d.Organisations[i].Details)] = append(d.OrgIdx["Details"][strings.ToLower(d.Organisations[i].Details)], d.Organisations[i])
		d.OrgIdx["SharedTickets"][strconv.FormatBool(d.Organisations[i].SharedTickets)] = append(d.OrgIdx["SharedTickets"][strconv.FormatBool(d.Organisations[i].SharedTickets)], d.Organisations[i])
		for _, tag := range d.Organisations[i].Tags {
			d.OrgIdx["Tags"][strings.ToLower(tag)] = append(d.OrgIdx["Tags"][strings.ToLower(tag)], d.Organisations[i])
		}
	}
}
