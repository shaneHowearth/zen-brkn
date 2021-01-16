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

// Organization -
type Organization struct {
	ID            int      `json:"_id"`
	URL           string   `json:"url"`
	ExternalID    string   `json:"external_id"`
	Name          string   `json:"name"`
	DomainNames   []string `json:"domain_names"`
	CreatedAt     string   `json:"created_at"`
	Details       string   `json:"details"`
	SharedTickets bool     `json:"shared_tickets"`
	Tags          []string `json:"tags"`
}

// Ticket -
type Ticket struct {
	ID             string   `json:"_id"`
	URL            string   `json:"url"`
	ExternalID     string   `json:"external_id"`
	CreatedAt      string   `json:"created_at"`
	Type           string   `json:"type"`
	Subject        string   `json:"subject"`
	Description    string   `json:"description"`
	Priority       string   `json:"priority"`
	Status         string   `json:"status"`
	SubmitterID    int      `json:"submitter_id"`
	AssigneeID     int      `json:"assignee_id"`
	OrganizationID int      `json:"organization_id"`
	Tags           []string `json:"tags"`
	HasIncidents   bool     `json:"has_incidents"`
	DueAt          string   `json:"due_at"`
	Via            string   `json:"via"`
}

// User -
type User struct {
	ID             int      `json:"_id"`
	URL            string   `json:"url"`
	ExternalID     string   `json:"external_id"`
	Name           string   `json:"name"`
	Alias          string   `json:"alias"`
	CreatedAt      string   `json:"created_at"`
	Active         bool     `json:"active"`
	Verified       bool     `json:"verified"`
	Shared         bool     `json:"shared"`
	Locale         string   `json:"locale"`
	Timezone       string   `json:"timezone"`
	LastLoginAt    string   `json:"last_login_at"`
	Email          string   `json:"email"`
	Phone          string   `json:"phone"`
	Signature      string   `json:"signature"`
	OrganizationID int      `json:"organization_id"`
	Tags           []string `json:"tags"`
	Suspended      bool     `json:"suspended"`
	Role           string   `json:"role"`
}

// datastore -
type datastore struct {
	filepaths map[string]string
	data      map[string][]map[string]interface{}
	terms     map[string]map[string]reflect.Type
}

// NewDatastore - New instance of a datastore, with checks to ensure required
// filepaths are specified.
// Allow return of unexported type:
// nolint: golint
func NewDatastore(filepaths map[string]string) (*datastore, error) {
	if _, ok := filepaths["Organizations"]; !ok {
		return nil, fmt.Errorf("no 'Organizations' filepath specified")
	}
	if _, ok := filepaths["Users"]; !ok {
		return nil, fmt.Errorf("no 'Users' filepath specified")
	}
	if _, ok := filepaths["Tickets"]; !ok {
		return nil, fmt.Errorf("no 'Tickets' filepath specified")
	}

	return &datastore{filepaths: filepaths, data: map[string][]map[string]interface{}{}, terms: map[string]map[string]reflect.Type{}}, nil
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

func (ds datastore) fileExists(filepath string) error {
	// check file exists, and we have perms
	if _, err := os.Stat(filepath); err == nil {
		return nil
	} else if os.IsNotExist(err) {
		return fmt.Errorf("%s does not exist", filepath)
	} else {
		// Some other error, eg. permissons
		return fmt.Errorf("%s caused error %w", filepath, err)
	}

}

func (ds datastore) loadFile(filepath string) ([]byte, error) {
	if err := ds.fileExists(filepath); err != nil {
		return nil, fmt.Errorf("unable to load file with error %w", err)
	}

	data, err := ioutil.ReadFile(filepath)
	if err != nil {
		return nil, fmt.Errorf("readfile error %w", err)
	}
	return data, nil
}

func (ds datastore) loadFileData(file string) ([]map[string]interface{}, error) {
	// OS independent path
	p := filepath.FromSlash(file)
	data, err := ds.loadFile(p)
	if err != nil {
		return nil, fmt.Errorf("%s loaddata error %w", file, err)
	}

	var tempData interface{}

	err = json.Unmarshal(data, &tempData)
	if err != nil {
		return nil, fmt.Errorf("%s unmarshal error %w", file, err)
	}

	m := []map[string]interface{}{}
	for i := range tempData.([]interface{}) {
		m = append(m, (tempData.([]interface{}))[i].(map[string]interface{}))
	}
	return m, nil
}

func (ds datastore) loadTerms(data map[string]interface{}) (map[string]reflect.Type, error) {
	keys := map[string]reflect.Type{}
	for k := range data {
		keys[k] = reflect.TypeOf(data[k])
	}
	return keys, nil

}

// LoadData - Load all the data from the files into RAM
func (ds datastore) LoadData() (err error) {
	ds.data["Tickets"], err = ds.loadFileData(ds.filepaths["Tickets"])
	if err != nil {
		return err
	}
	ds.terms["Tickets"], err = ds.loadTerms(ds.data["Tickets"][0])
	if err != nil {
		return err
	}

	ds.data["Users"], err = ds.loadFileData(ds.filepaths["Users"])
	if err != nil {
		return err
	}
	ds.terms["Users"], err = ds.loadTerms(ds.data["Users"][0])
	if err != nil {
		return err
	}

	ds.data["Organizations"], err = ds.loadFileData(ds.filepaths["Organizations"])
	if err != nil {
		return err
	}
	ds.terms["Organizations"], err = ds.loadTerms(ds.data["Organizations"][0])

	return err
}
