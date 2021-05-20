// Package ui -
package ui

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"sync"
	"text/tabwriter"

	"github.com/shanehowearth/zen/zen"
)

var _ zen.UI = (*CLI)(nil)

// CLI -
type CLI struct {
	m sync.RWMutex
	w *tabwriter.Writer
}

// make os.Stdout changeable for tests.
var osStdout = os.Stdout

// NewWriter -
func (c *CLI) NewWriter() {
	c.m.Lock()
	defer c.m.Unlock()
	c.w = tabwriter.NewWriter(osStdout, 0, 0, 1, ' ', 0)

}

func (c *CLI) toScreen(s []string) {
	if c.w == nil {
		c.NewWriter()
	}
	for i := range s {
		fmt.Fprint(c.w, s[i])
	}
	c.w.Flush()

}

// WelcomeMenu -
func (c *CLI) WelcomeMenu() (string, error) {
	c.toScreen([]string{`Welcome to Zendesk Search
Type 'quit' to exit at any time, Press 'Enter' to continue



	Select search options:
	 * Press 1 to search Zendesk
	 * Press 2 to view a list of searchable fields
	 * Type 'quit' or 'q' to exit

`})
	return c.GetCommand()
}

// DataMenu -
func (c *CLI) DataMenu(groups []string) {
	// menu := "Select 1) Users or 2) Tickets or 3) Organizations"
	menu := fmt.Sprintf("Select 1) %s", groups[0])
	for i := 2; i <= len(groups); i++ {
		menu += fmt.Sprintf(" or %d) %s", i, strings.Title(strings.ToLower(groups[i-1])))
	}
	c.toScreen([]string{fmt.Sprintf("%s\n", menu)})
}

// TermQuestion -
func (c *CLI) TermQuestion() (string, error) {
	c.toScreen([]string{"Enter search term"})
	return c.GetCommand()
}

// GroupQuestion -
func (c *CLI) GroupQuestion() (string, error) {
	c.toScreen([]string{"Enter search group"})
	return c.GetCommand()
}

// ValueQuestion -
func (c *CLI) ValueQuestion() (string, error) {
	c.toScreen([]string{"Enter search value"})
	return c.GetCommand()
}

// ShowTerms -
func (c *CLI) ShowTerms(t map[string][]string) {
	for k := range t {
		c.toScreen([]string{"----------------------------------------------------\n"})
		c.toScreen([]string{k})
		ts := make([]string, len(t[k]))
		for i := range t[k] {
			ts[i] = t[k][i] + "\n"
		}
		c.toScreen(ts)
	}
}

// ShowResults -
func (c *CLI) ShowResults(results []map[string][]map[string][]string) {
	for i := range results {
		c.toScreen([]string{"----------------------------------------------------\n"})
		d := []string{}
		for k := range results[i] {
			d = append(d, fmt.Sprintf("%s\n", k))
			for j := range results[i][k] {
				for jk, jv := range results[i][k][j] {
					d = append(d, fmt.Sprintf("%s\t\t\t%s\n", jk, strings.NewReplacer("[", "", "]", "").Replace(fmt.Sprintf("%v", jv))))
				}
				d = append(d, "--\n")
			}
		}
		c.toScreen(d)
	}
}

// ShowError -
func (c *CLI) ShowError(s string) {
	c.toScreen([]string{fmt.Sprintf("ERROR %s", s)})
}

// GetCommand -
func (c *CLI) GetCommand() (string, error) {
	// Prompt
	c.toScreen([]string{"> "})

	reader := bufio.NewReader(os.Stdin)
	cmdString, err := reader.ReadString('\n')
	if err != nil {
		return "", fmt.Errorf("unable to get user input with error %w", err)
	}
	return strings.TrimSpace(cmdString), nil
}
