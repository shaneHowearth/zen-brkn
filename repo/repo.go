// Package repo -
package repo

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"reflect"
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
	tickets := []*Ticket{}
	users := []*User{}
	organisations := []*Organisation{}
	files := []struct {
		name      string
		filename  string
		container interface{}
	}{
		{"ticket", "data/tickets.json", &tickets},
		{"user", "data/users.json", &users},
		{"organisation", "data/organizations.json", &organisations},
	}

	d.Data = map[string][]map[string][]string{}
	d.Indexes = map[string]map[string]map[string][]item{}
	d.Terms = map[string]map[string]struct{}{}
	for i := range files {
		p := filepath.FromSlash(files[i].filename)

		// load the file into memory
		data, err := d.loadFile(p)
		if err != nil {
			return fmt.Errorf("%s loaddata error %w", files[i].filename, err)
		}

		// extract the json from the file and load it into structs
		err = jsonUnmarshal(data, files[i].container)
		if err != nil {
			return fmt.Errorf("%s unmarshal error %w", files[i].filename, err)
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
