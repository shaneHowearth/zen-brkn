// Package simple -
package simple

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"text/tabwriter"
)

// Simple -
type Simple struct{}

// ShowHelp -
func (s Simple) ShowHelp() {
	fmt.Print(`Welcome to Zendesk Search
Type 'quit' to exit at any time, Press 'Enter' to continue



		Select search options:
		* Press 1 to search Zendesk
		* Press 2 to view a list of searchable fields
		* Type 'quit' to exit

`)
}

// Allow os.Exit to be faked in tests
var osExit = os.Exit

// Exit -
func (s Simple) Exit() {
	osExit(0)
}

func setGroup(group string) string {
	switch group {
	case "1":
		group = "Users"
	case "2":
		group = "Tickets"
	case "3":
		group = "Organizations"
	}
	return group
}

// GetCommand - Read command(s) from terminal
func (s Simple) GetCommand() (map[string]string, error) {
	s.ShowHelp()

	m := map[string]string{}

	// Search option
	reader := bufio.NewReader(os.Stdin)
	cmdString, err := reader.ReadString('\n')
	if err != nil {
		return nil, fmt.Errorf("unable to get user search option with error %w", err)
	}
	m["command"] = strings.TrimSpace(cmdString)
	// Quit option
	if strings.Contains(cmdString, "q") {
		return m, nil
	}
	// List search terms
	// TODO - ensure that only one rune exists here, what happens if the user
	// hits "12"
	searchTerms := false
	if strings.Contains(cmdString, "2") {
		searchTerms = true
	}
	// Search group
	fmt.Println("Select 1) Users or 2) Tickets or 3) Organizations")
	cmdString, err = reader.ReadString('\n')
	if err != nil {
		return nil, fmt.Errorf("unable to get search group with error %w", err)
	}
	if strings.Contains(cmdString, "q") {
		m["command"] = strings.TrimSpace(cmdString)
		return m, nil
	}
	m["group"] = setGroup(strings.TrimSpace(cmdString))
	if searchTerms {
		return m, nil
	}

	// Search Term
	fmt.Println("Enter search term")
	cmdString, err = reader.ReadString('\n')
	if err != nil {
		return nil, fmt.Errorf("unable to get search term with error %w", err)
	}
	m["term"] = strings.TrimSpace(cmdString)

	// Search Value
	fmt.Println("Enter search value")
	cmdString, err = reader.ReadString('\n')
	if err != nil {
		return nil, fmt.Errorf("unable to get search value with error %w", err)
	}
	m["value"] = strings.TrimSpace(cmdString)

	return m, nil
}

// ShowResults - Print out results
func (s Simple) ShowResults(output []string) error {
	// Align columns to the left
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', 0)
	for i := range output {
		fmt.Fprintln(w, output[i])
	}
	w.Flush()
	return nil
}
