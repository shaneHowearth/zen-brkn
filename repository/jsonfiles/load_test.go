package jsonfiles

import (
	"fmt"
	"os"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewDatastore(t *testing.T) {
	testcases := map[string]struct {
		filepaths map[string]string
		// ds        *datastore
		terms  map[string]map[string]reflect.Type
		groups []string
		err    error
		osErr  func(string) (os.FileInfo, error)
		ioErr  func(string) ([]byte, error)
		jsErr  func([]byte, interface{}) error
	}{
		"Single field, single instance, in single file": {
			filepaths: map[string]string{"Test": "data/testdata/single.json"},
			terms:     map[string]map[string]reflect.Type{"Test": {"_id": reflect.TypeOf(101.1)}},
			groups:    []string{"Test"},
		},
		"Bad filepath": {
			filepaths: map[string]string{"Test": "data/testdata/nonexistant.json"},
			err:       fmt.Errorf("data/testdata/nonexistant.json loaddata error unable to load file with error data/testdata/nonexistant.json does not exist"),
		},
		"Bad file": {
			// Note: This returns an error that *isn't* a bad filepath
			filepaths: map[string]string{"Test": "data/testdata/nonexistant.json"},
			err:       fmt.Errorf("data/testdata/nonexistant.json loaddata error unable to load file with error data/testdata/nonexistant.json caused error fake error"),
			osErr:     func(string) (os.FileInfo, error) { return nil, fmt.Errorf("fake error") },
		},
		"Cannot read file": {
			filepaths: map[string]string{"Test": "data/testdata/single.json"},
			err:       fmt.Errorf("data/testdata/single.json loaddata error readfile error fake read error"),
			ioErr:     func(string) ([]byte, error) { return nil, fmt.Errorf("fake read error") },
		},
		"json unmarshal error": {
			filepaths: map[string]string{"Test": "data/testdata/single.json"},
			err:       fmt.Errorf("data/testdata/single.json unmarshal error fake unmarshal error"),
			jsErr:     func([]byte, interface{}) error { return fmt.Errorf("fake unmarshal error") },
		},
		"No Filepaths": {
			filepaths: map[string]string{},
			err:       fmt.Errorf("no files supplied, cannot continue"),
		},
	}
	for name, tc := range testcases {
		t.Run(name, func(t *testing.T) {
			if tc.jsErr != nil {
				orig := jsonUnmarshal
				jsonUnmarshal = tc.jsErr
				defer func() { jsonUnmarshal = orig }()
			}
			if tc.ioErr != nil {
				orig := ioutilReadFile
				ioutilReadFile = tc.ioErr
				defer func() { ioutilReadFile = orig }()
			}
			if tc.osErr != nil {
				orig := osStat
				osStat = tc.osErr
				defer func() { osStat = orig }()
			}
			tds, err := NewDatastore(tc.filepaths)
			_, _ = tds, err
			if tc.err != nil {
				assert.Nil(t, tds, "Expected a nil datastore")
				assert.NotNil(t, err, "Expected an error but did not get one")
				assert.EqualError(t, tc.err, err.Error(), "Errors do not match")
			} else {
				assert.Nil(t, err, "Got an unexpected error")
				assert.Equal(t, tc.terms, tds.terms, "Terms did not match")
				assert.Equal(t, tc.groups, tds.groups, "Terms did not match")
			}
		})
	}
}

func TestGetGroupNames(t *testing.T) {
	filepaths := map[string]string{"Test": "data/testdata/single.json"}
	tds, err := NewDatastore(filepaths)
	assert.Nil(t, err, "No error was expected")
	assert.NotNil(t, tds, "A datastore was expected")
	output := tds.GetGroupNames()
	expectedOutput := []string{"Test"}
	assert.Equal(t, output, expectedOutput, "Got different group names")
}
func TestGetGroup(t *testing.T) {
	testcases := map[string]struct {
		filepaths      map[string]string
		groupName      string
		expectedOutput []map[string]interface{}
		expectedErr    error
	}{
		"Single Group": {
			filepaths:      map[string]string{"Test": "data/testdata/single.json"},
			groupName:      "Test",
			expectedOutput: []map[string]interface{}{{"_id": float64(101)}},
		},
		"Non-existant group": {
			filepaths:   map[string]string{"Test": "data/testdata/single.json"},
			groupName:   "nothing",
			expectedErr: fmt.Errorf("nothing does not exist"),
		},
	}
	for name, tc := range testcases {
		t.Run(name, func(t *testing.T) {

			tds, err := NewDatastore(tc.filepaths)
			assert.Nil(t, err, "No error was expected")
			assert.NotNil(t, tds, "A datastore was expected")
			output, err := tds.GetGroup(tc.groupName)
			assert.Equal(t, tc.expectedOutput, output, "Outputs did not match")
			if tc.expectedErr != nil {
				assert.NotNil(t, err, "Expected an error")
				assert.EqualError(t, tc.expectedErr, err.Error(), "Errors do not match")
			} else {
				assert.Nil(t, err, "Did not expect an error")
			}
		})
	}
}
func TestGetTerms(t *testing.T) {
	testcases := map[string]struct {
		filepaths      map[string]string
		groupName      string
		expectedOutput map[string]reflect.Type
		expectedErr    error
	}{
		"Single Group": {
			filepaths:      map[string]string{"Test": "data/testdata/single.json"},
			groupName:      "Test",
			expectedOutput: map[string]reflect.Type{"_id": reflect.TypeOf(float64(101))},
		},
		"Non-existant group": {
			filepaths:   map[string]string{"Test": "data/testdata/single.json"},
			groupName:   "nothing",
			expectedErr: fmt.Errorf("nothing does not exist"),
		},
	}
	for name, tc := range testcases {
		t.Run(name, func(t *testing.T) {

			tds, err := NewDatastore(tc.filepaths)
			assert.Nil(t, err, "No error was expected")
			assert.NotNil(t, tds, "A datastore was expected")
			output, err := tds.GetTerms(tc.groupName)
			assert.Equal(t, tc.expectedOutput, output, "Outputs did not match")
			if tc.expectedErr != nil {
				assert.NotNil(t, err, "Expected an error")
				assert.EqualError(t, tc.expectedErr, err.Error(), "Errors do not match")
			} else {
				assert.Nil(t, err, "Did not expect an error")
			}
		})
	}
}
