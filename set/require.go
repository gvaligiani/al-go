package set

import (
	"testing"
)

// Require tests a set against an expected one ( same number of items, same values )
func Require[T comparable](t *testing.T, expected, got map[T]struct{}) {
	if len(expected) != len(got) {
		t.Errorf("invalid length! ( expected: %d, got: %d )", len(expected), len(got))
	}
	for item := range expected {
		if _, found := got[item]; !found {
			t.Errorf("missing item %v", item)
		}
	}
	for item := range got {
		if _, found := expected[item]; !found {
			t.Errorf("unexpected item %v", item)
		}
	}
}
