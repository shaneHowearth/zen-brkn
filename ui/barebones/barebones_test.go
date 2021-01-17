package barebones_test

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"testing"

	"github.com/shanehowearth/zen/ui/barebones"
	"github.com/stretchr/testify/assert"
)

func TestShowHelp(t *testing.T) {
	s := barebones.Bb{}
	// Back up os.Stdout, so we can restore it later
	orig := os.Stdout
	r, w, _ := os.Pipe()
	// Set os.Stdout to our writer
	os.Stdout = w
	defer func() { os.Stdout = orig }()

	// Run command to be tested
	s.ShowHelp()
	// Close the writer, so we can move on to the next part of the test
	w.Close()
	var buf bytes.Buffer
	if _, err := io.Copy(&buf, r); err != nil {
		log.Fatalf("Copy error, cannot continue %v\n", err)
	}
	assert.Equal(t, buf.String(), `Welcome to Zendesk Search
Type 'quit' to exit at any time, Press 'Enter' to continue



		Select search options:
		* Press 1 to search Zendesk
		* Press 2 to view a list of searchable fields
		* Type 'quit' to exit

`, "Help string does not match")
}

func TestGetCommand(t *testing.T) {
	s := barebones.Bb{}
	testcases := map[string]struct {
		output map[string]string
		oErr   error
		input  []string
		rErr   error
	}{
		"Quit": {
			input:  []string{"quit\n"},
			output: map[string]string{"command": "quit"},
		},
		"Search error": {
			input: []string{"meaningless\n"},
			rErr:  fmt.Errorf("unable to get user search option with error EOF"),
		},
		// "Searchable fields": {
		// input: []string{"2\n"},
		// output: map[string]string(nil)
		// },
	}

	for name, tc := range testcases {
		t.Run(name, func(t *testing.T) {

			r, w, err := os.Pipe()
			if err != nil {
				t.Fatal(err)
			}
			orig := os.Stdin
			// Restore stdin right after the test.
			defer func() {
				os.Stdin = orig
			}()
			os.Stdin = r

			input := []byte(tc.input[0])
			_, err = w.Write(input)
			if err != nil {
				t.Error(err)
			}
			w.Close()

			if tc.rErr != nil {
				r = nil
			}
			// if len(tc.input) > 1 {
			// r, w, err = os.Pipe()
			// if err != nil {
			// t.Fatal(err)
			// }
			// input = []byte(tc.input[1])
			// _, err = w.Write(input)
			// if err != nil {
			// t.Error(err)
			// }
			// w.Close()
			// }

			output, err := s.GetCommand()

			if tc.rErr != nil {
				assert.Errorf(t, tc.rErr, err.Error(), "Error isn't what was expected")
			} else {
				assert.Nilf(t, err, "Got an unexpected error %v", err)
			}
			assert.Equal(t, tc.output, output, "Output did not match")
		})
	}
}

func TestShowResults(t *testing.T) {
	s := barebones.Bb{}
	testcases := map[string]struct {
		input  []string
		output string
	}{
		"Single item": {
			input:  []string{"Test String"},
			output: "Test String\n",
		},
	}
	for name, tc := range testcases {
		t.Run(name, func(t *testing.T) {
			orig := os.Stdout
			r, w, _ := os.Pipe()
			os.Stdout = w

			defer func() { os.Stdout = orig }()
			// Run command to be tested
			s.ShowResults(tc.input)
			w.Close()
			var buf bytes.Buffer
			if _, err := io.Copy(&buf, r); err != nil {
				log.Fatalf("Copy error, cannot continue %v\n", err)
			}
			assert.Equal(t, buf.String(), tc.output, "Output string does not match")
		})
	}
}
