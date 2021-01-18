// Package zensearch -
package zensearch

import (
	"fmt"
	"log"
	"reflect"
	"sort"

	storage "github.com/shanehowearth/zen/repository"
	"github.com/shanehowearth/zen/search"
	userinterface "github.com/shanehowearth/zen/ui"
)

type zen struct {
	datastore storage.Store
	ui        userinterface.UI
	searcher  search.Search
}

func isNilFixed(i interface{}) bool {
	if i == nil {
		return true

	}
	switch reflect.TypeOf(i).Kind() {
	case reflect.Ptr, reflect.Map, reflect.Array, reflect.Chan, reflect.Slice:
		return reflect.ValueOf(i).IsNil()

	}
	return false

}

// NewZen -
// Quiet linter so pointer to non-exported type can be returned
// nolint: golint
func NewZen(d storage.Store, u userinterface.UI, s search.Search) (*zen, error) {
	if isNilFixed(d) {
		return nil, fmt.Errorf("no storage.Store supplied")
	}
	if isNilFixed(u) {
		return nil, fmt.Errorf("no userinterface.UI supplied")
	}
	if isNilFixed(s) {
		return nil, fmt.Errorf("no search.Search supplied")
	}
	return &zen{datastore: d, ui: u, searcher: s}, nil
}

// Make log.Printf changeable for testing
var logPrintf = log.Printf

// Run -
func (z *zen) Run() {
	// Main loop
	for {
		cmds, err := z.ui.GetCommand()
		if err != nil {
			logPrintf("ui.GetCommand error %v", err)
			break
		}
		switch cmds["command"] {
		case "quit", "q":
			// Quit
			z.ui.Exit()
		case "2":
			// Get available search terms
			for _, group := range z.datastore.GetGroupNames() {

				terms, err := z.datastore.GetTerms(group)
				if err != nil {
					z.ui.ShowResults([]string{"Cannot retrieve terms"})
					break
				}
				termVals := []string{fmt.Sprintf("Search %s with", group)}
				tmpVals := []string{}
				for k, v := range terms {
					tmpVals = append(tmpVals, fmt.Sprintf("%s, \t\t\t%s", k, v))
				}
				sort.Strings(tmpVals)
				termVals = append(termVals, tmpVals...)
				termVals = append(termVals, fmt.Sprint("---------------"))
				z.ui.ShowResults(termVals)
			}
		case "1":
			// Get lines from the specified group that contain the specified
			// search term
			group, err := z.datastore.GetGroup(cmds["group"])
			if err != nil {
				z.ui.ShowResults([]string{fmt.Sprintf("Cannot retrieve %s", cmds["group"])})
				break
			}

			d, err := z.searcher.Contains(cmds["value"], cmds["term"], group)
			if err != nil {
				z.ui.ShowResults([]string{fmt.Sprintf("Search %q of %q returned error: %v", cmds["term"], cmds["group"], err)})
				break
			}
			dVals := []string{}
			for k, v := range d {
				dVals = append(dVals, fmt.Sprintf("%s, \t%v", k, v))
			}
			sort.Strings(dVals)
			if len(dVals) == 0 {
				dVals = append(dVals, fmt.Sprintf("Search %q for %q with value of %v", cmds["group"], cmds["term"], cmds["value"]))
				dVals = append(dVals, "No results found")
			}
			z.ui.ShowResults(dVals)
		}
	}
}
