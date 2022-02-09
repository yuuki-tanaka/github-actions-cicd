package main

import "testing"

func TestGreeting(t *testing.B) {
	want := "Hello"
	got := Greeting()

	if want != got {
		t.Errorf("got:%q , want:%q", got, want)
	}
}
