package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Fire")
	want := "Hello, Fire"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
