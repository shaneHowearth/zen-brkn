// Package jsonfiles -
package jsonfiles

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
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

func fileExists(filepath string) error {
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

func loadData(filepath string) ([]byte, error) {
	if err := fileExists(filepath); err != nil {
		return nil, fmt.Errorf("unable to load file with error %w", err)
	}

	data, err := ioutil.ReadFile(filepath)
	if err != nil {
		return nil, fmt.Errorf("readfile error %w", err)
	}
	return data, nil
}

// LoadTicketData -
func LoadTicketData(filename, path string) error {
	// OS independent path
	p := filepath.FromSlash(filepath.Join(path, filename))
	data, err := loadData(p)
	if err != nil {
		return fmt.Errorf("loadTicketData loaddata error %w", err)
	}

	var ticket []Ticket

	err = json.Unmarshal(data, &ticket)
	if err != nil {
		return fmt.Errorf("loadTicketData unmarshal error %w", err)
	}

	return nil
}

// LoadOrgData -
func LoadOrgData(filename, path string) error {
	// OS independent path
	p := filepath.FromSlash(filepath.Join(path, filename))
	data, err := loadData(p)
	if err != nil {
		return fmt.Errorf("loadOrgData loaddata error %w", err)
	}

	var org []Organization

	err = json.Unmarshal(data, &org)
	if err != nil {
		return fmt.Errorf("loadOrgData unmarshal error %w", err)
	}

	return nil
}

// LoadUserData -
func LoadUserData(filename, path string) error {
	// OS independent path
	p := filepath.FromSlash(filepath.Join(path, filename))
	data, err := loadData(p)
	if err != nil {
		return fmt.Errorf("loadUserData loaddata error %w", err)
	}

	var user []User

	err = json.Unmarshal(data, &user)
	if err != nil {
		return fmt.Errorf("loadUserData unmarshal error %w", err)
	}

	return nil
}
