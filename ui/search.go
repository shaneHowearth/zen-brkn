// Package userinterface -
package userinterface

// UI -
type UI interface {
	GetCommand() (map[string]string, error)
	ShowResults([]string) error
	Exit()
}
