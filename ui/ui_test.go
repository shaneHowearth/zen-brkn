package ui_test

import (
	"testing"

	"github.com/shanehowearth/zen/ui"
	"github.com/stretchr/testify/assert"
)

func TestWelcomeMenu(t *testing.T) {
	expected := `Welcome to Zendesk Search
Type 'quit' to exit at any time, Press 'Enter' to continue



	Select search options:
	 * Press 1 to search Zendesk
	 * Press 2 to view a list of searchable fields
	 * Type 'quit' or 'q' to exit

`
	c := ui.CLI{}
	output := c.WelcomeMenu()
	assert.Equal(t, expected, output, "Welcome Menu output did not match")
}

func TestDataMenu(t *testing.T) {
	expected := "Select 1) Users or 2) Tickets or 3) Organizations"
	c := ui.CLI{}
	output := c.DataMenu()
	assert.Equal(t, expected, output, "Data Menu output did not match")
}
