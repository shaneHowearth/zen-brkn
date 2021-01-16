// Package storage -
package storage

import "reflect"

// Store - interface that defines methods required to access stored data
type Store interface {
	LoadData() error
	GetGroup(string) ([]map[string]interface{}, error)
	GetTerms(string) (map[string]reflect.Type, error)
}
