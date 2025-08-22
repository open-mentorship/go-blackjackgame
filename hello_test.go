package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("Greeting people", func(t *testing.T) {
		got := Hello("world")
		want := "Hello world!"

		if got != want {
			t.Errorf("got %q but want %q", got, want)
		}
	})
	t.Run("Greet 'Hello world!' when empty string is supplied", func(t *testing.T) {
		got := Hello("")
		want := "Hello world!"

		if got != want {
			t.Errorf("got %q but want %q", got, want)
		}
	})
}
