package repo

import "fmt"

// FindMatches -
func (d *Data) FindMatches(group, term, value string) ([]map[string]string, error) {
	return nil, fmt.Errorf("%s does not have an index", group)
}

func (d Data) ToDTO(interface{}) ([]map[string]string, error) {
	return nil, nil
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
