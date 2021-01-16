// Package simple -
package simple

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

// Simple -
type Simple struct{}

// Contains -
func (s Simple) Contains(search, field string, data []map[string]interface{}) (map[string]interface{}, error) {
	results := map[string]interface{}{}
	search = strings.TrimSpace(search)
	switch data[0][field].(type) {
	case bool:
		bSearch := false
		switch search {
		case "t", "true":
			bSearch = true
		case "f", "false":
			bSearch = false
		default:
			return nil, fmt.Errorf("cannot convert %s to bool", search)
		}
		for i := range data {
			// Will only return the first one
			if data[i][field].(bool) == bSearch {
				return data[i], nil
			}
		}
	case string:
		for i := range data {
			if strings.Contains(data[i][field].(string), search) {
				return data[i], nil
			}
		}
	case int:
		i, err := strconv.Atoi(search)
		if err != nil {
			return nil, fmt.Errorf("cannont convert %s with error %v", search, err)
		}
		for j := range data {
			if data[j][field].(int) == i {
				return data[j], nil
			}
		}
	case []string:
		for i := range data {
			for ss := range (data[i][field]).([]string) {
				if strings.Contains(((data[i][field]).([]string))[ss], search) {
					return data[i], nil
				}
			}
		}
	default:
		return nil, fmt.Errorf("unknown type %s", reflect.TypeOf(data[0][field]))
	}
	// Nothing found
	return results, nil
}
