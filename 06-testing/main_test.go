package main

import "testing"

func TestSay(t *testing.T) {
	want := "Hello, Bob!"
	got := say([]string{"Bob"})

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
