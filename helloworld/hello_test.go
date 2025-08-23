package main

import "testing"

// func TestHello(t *testing.T) {
//     got := Hello("John")
//     want := "Hello, John"

//     if got != want {
//         t.Errorf("got %q want %q", got, want)
//     }
// }

// func TestHello(t *testing.T) {
//     t.Run("saying hello to people", func(t *testing.T) {
//         got := Hello("John")
//         want := "Hello, John"

//         if got != want {
//             t.Errorf("got %q want %q", got, want)
//         }
//     })
//     t.Run("say 'Hello, World' when an empty string is supplied", func(t *testing.T) {
//         got := Hello("")
//         want := "Hello, World"

//         if got != want {
//             t.Errorf("got %q want %q", got, want)
//         }
//     })
// }

func TestHello(t *testing.T) {
    t.Run("saying hello to people", func(t *testing.T) {
        got := Hello("Chris", "")
        want := "Hello, Chris"
        assertCorrectMessage(t, got, want)
    })

    t.Run("empty string defaults to 'world'", func(t *testing.T) {
        got := Hello("", "")
        want := "Hello, World"
        assertCorrectMessage(t, got, want)
    })
    t.Run("in Spanish", func(t *testing.T) {
        got := Hello("Lucia", "Spanish")
        want := "Hola, Lucia"
        assertCorrectMessage(t, got, want)
    })
    t.Run("in French", func(t *testing.T) {
        got := Hello("Louise", "French")
        want := "Bonjour, Louise"
        assertCorrectMessage(t, got, want)
    })
}

func assertCorrectMessage(t testing.TB, got, want string) {
    // tell the test suite that this method is a helper
    t.Helper()

    if got != want {
        t.Errorf("got %q want %q", got, want)
    }
}
