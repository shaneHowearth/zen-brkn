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

func (z *zen) searchTerms() []string {
	vals := []string{}
	// Get available search terms
	for _, group := range z.datastore.GetGroupNames() {
		terms, err := z.datastore.GetTerms(group)
		if err != nil {
			return []string{"Cannot retrieve terms"}
		}
		vals = append(vals, fmt.Sprintf("Search %s with", group))
		tmpVals := []string{}
		for k, v := range terms {
			tmpVals = append(tmpVals, fmt.Sprintf("%s, \t\t\t%s", k, v))
		}
		sort.Strings(tmpVals)
		vals = append(vals, tmpVals...)
		vals = append(vals, "---------------")

	}
	return vals
}

func (z *zen) findMatch(cmds map[string]string) []string {
	// Get lines from the specified group that contain the specified
	// search term
	group, err := z.datastore.GetGroup(cmds["group"])
	if err != nil {
		return []string{fmt.Sprintf("Cannot retrieve %q", cmds["group"])}
	}

	d, err := z.searcher.Contains(cmds["value"], cmds["term"], group)
	if err != nil {
		return []string{fmt.Sprintf("Search %q of %q returned error: %v", cmds["term"], cmds["group"], err)}
	}

	vals := []string{}
	for k, v := range d {
		vals = append(vals, fmt.Sprintf("%s, \t%v", k, v))

	}
	sort.Strings(vals)
	if len(vals) == 0 {
		vals = append(vals, fmt.Sprintf("Search %q for %q with value of %q", cmds["group"], cmds["term"], cmds["value"]))
		vals = append(vals, "No results found")

	}
	return vals
}

// Run -
func (z *zen) Run() {
	// Main loop
	for {
		vals := []string{}
		cmds, err := z.ui.GetCommand(z.datastore.GetGroupNames())
		if err != nil {
			// Going to DIAF here because the error may be that the user cannot
			// provide input that the application can handle, which might stop a
			// 'quit' being sent.
			log.Printf("cannot continue ui.GetCommand error %v", err)
			z.ui.Exit()
		}
		switch cmds["command"] {
		case "quit", "q":
			// Quit
			z.ui.Exit()
		case "2":
			vals = z.searchTerms()
		case "1":
			vals = z.findMatch(cmds)
		}
		if err := z.ui.ShowResults(vals); err != nil {
			log.Printf("Show Results returned an error %v", err)
		}
	}
}
