// Package jsonfiles -
package jsonfiles

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"reflect"
	"strings"
)

// datastore -
type datastore struct {
	// filepaths map[string]string
	data   map[string][]map[string]interface{}
	terms  map[string]map[string]reflect.Type
	groups []string
}

// NewDatastore - New instance of a datastore, with checks to ensure required
// filepaths are specified.
// Allow return of unexported type:
// nolint: golint
func NewDatastore(filepaths map[string]string) (*datastore, error) {
	if len(filepaths) == 0 {
		return nil, fmt.Errorf("no files supplied, cannot continue")
	}
	data := map[string][]map[string]interface{}{}
	terms := map[string]map[string]reflect.Type{}
	groups := []string{}
	ds := datastore{
		// filepaths: filepaths,
		data:   data,
		terms:  terms,
		groups: groups,
	}
	var err error
	for k := range filepaths {
		ds.groups = append(ds.groups, k)
		ds.data[k], err = ds.loadFileData(filepaths[k])
		if err != nil {
			return nil, err
		}
		ds.terms[k] = ds.loadTerms(ds.data[k][0])
	}
	return &ds, nil

}

// GetGroupNames - Get all the names of groups that this datastore knows about.
func (ds datastore) GetGroupNames() []string {
	return ds.groups
}

// GetGroup - Get all of the items for this group.
func (ds datastore) GetGroup(group string) ([]map[string]interface{}, error) {
	g := strings.Title(strings.ToLower(group))
	if _, ok := ds.data[g]; ok {
		return ds.data[g], nil
	}
	return nil, fmt.Errorf("%s does not exist", group)
}

// GetTerms - get all of the terms that can be used to search this group.
func (ds datastore) GetTerms(group string) (map[string]reflect.Type, error) {
	g := strings.Title(strings.ToLower(group))
	if _, ok := ds.terms[g]; ok {
		return ds.terms[g], nil
	}
	return nil, fmt.Errorf("%s does not exist", group)
}

// Make os.Stat changeable for testing.
var osStat = os.Stat

// Check file exists, and we have permissons to access it.
func (ds datastore) fileExists(filepath string) error {
	if _, err := osStat(filepath); err == nil {
		return nil
	} else if os.IsNotExist(err) {
		return fmt.Errorf("%s does not exist", filepath)
	} else {
		// Some other error, eg. permissons
		return fmt.Errorf("%s caused error %w", filepath, err)
	}

}

// Make ioutil.ReadFile changeable for test purposes
var ioutilReadFile = ioutil.ReadFile

// Read file contents into memory.
func (ds datastore) loadFile(filepath string) ([]byte, error) {
	if err := ds.fileExists(filepath); err != nil {
		return nil, fmt.Errorf("unable to load file with error %w", err)
	}

	data, err := ioutilReadFile(filepath)
	if err != nil {
		return nil, fmt.Errorf("readfile error %w", err)
	}
	return data, nil
}

// Make json.Unmarshal changeable for testing.
var jsonUnmarshal = json.Unmarshal

func (ds datastore) loadFileData(file string) ([]map[string]interface{}, error) {
	// OS independent path
	p := filepath.FromSlash(file)
	data, err := ds.loadFile(p)
	if err != nil {
		return nil, fmt.Errorf("%s loaddata error %w", file, err)
	}

	var tempData interface{}

	err = jsonUnmarshal(data, &tempData)
	if err != nil {
		return nil, fmt.Errorf("%s unmarshal error %w", file, err)
	}

	m := []map[string]interface{}{}
	for i := range tempData.([]interface{}) {
		m = append(m, (tempData.([]interface{}))[i].(map[string]interface{}))
	}
	return m, nil
}

func (ds datastore) loadTerms(data map[string]interface{}) map[string]reflect.Type {
	keys := map[string]reflect.Type{}
	for k := range data {
		keys[k] = reflect.TypeOf(data[k])
	}
	return keys

}
