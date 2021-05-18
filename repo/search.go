package repo

import "strings"

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
