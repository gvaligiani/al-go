package dict

import (
	"testing"
)

// Require tests a map against an expected one ( same number of items, same values )
func Require[K comparable, V comparable](t *testing.T, expected, got map[K]V) {
	if len(expected) != len(got) {
		t.Errorf("invalid length! ( expected: %d, got: %d )", len(expected), len(got))
	}
	for key, expectedValue := range expected {
		if gotValue, found := got[key]; !found {
			t.Errorf("missing item %v", key)
		} else if expectedValue != gotValue {
			t.Errorf("invalid item %v! ( expected: %v, got: %v )", key, expectedValue, gotValue)
		}
	}
	for item := range got {
		if _, found := expected[item]; !found {
			t.Errorf("unexpected item: %v", item)
		}
	}
}
