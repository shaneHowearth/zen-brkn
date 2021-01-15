// Package search -
package search

// Search -
type Search interface {
	Contains(substr string, data []string) ([]string, error)
}
