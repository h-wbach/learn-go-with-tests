package main

import "testing"

func TestName(t *testing.T) {
	got := Name()
	want := "Hunter"

	if got != want {
		t.Errorf("got: %q, want: %q", got, want)
	}
}
