// Package simple -
package simple

import (
	"fmt"
	"math"
	"reflect"
	"strconv"
	"strings"
)

// Simple -
type Simple struct{}

const epsilon = 1e-9

func nearlyEqualFloats(a, b float64) bool {
	absA := math.Abs(a)
	absB := math.Abs(b)
	diff := math.Abs(a - b)

	if a == b { // shortcut, handles infinities
		return true
	} else if a == 0 || b == 0 || (absA+absB < math.SmallestNonzeroFloat64) {
		// a or b is zero or both are extremely close to it relative error is
		// less meaningful here
		return diff < (epsilon * math.SmallestNonzeroFloat64)

	} else { // use relative error
		return diff/math.Min((absA+absB), math.MaxFloat64) < epsilon
	}

}

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
			return nil, fmt.Errorf("cannot use %s for comparison to boolean field", search)
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
	case float32, float64:
		f64, err := strconv.ParseFloat(search, 64)
		if err != nil {
			return nil, fmt.Errorf("cannot use %s for comparison to float field", search)
		}
		for j := range data {
			if nearlyEqualFloats(data[j][field].(float64), f64) {
				return data[j], nil
			}
		}
	case int:
		i, err := strconv.Atoi(search)
		if err != nil {
			return nil, fmt.Errorf("cannot use %s for comparison to int field", search)
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
