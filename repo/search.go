package repo

import (
	"fmt"
	"strings"
)

// Note: Search of the data is closely coupled to how the data is stored, that
// is, if the data is stored as a BST then search operations are going to be
// different than search in a Off the shelf datastore like Redis or MongoDB, or
// even, the search of a B+ tree.

// FindMatches -
func (d *Data) FindMatches(group, term, value string) ([]map[string][]string, error) {
	// Case insensitive
	value = strings.ToLower(value)
	matches := []map[string][]string{}
	for i := range d.Indexes[group][term][value] {
		matches = append(matches, d.Indexes[group][term][value][i].ToDTO())
	}
	return matches, nil
}

// FindRelated -
func (d *Data) FindRelated(group string, origin map[string][]string) (map[string][]map[string][]string, error) {
	// Where the data exists, values from any related entities should be
	// included in the results, i.e. searching organization by id should return
	// its tickets and users.
	// Tickets
	// OrganizationID int      `json:"organization_id"`
	// SubmitterID    int      `json:"submitter_id"` - UserId?
	// Users
	// OrganizationID int      `json:"organization_id"`

	// We're locked to the data on what a relationship is.
	matches := map[string][]map[string][]string{}
	switch group {
	case "organisation":
		tickets := d.Indexes["tickets"]["organization_id"][origin["_id"][0]]
		users := d.Indexes["users"]["organization_id"][origin["_id"][0]]
		matches["tickets"] = []map[string][]string{}
		for i := range tickets {
			matches["tickets"] = append(matches["tickets"], tickets[i].ToDTO())
		}
		matches["users"] = []map[string][]string{}
		for i := range users {
			matches["users"] = append(matches["users"], users[i].ToDTO())
		}
	case "ticket":
		organisation := d.Indexes["organisation"]["_id"][origin["organization_id"][0]]
		users := d.Indexes["users"]["_id"][origin["submitter_id"][0]]
		matches["organisations"] = []map[string][]string{}
		for i := range organisation {
			matches["organisations"] = append(matches["organisations"], organisation[i].ToDTO())
		}
		matches["users"] = []map[string][]string{}
		for i := range users {
			matches["users"] = append(matches["users"], users[i].ToDTO())
		}
	case "user":
		organisation := d.Indexes["organisation"]["_id"][origin["organization_id"][0]]
		tickets := d.Indexes["ticket"]["submitter_id"][origin["_id"][0]]
		matches["organisations"] = []map[string][]string{}
		for i := range organisation {
			matches["organisations"] = append(matches["organisations"], organisation[i].ToDTO())
		}
		matches["tickets"] = []map[string][]string{}
		for i := range tickets {
			matches["tickets"] = append(matches["tickets"], tickets[i].ToDTO())
		}
	default:
		return nil, fmt.Errorf("%s has not been implemented yet", group)
	}
	return matches, nil
}

// GetGroups -
func (d *Data) GetGroups() ([]string, error) {
	return d.Groups, nil
}

// GetTerms -
func (d *Data) GetTerms(group string) (map[string]struct{}, error) {
	return d.Terms[group], nil
}
