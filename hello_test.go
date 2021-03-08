package main

import (
	"testing"
)

/*
TestHello function will test various test cases for the hello program.
*/
func TestHello(t *testing.T) {

	assertCorrectMessage := func(t testing.TB, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got: %q want: %q", got, want)
		}
	}

	t.Run("Saying Hello to people.", func(t *testing.T) {
		got := Hello("Shivam", "English")
		want := "Hello, Shivam"
		assertCorrectMessage(t, got, want)

	})

	t.Run("Saying Hello, World when empty string is supplied.", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)

	})

	t.Run("Saying Hello in Spanish.", func(t *testing.T) {
		got := Hello("Shivam", "Spanish")
		want := "Hola, Shivam"
		assertCorrectMessage(t, got, want)

	})

	t.Run("Saying Hello in French.", func(t *testing.T) {
		got := Hello("Shivam", "French")
		want := "Bonjour, Shivam"
		assertCorrectMessage(t, got, want)
	})

	t.Run("Saying Hello name when empty language is supplied.", func(t *testing.T) {
		got := Hello("Shivam", "")
		want := "Hello, Shivam"
		assertCorrectMessage(t, got, want)

	})
}
