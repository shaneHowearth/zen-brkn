package barebones

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExit(t *testing.T) {
	b := Bb{}
	osExit = func(i int) { panic(fmt.Sprintf("Called with %d", i)) }
	defer func() { osExit = os.Exit }()
	assert.PanicsWithValue(t, "Called with 0", b.Exit, "os.Exit was not called")
}

func TestShowResults(t *testing.T) {
	s := Bb{}
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
			orig := osStdout
			r, w, err := os.Pipe()
			if err != nil {
				fmt.Println("DIAF", err)
				os.Exit(0)
			}
			osStdout = w

			defer func() { osStdout = orig }()
			// Run command to be tested
			s.ShowResults(tc.input)

			w.Close()
			// var buf bytes.Buffer
			var buf strings.Builder
			if _, err := io.Copy(&buf, r); err != nil {
				log.Fatalf("Copy error, cannot continue %v\n", err)
			}
			assert.Equal(t, tc.output, buf.String(), "Output string does not match")
		})
	}
}
