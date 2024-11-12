package list_test

import (
	"testing"

	"github.com/gvaligiani/al-go/list"
)

func TestRequire(t *testing.T) {
	list.Require(t, list.New(1, 3, 4, 3, 2), list.New(1, 3, 4, 3, 2))
	list.RequireUnordered(t, list.New(1, 3, 4, 3, 2), list.New(1, 2, 3, 3, 4))
}
