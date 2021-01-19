// Package userinterface -
package userinterface

// UI -
type UI interface {
	GetCommand([]string) (map[string]string, error)
	ShowResults([]string) error
	Exit()
}
