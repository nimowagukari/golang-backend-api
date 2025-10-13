package main

import "testing"

func TestGreeting(t *testing.T) {
	got := greeting("Hoge")
	want := "Hello, Hoge"
	if got != want {
		t.Errorf("got: %s, want: %s", got, want)
	}
}
