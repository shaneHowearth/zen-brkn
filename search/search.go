// Package search -
package search

// Search -
type Search interface {
	Contains(substr, field string, data []map[string]interface{}) (map[string]interface{}, error)
}
