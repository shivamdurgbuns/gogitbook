package main

import (
	"testing"
)

func TestHello(t *testing.T) {

	assertCorrectMessage := func(t testing.TB, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got: %q want: %q", got, want)
		}
	}

	t.Run("Saying Hello to people.", func(t *testing.T) {
		got := hello("Shivam", "English")
		want := "Hello, Shivam"
		assertCorrectMessage(t, got, want)

	})

	t.Run("Saying Hello, World when empty string is supplied.", func(t *testing.T) {
		got := hello("", "")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)

	})

	t.Run("Saying Hello in Spanish.", func(t *testing.T) {
		got := hello("Shivam", "Spanish")
		want := "Hola, Shivam"
		assertCorrectMessage(t, got, want)

	})

	t.Run("Saying Hello in French.", func(t *testing.T) {
		got := hello("Shivam", "French")
		want := "Bonjour, Shivam"
		assertCorrectMessage(t, got, want)
	})

	t.Run("Saying Hello name when empty language is supplied.", func(t *testing.T) {
		got := hello("Shivam", "")
		want := "Hello, Shivam"
		assertCorrectMessage(t, got, want)

	})
}
