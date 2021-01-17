package simple

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExit(t *testing.T) {
	s := Simple{}
	osExit = func(i int) { panic(fmt.Sprintf("Called with %d", i)) }
	defer func() { osExit = os.Exit }()
	assert.PanicsWithValue(t, "Called with 0", s.Exit, "os.Exit was not called")
}
