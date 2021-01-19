package zensearch

import (
	"fmt"
	"os"
	"reflect"
	"testing"

	storage "github.com/shanehowearth/zen/repository"
	"github.com/shanehowearth/zen/search"
	userinterface "github.com/shanehowearth/zen/ui"
	"github.com/stretchr/testify/assert"
)

// Mocked instances

// Storage.store
type storeMock struct{}

var testGroup []map[string]interface{}
var testGroupErr error

func (st storeMock) GetGroup(string) ([]map[string]interface{}, error) {
	return testGroup, testGroupErr
}

var testGroupNames []string

func (st storeMock) GetGroupNames() []string {
	return testGroupNames
}

var testTerms map[string]reflect.Type
var testTermsErr error

func (st storeMock) GetTerms(string) (map[string]reflect.Type, error) {
	return testTerms, testTermsErr
}

// userinterface.UI
type uiMock struct{}

var testCommand map[string]string
var testCommandErr error

func (u uiMock) GetCommand() (map[string]string, error) {
	return testCommand, testCommandErr
}

var resultsErr error

func (u uiMock) ShowResults([]string) error {
	return resultsErr
}
func (u uiMock) Exit() {
	os.Exit(0)
}

// search.Search
type searchMock struct{}

var testContains map[string]interface{}
var testContainsErr error

func (se searchMock) Contains(substr, field string, data []map[string]interface{}) (map[string]interface{}, error) {
	return testContains, testContainsErr
}

// Tests start here
func TestNewZen(t *testing.T) {
	testcases := map[string]struct {
		store  storage.Store
		ui     userinterface.UI
		search search.Search
		err    error
	}{
		"New Zen": {
			store:  storeMock{},
			ui:     uiMock{},
			search: searchMock{},
		},
		"No datastore": {
			ui:     uiMock{},
			search: searchMock{},
			err:    fmt.Errorf("no storage.Store supplied"),
		},
		"No ui": {
			store:  storeMock{},
			search: searchMock{},
			err:    fmt.Errorf("no userinterface.UI supplied"),
		},
		"No search": {
			store: storeMock{},
			ui:    uiMock{},
			err:   fmt.Errorf("no search.Search supplied"),
		},
	}
	for name, tc := range testcases {
		t.Run(name, func(t *testing.T) {})
		z, err := NewZen(tc.store, tc.ui, tc.search)
		if tc.err != nil {
			assert.NotNil(t, err, "Was expecting an error")
			assert.Nil(t, z, "Expecting a nil instance of zen")
			assert.EqualError(t, tc.err, err.Error(), "Got different error")
		} else {
			assert.Nil(t, err, "Not expecting an error")
			assert.NotNil(t, z, "Expecting a non-nil instance of zen")
		}
	}
}

func TestSearchTerms(t *testing.T) {
	z, err := NewZen(storeMock{}, uiMock{}, searchMock{})
	assert.Nilf(t, err, "Zen creation returned error %v", err)
	assert.NotNil(t, z, "Expecting a non-nil instance of zen")
	testcases := map[string]struct {
		output     []string
		err        error
		groupNames []string
		testTerms  map[string]reflect.Type
	}{
		"Test terms error": {
			err:        fmt.Errorf("fake error"),
			output:     []string{"Cannot retrieve terms"},
			groupNames: []string{"A"},
		},
		"Successful execution": {
			groupNames: []string{"A"},
			output:     []string{"Search A with", "testName, \t\t\tstring", "---------------"},
			testTerms:  map[string]reflect.Type{"testName": reflect.TypeOf("a")},
		},
	}
	for name, tc := range testcases {
		t.Run(name, func(t *testing.T) {
			testGroupNames = tc.groupNames
			testTermsErr = tc.err
			testTerms = tc.testTerms
			output := z.searchTerms()
			assert.Equal(t, tc.output, output, "Output wasn't what was expected")
		})
	}
}

func TestFindMatch(t *testing.T) {
	z, err := NewZen(storeMock{}, uiMock{}, searchMock{})
	assert.Nilf(t, err, "Zen creation returned error %v", err)
	assert.NotNil(t, z, "Expecting a non-nil instance of zen")
	testcases := map[string]struct {
		getGroupErr error
		containsErr error
		cmds        map[string]string
		output      []string
		contains    map[string]interface{}
	}{
		"Get Group error": {
			cmds:        map[string]string{"group": "test"},
			getGroupErr: fmt.Errorf("fake group err"),
			output:      []string{fmt.Sprintf("Cannot retrieve %q", "test")},
		},
		"Contains error": {
			cmds:        map[string]string{"value": "test", "term": "test term", "group": "test group"},
			containsErr: fmt.Errorf("fake contains err"),
			output:      []string{fmt.Sprintf("Search %q of %q returned error: %v", "test term", "test group", fmt.Errorf("fake contains err"))},
		},
		"Nothing found": {
			cmds:   map[string]string{"value": "test", "term": "test term", "group": "test group"},
			output: []string{fmt.Sprintf("Search %q for %q with value of %q", "test group", "test term", "test"), "No results found"},
		},
		"One row": {
			cmds:     map[string]string{"value": "test", "term": "test term", "group": "test group"},
			contains: map[string]interface{}{"testKey": "testVal"},
			output:   []string{"testKey, \ttestVal"},
		},
	}
	for name, tc := range testcases {
		t.Run(name, func(t *testing.T) {
			testGroupErr = tc.getGroupErr
			testContains = tc.contains
			testContainsErr = tc.containsErr
			output := z.findMatch(tc.cmds)
			assert.Equal(t, tc.output, output, "findMatch output wasn't what was expected")
		})
	}
}
