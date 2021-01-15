// Package simple -
package simple

import "strings"

// Simple -
type Simple struct{}

// Contains -
func (s Simple) Contains(search string, data []string) ([]string, error) {
	results := []string{}
	for i := range data {
		if strings.Contains(data[i], search) {
			results = append(results, data[i])
		}
	}
	return results, nil
}
