package repo

import "fmt"

// FindMatches -
func (d *Data) FindMatches(group, term, value string) ([]map[string]string, error) {
	// need to signal which index to use, and which ones exist
	// Hardcoding :(
	switch group {
	case "organisation":
		return d.ToDTO(d.OrgIdx[term][value]), nil
	case "ticket":
		return d.ToDTO(d.TicketIdx[term][value]), nil
	case "user":
		return d.ToDTO(d.UserIdx[term][value]), nil
	}
	return nil, fmt.Errorf("%s does not have an index", group)
}

func (d Data) ToDTO(interface{}) ([]map[string]string, error)

// FindRelated -
func (d *Data) FindRelated(map[string]string) ([]map[string]string, error) {
	return nil, nil
}

// GetGroups -
func (d *Data) GetGroups() ([]string, error) {
	return d.Groups, nil
}

// GetTerms -
func (d *Data) GetTerms(group string) (map[string]struct{}, error) {
	return d.Terms[group], nil
}
