package iteration

import (
	"testing"
)

func TestRepeat(t *testing.T) {

	assertCorrectMessage := func(t testing.TB, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got: %q want := %q", got, want)
		}
	}

	t.Run("Checking if string and number of repetations is give.", func(t *testing.T) {

		repeated := Repeat("a", 0)
		want := ""
		assertCorrectMessage(t, repeated, want)
	})

	t.Run("Checking if string and number of repetations is give.", func(t *testing.T) {

		repeated := Repeat("aman", 6)
		want := "amanamanamanamanamanaman"
		assertCorrectMessage(t, repeated, want)
	})
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 10)
	}
}
