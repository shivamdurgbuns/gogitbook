package main

import (
	"testing"
)

func TestHello(t *testing.T) {
	got := hello("Shivam")
	want := "Hello Shivam!"

	if got != want {
		t.Errorf("got: %q want: %q", got, want)
	}
}
