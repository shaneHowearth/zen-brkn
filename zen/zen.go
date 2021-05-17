// Package zen -
package zen

import (
	"fmt"
	"reflect"
	"strconv"
)

type brain struct {
	Data Data
	UI   UI
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

// NewBrain -
// Ignore linter complaint returning a non-nexported type
// nolint:golint
func NewBrain(data Data, ui UI) (*brain, error) {
	if isNilFixed(data) {
		return nil, fmt.Errorf("nil data object supplied")
	}
	if isNilFixed(ui) {
		return nil, fmt.Errorf("nil UI upplied")
	}
	return &brain{Data: data, UI: ui}, nil
}

// Data -
type Data interface {
	FindMatches(group, term, value string) ([]map[string]string, error)
	FindRelated(map[string]string) ([]map[string]string, error)
	GetGroups() ([]string, error)
	GetTerms(string) (map[string]struct{}, error)
}

// UI -
type UI interface {
	WelcomeMenu() string
	DataMenu() string
	GetCommand() (string, error)
	ShowGroupQuestion() error
	ShowTermQuestion() error
	ShowValueQuestion() error
	ShowResults([]map[string]string)
	ShowError(string)
}

func (b *brain) getGroup(maxTries int) (string, error) {
	groups := b.Data.GetGroups()
	groupNum := 0
	i := 0
	for {
		if i >= maxTries {
			b.UI.ShowError("Too many attempts, sorry.")
			return "", fmt.Errorf("too many attempts")
		}
		if err := b.UI.ShowGroupQuestion(); err != nil {
			return "", fmt.Errorf("search group question error %w", err)
		}
		group, err := b.UI.GetCommand()
		if err != nil {
			b.UI.ShowError("Error getting input, cannot continue.")
			return "", fmt.Errorf("input error %w", err)
		}
		if group == "q" || group == "quit" {
			return "", fmt.Errorf("quit signal")
		}
		groupNum, err = strconv.Atoi(group)
		if err != nil || groupNum <= 1 || groupNum > len(groups) {
			b.UI.ShowError(fmt.Sprintf("Invalid selection, please choose between 1 and %d, or 'quit' at any time.", len(groups)))
			i++
			continue
		}
		break
	}
	return groups[groupNum-1], nil
}

func (b *brain) getTerm(maxTries int, group string) (string, error) {
	terms, err := b.Data.GetTerms(group)
	if err != nil {
		// Group doesn't exist - how did /that/ happen?
	}
	i := 0
	termName := ""
	for {
		if i >= maxTries {
			b.UI.ShowError("Too many attempts, sorry.")
			return "", fmt.Errorf("too many attempts")
		}
		if err = b.UI.ShowTermQuestion(); err != nil {
			return "", fmt.Errorf("search term question error %w", err)
		}
		termName, err = b.UI.GetCommand()
		if err != nil {
			b.UI.ShowError("Error getting input, cannot continue.")
			return "", fmt.Errorf("input error %w", err)
		}
		if termName == "q" || termName == "quit" {
			return "", fmt.Errorf("quit signal")
		}
		if _, ok := terms[termName]; !ok {
			b.UI.ShowError("Invalid selection, please choose a valid search term, or 'quit' at any time.")
			i++
			continue
		}
		break
	}
	return termName, nil
}

func (b *brain) getValue(maxTries int, group, term string) (string, error) {
	if err := b.UI.ShowValueQuestion(); err != nil {
		return "", fmt.Errorf("search value question error %w", err)
	}
	termValue, err := b.UI.GetCommand()
	if err != nil {
		b.UI.ShowError("Error getting input, cannot continue.")
		return "", fmt.Errorf("input error %w", err)
	}
	return termValue, nil
}

// MaxTries - Maximum times a user can try to enter data for a single question
const MaxTries = 3

func (b *brain) Forever() {
	for {
		b.UI.WelcomeMenu()
		b.UI.DataMenu()
		group, err := b.getGroup(MaxTries)
		if err != nil {
		}
		term, err := b.getTerm(MaxTries, group)
		if err != nil {
		}
		value, err := b.getValue(MaxTries, group, term)
		if err != nil {
		}

		found := []map[string]string{}
		matches := b.Data.FindMatches(group, term, value)
		found = append(found, matches...)
		for idx := range matches {
			found = append(found, b.Data.FindRelated(matches[idx])...)
		}
		b.UI.ShowResults(found)

	}
}
