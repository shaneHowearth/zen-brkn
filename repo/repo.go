// Package repo -
package repo

import (
	_ "embed" // embed
	"encoding/json"
	"fmt"
	"os"
	"reflect"

	"github.com/shanehowearth/zen/zen"
)

// item - all Data types must implement this interface
type item interface {
	ToDTO() map[string][]string
	CreateIndex(interface{}, string) map[string]map[string][]item
}

// Data -
type Data struct {
	// Book keeping
	Groups []string
	Terms  map[string]map[string]struct{}

	Data    map[string][]map[string][]string        // map[data group][]map[fieldname]field value
	Indexes map[string]map[string]map[string][]item // map[data group name]map[field name]map[field value][]item
}

var _ zen.Data = (*Data)(nil)

// Make os.Stat changeable for testing.
var osStat = os.Stat

// Check file exists, and we have permissons to access it.
func (d *Data) fileExists(filepath string) error {
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
func (d *Data) loadFile(filepath string) ([]byte, error) {
	if err := d.fileExists(filepath); err != nil {
		return nil, fmt.Errorf("unable to load file with error %w", err)
	}

	data, err := osReadFile(filepath)
	if err != nil {
		return nil, fmt.Errorf("readfile error %w", err)
	}
	return data, nil
}

// Initialise -
func (d *Data) Initialise() {
	err := d.LoadJSON()
	if err != nil {
		fmt.Printf("Loading data caused an error %v\n", err)
	}
}

// Make json.Unmarshal changeable for testing.
var jsonUnmarshal = json.Unmarshal

//go:embed data/tickets.json
var ticketData []byte

//go:embed data/users.json
var userData []byte

//go:embed data/organizations.json
var organisationData []byte

// LoadJSON -
func (d *Data) LoadJSON() error {
	tickets := []*Ticket{}
	users := []*User{}
	organisations := []*Organisation{}
	files := []struct {
		name      string
		fileData  []byte
		container interface{}
	}{
		{"ticket", ticketData, &tickets},
		{"user", userData, &users},
		{"organisation", organisationData, &organisations},
	}

	d.Data = map[string][]map[string][]string{}
	d.Indexes = map[string]map[string]map[string][]item{}
	d.Terms = map[string]map[string]struct{}{}
	for i := range files {
		d.Groups = append(d.Groups, files[i].name)
		// extract the json from the file and load it into structs
		err := jsonUnmarshal(files[i].fileData, files[i].container)
		if err != nil {
			return fmt.Errorf("%s unmarshal error %w", files[i].name, err)
		}

		// Convert the container into a DTO
		list := reflect.ValueOf(files[i].container)
		for j := 0; j < reflect.Indirect(list).Len(); j++ {
			if _, ok := d.Data[files[i].name]; !ok {
				d.Data[files[i].name] = []map[string][]string{}
			}
			d.Data[files[i].name] = append(d.Data[files[i].name], reflect.Indirect(list).Index(j).Interface().(item).ToDTO())
		}

		// Create Index for this group
		if reflect.Indirect(list).Len() >= 1 {
			d.Indexes[files[i].name] = reflect.Indirect(list).Index(0).Interface().(item).CreateIndex(reflect.Indirect(list).Interface(), files[i].name)
			d.Terms[files[i].name] = map[string]struct{}{}
			for k := range d.Indexes[files[i].name] {
				d.Terms[files[i].name][k] = struct{}{}
			}
		}
	}

	return nil
}
