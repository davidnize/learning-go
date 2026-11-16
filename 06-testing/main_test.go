package main

import "testing"

func TestSay(t *testing.T) {
	want := "Hello, Bob!"
	got := say([]string{"Bob"})

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
func TestSayDefault(t *testing.T) {
	want := "Hello, World!"
	got := say([]string{})

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
