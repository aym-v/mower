package test

import "testing"

// AssertEqual asserts that two objects are equal.
func AssertEqual(t *testing.T, got interface{}, want interface{}) {
	if got != want {
		t.Fatalf("got %s; want %s", got, want)
	}
}
