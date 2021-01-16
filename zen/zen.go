// Package zensearch -
package zensearch

import (
	"fmt"
	"log"
	"reflect"

	storage "github.com/shanehowearth/zen/repository"
	"github.com/shanehowearth/zen/search"
	userinterface "github.com/shanehowearth/zen/ui"
)

type zen struct {
	datastore storage.Store
	ui        userinterface.UI
	searcher  search.Search
}

// NewZen -
// Quiet linter so pointer to non-exported type can be returned
// nolint: golint
func NewZen(d storage.Store, u userinterface.UI, s search.Search) (*zen, error) {
	if reflect.ValueOf(d).IsNil() {
		return nil, fmt.Errorf("no storage.Store supplied")
	}
	if reflect.ValueOf(u).IsNil() {
		return nil, fmt.Errorf("no userinterface.UI supplied")
	}
	if reflect.ValueOf(s).IsNil() {
		return nil, fmt.Errorf("no search.Search supplied")
	}
	return &zen{datastore: d, ui: u, searcher: s}, nil
}

// Run -
func (z *zen) Run() {
	if err := z.datastore.LoadData(); err != nil {
		log.Printf("unable to load data with error %v", err)
		z.ui.Exit()
	}

	// Main loop
	for {
		cmds, err := z.ui.GetCommand()
		if err != nil {
			log.Printf("ui.GetCommand error %v", err)
		}
		switch cmds["command"] {
		case "quit", "q":
			// Quit
			z.ui.Exit()
		case "2":
			// Get available search terms
			terms, err := z.datastore.GetTerms(cmds["group"])
			if err != nil {
				z.ui.ShowResults([]string{"Cannot retrieve terms"})
				break
			}
			termVals := []string{}
			for k, v := range terms {
				termVals = append(termVals, fmt.Sprintf("%s, \t%s", k, v))
			}
			z.ui.ShowResults(termVals)
		case "1":
			// Get lines from the specified group that contain the specified
			// search term
			group, err := z.datastore.GetGroup(cmds["group"])
			if err != nil {
				z.ui.ShowResults([]string{fmt.Sprintf("Cannot retrieve %s", cmds["group"])})
				break
			}
			// Get the terms
			// terms, err := z.datastore.GetTerms(cmds["group"])
			// if err != nil {
			// z.ui.ShowResults([]string{"Cannot retrieve terms"})
			// break
			// }

			d, err := z.searcher.Contains(cmds["value"], cmds["term"], group)
			if err != nil {
				z.ui.ShowResults([]string{fmt.Sprintf("Cannot complete search of %s for %s", cmds["term"], group)})
				break
			}
			dVals := []string{}
			for k, v := range d {
				dVals = append(dVals, fmt.Sprintf("%s, \t%s", k, v))

			}
			z.ui.ShowResults(dVals)
		}
	}
}
