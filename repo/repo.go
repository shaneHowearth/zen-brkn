// Package repo -
package repo

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

// Data -
type Data struct {
	Organisations []*Organisation
	Users         []*User
	Tickets       []*Ticket
	OrgIdx        map[string]map[string][]*Organisation // map[fieldname]map[fieldvalue][]*Organisation
	UserIdx       map[string]map[string][]*User         // map[fieldname]map[fieldvalue][]*User
	TicketIdx     map[string]map[string][]*Ticket       // map[fieldname]map[fieldvalue][]*Ticket
}

// Make os.Stat changeable for testing.
var osStat = os.Stat

// Check file exists, and we have permissons to access it.
func (d Data) fileExists(filepath string) error {
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
var osReadFile = os.ReadFile

// Read file contents into memory.
func (d Data) loadFile(filepath string) ([]byte, error) {
	if err := d.fileExists(filepath); err != nil {
		return nil, fmt.Errorf("unable to load file with error %w", err)
	}

	data, err := osReadFile(filepath)
	if err != nil {
		return nil, fmt.Errorf("readfile error %w", err)
	}
	return data, nil
}

// Make json.Unmarshal changeable for testing.
var jsonUnmarshal = json.Unmarshal

// LoadJSON -
func (d *Data) LoadJSON() error {
	files := []struct {
		name      string
		container interface{}
	}{
		{"data/tickets.json", &d.Tickets},
		{"data/users.json", &d.Users},
		{"data/organizations.json", &d.Organisations},
	}

	for i := range files {
		p := filepath.FromSlash(files[i].name)

		// load the file into memory
		data, err := d.loadFile(p)
		if err != nil {
			return fmt.Errorf("%s loaddata error %w", files[i].name, err)
		}

		// extract the json from the file and load it into structs
		err = jsonUnmarshal(data, files[i].container)
		if err != nil {
			return fmt.Errorf("%s unmarshal error %w", files[i].name, err)
		}
	}

	// Create indexes
	d.TicketIndexes()
	d.OrganisationIndexes()
	d.UserIndexes()
	return nil
}
