package list

import (
	"testing"
)

// Require tests a list against an expected one ( same number of items, same values with same order )
func Require[T comparable](t *testing.T, expected, got []T) {
	if len(expected) != len(got) {
		t.Errorf("invalid length! ( expected: %d, got: %d )", len(expected), len(got))
	}
	for i, item := range expected {
		if item != got[i] {
			t.Errorf("invalid item #%d! ( expected: %v, got: %v )", i, item, got[i])
		}
	}
}

// RequireUnordered tests a list against an expected one ( same number of items, same values with different order )
//
//	>>> helpful when testing list coming from a map ( order not guaranteed )
func RequireUnordered[T comparable](t *testing.T, expected, got []T) {
	if len(expected) != len(got) {
		t.Errorf("invalid length! ( expected: %d, got: %d )", len(expected), len(got))
	}
	gotIndexes := make(map[int]struct{}, len(expected))
	for _, expectedItem := range expected {
		found := false
		for j, gotItem := range got {
			if _, hasIndex := gotIndexes[j]; hasIndex || expectedItem != gotItem {
				continue
			}
			gotIndexes[j] = struct{}{}
			found = true
			break
		}
		if !found {
			t.Errorf("missing item %v!", expectedItem)
		}
	}
}
