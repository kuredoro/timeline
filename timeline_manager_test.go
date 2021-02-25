package timeline_test

import (
    "testing"

    "github.com/kuredoro/timeline"
)

func TestHello(t *testing.T) {
    t.Run("hello prints hello", func(t *testing.T) {
        got := timeline.Hello()
        want := "Hello"

        if got != want {
            t.Errorf("hello returned %q, want %q", got, want)
        }
    })
}
