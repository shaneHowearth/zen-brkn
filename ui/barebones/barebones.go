// Package barebones -
package barebones

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"text/tabwriter"
)

// Bb -
type Bb struct{}

// ShowHelp -
func (b Bb) ShowHelp() {
	fmt.Print(`

Welcome to Zendesk Search
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
func (b Bb) Exit() {
	osExit(0)
}

// GetCommand - Read command(s) from terminal
func (b Bb) GetCommand(groups []string) (map[string]string, error) {
	b.ShowHelp()
	fmt.Print("> ")

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
	if strings.Contains(cmdString, "2") {
		return m, nil
	}
	// Search group
	getGroup := func() (string, error) {

		s := []string{}
		for i := range groups {
			s = append(s, fmt.Sprintf("%d) %s", i+1, groups[i]))
		}
		joined := ""
		if len(s) > 1 {
			joined = strings.Join(s[:len(s)-1], " or ")
			fmt.Printf("Select %s %s\n> ", joined, s[len(s)-1])
		} else {
			fmt.Printf("Select %s\n", s)
		}
		cmdString, err = reader.ReadString('\n')
		return cmdString, err
	}
	maxTries := 3
	var num int
	for i := 0; i < maxTries; i++ {
		// give the user 3 shots to enter decent input
		input := ""
		input, err = getGroup()
		if err != nil {
			return nil, fmt.Errorf("unable to get search group with error %w", err)
		}
		if strings.Contains(input, "q") {
			m["command"] = strings.TrimSpace(input)
			return m, nil
		}
		num, err = strconv.Atoi(strings.TrimSpace(cmdString))
		if i+1 < maxTries {
			if err != nil {
				fmt.Printf("Please enter a valid group number, you have %d tries left\n", maxTries-i-1)
				continue
			}
			if num < 1 || num > len(groups) {
				fmt.Printf("Please enter a valid group number, you have %d tries left\n", maxTries-i-1)
				continue
			}
		} else {
			m["command"] = strings.TrimSpace("q")
			return m, nil

		}
		break
	}
	m["group"] = groups[num-1]

	// Search Term
	fmt.Println("Enter search term")
	fmt.Print("> ")
	cmdString, err = reader.ReadString('\n')
	if err != nil {
		return nil, fmt.Errorf("unable to get search term with error %w", err)
	}
	m["term"] = strings.TrimSpace(cmdString)

	// Search Value
	fmt.Println("Enter search value")
	fmt.Print("> ")
	cmdString, err = reader.ReadString('\n')
	if err != nil {
		return nil, fmt.Errorf("unable to get search value with error %w", err)
	}
	m["value"] = strings.TrimSpace(cmdString)

	return m, nil
}

// make os.Stdout changeable for tests.
var osStdout = os.Stdout

// ShowResults - Print out results
func (b Bb) ShowResults(input []string) error {
	fmt.Println()
	// Align columns to the left
	w := tabwriter.NewWriter(osStdout, 0, 0, 1, ' ', 0)
	for i := range input {
		fmt.Fprintln(w, input[i])
	}
	w.Flush()
	return nil
}
