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

func AssertFacets(t *testing.T, manager *TimelineManager, wantIdle, wantCrossed string) {
    t.Helper()

    gotIdle, gotCrossed := manager.Facets()

    if gotIdle != wantIdle {
        t.Errorf("got idle facet %q, want %q. timelines=%v", gotIdle, wantIdle, manager.timelines)
    }

    if gotCrossed != wantCrossed {
        t.Errorf("got crossed facet %q, want %q. timelines=%v", gotCrossed, wantCrossed, manager.timelines)
    }
}
