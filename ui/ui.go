// Package ui -
package ui

// CLI -
type CLI struct{}

// WelcomeMenu -
func (c *CLI) WelcomeMenu() string {
	return `Welcome to Zendesk Search
Type 'quit' to exit at any time, Press 'Enter' to continue



	Select search options:
	 * Press 1 to search Zendesk
	 * Press 2 to view a list of searchable fields
	 * Type 'quit' or 'q' to exit

`
}

// DataMenu -
func (c *CLI) DataMenu() string {
	return "Select 1) Users or 2) Tickets or 3) Organizations"
}
