package search

type UI interface {
	GetCommand() (map[string]string, error)
	ShowResults([]string) error
}
