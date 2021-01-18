// Package storage -
package storage

import "reflect"

// Store - interface that defines methods required to access stored data
type Store interface {
	GetGroup(string) ([]map[string]interface{}, error)
	GetGroupNames() []string
	GetTerms(string) (map[string]reflect.Type, error)
}
