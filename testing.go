package timeline

import (
    "testing"
    "reflect"
)

func AssertTimelines(t *testing.T, got, want []int) {
    t.Helper()

    if !reflect.DeepEqual(got, want) {
        t.Errorf("got timelines %v, want %v", got, want)
    }
}
