package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("Greeting everyone", func(t *testing.T) {
		got := Hello("world", "")
		want := "Hello world"

		assertCorrectMessage(t, got, want)
	})
	t.Run("Greet 'Hello world' when empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello world"
		assertCorrectMessage(t, got, want)
	})
	t.Run("Spanish greeting", func(t *testing.T) {
		got := Hello("world", spanish)
		want := "Hola world"
		assertCorrectMessage(t, got, want)
	})
	t.Run("French greeting", func(t *testing.T) {
		got := Hello("world", french)
		want := "Ca va world"
		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q but want %q", got, want)
	}
}
