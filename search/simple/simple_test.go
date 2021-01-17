package simple_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/shanehowearth/zen/search/simple"
	"github.com/stretchr/testify/assert"
)

func TestContains(t *testing.T) {
	s := simple.Searcher{}
	testcases := map[string]struct {
		search string
		field  string
		data   []map[string]interface{}
		output map[string]interface{}
		err    error
	}{
		"False Boolean": {search: "f", field: "boolean",
			data:   []map[string]interface{}{{"boolean": false}, {"boolean": true}},
			output: map[string]interface{}{"boolean": false},
		},
		"True Boolean": {
			search: "true", field: "boolean",
			data:   []map[string]interface{}{{"boolean": false}, {"boolean": true}},
			output: map[string]interface{}{"boolean": true},
		},
		"Bad Boolean": {
			search: "badInput", field: "boolean",
			data:   []map[string]interface{}{{"boolean": false}, {"boolean": true}},
			output: nil,
			err:    fmt.Errorf("cannot use %s for comparison to boolean field", "badInput"),
		},
		"No Boolean Match": {
			search: "false", field: "boolean",
			data:   []map[string]interface{}{{"boolean": true}, {"boolean": true}},
			output: map[string]interface{}{},
		},
		"String": {
			search: "exists", field: "string",
			data:   []map[string]interface{}{{"string": "Not a match"}, {"string": "match exists"}},
			output: map[string]interface{}{"string": "match exists"},
		},
		"No String Match": {
			search: "false", field: "string",
			data:   []map[string]interface{}{{"string": "No match here"}, {"string": "Not a match"}},
			output: map[string]interface{}{},
		},
		"Int": {
			search: "100", field: "int",
			data:   []map[string]interface{}{{"int": 100}, {"int": 105}},
			output: map[string]interface{}{"int": 100},
		},
		"No Int Match": {
			search: "110", field: "int",
			data:   []map[string]interface{}{{"int": 100}, {"int": 105}},
			output: map[string]interface{}{},
		},
		"Bad Int Match": {
			search: "11a", field: "int",
			data:   []map[string]interface{}{{"int": 100}, {"int": 105}},
			output: nil,
			err:    fmt.Errorf("cannot use %s for comparison to int field", "11a"),
		},
		"Slice of String": {
			search: "exists", field: "[]string",
			data:   []map[string]interface{}{{"[]string": []string{"Not", "a match"}}, {"[]string": []string{"match", "exists"}}},
			output: map[string]interface{}{"[]string": []string{"match", "exists"}},
		},
		"No Slice of String Match": {
			search: "false", field: "[]string",
			data:   []map[string]interface{}{{"[]string": []string{"Not", "a match"}}, {"[]string": []string{"match", "exists"}}},
			output: map[string]interface{}{},
		},
		"Float": {
			search: "100.3", field: "float32",
			data:   []map[string]interface{}{{"float32": 100.3}, {"float32": 105.7}},
			output: map[string]interface{}{"float32": 100.3},
		},
		"No Float Match": {
			search: "110.5", field: "float32",
			data:   []map[string]interface{}{{"float32": 100.6}, {"float32": 105.7}},
			output: map[string]interface{}{},
		},
		"Bad Float Match": {
			search: "11a", field: "float32",
			data:   []map[string]interface{}{{"float32": 100.2}, {"float32": 105.1}},
			output: nil,
			err:    fmt.Errorf("cannot use %s for comparison to float field", "11a"),
		},
		"Zero Float": {
			search: "10.3", field: "float32",
			data:   []map[string]interface{}{{"float32": 0.0}, {"float32": 10.3}},
			output: map[string]interface{}{"float32": 10.3},
		},
		"Default": {
			search: "a", field: "rune",
			data:   []map[string]interface{}{{"rune": 'a'}, {"rune": 'b'}},
			output: nil,
			err:    fmt.Errorf("unknown type %s", reflect.TypeOf('a')),
		},
	}
	for name, tc := range testcases {
		t.Run(name, func(t *testing.T) {
			output, err := s.Contains(tc.search, tc.field, tc.data)
			if tc.err == nil {
				assert.Nil(t, err, "Was not expecting an error")
			} else {
				assert.NotNil(t, err, "Was expecting an error")
			}
			assert.Equal(t, tc.output, output, "Did not get expected output")
		})
	}
}
