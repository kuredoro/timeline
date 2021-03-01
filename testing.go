package timeline

import (
	"reflect"
	"strings"
	"testing"
)

func AssertTimelines(t *testing.T, got, want []int) {
    t.Helper()

    if !reflect.DeepEqual(got, want) {
        t.Errorf("got timelines %v, want %v", got, want)
    }
}

func AssertFacetsOfSingleChars(t *testing.T, manager *TimelineManager, wantIdleStr, wantCrossedStr string) {
    t.Helper()

    wantIdle := strings.Split(wantIdleStr, "")
    wantCrossed := strings.Split(wantCrossedStr, "")

    if len(wantIdle) == 0 {
        wantIdle = nil
    }

    if len(wantCrossed) == 0 {
        wantCrossed = nil
    }

    gotIdle, gotCrossed := manager.Facets()

    if !reflect.DeepEqual(gotIdle, wantIdle) {
        t.Errorf("got idle facet %#v, want %#v. timelines=%v", gotIdle, wantIdle, manager.timelines)
    }

    if !reflect.DeepEqual(gotCrossed, wantCrossed) {
        t.Errorf("got crossed facet %#v, want %#v. timelines=%v", gotCrossed, wantCrossed, manager.timelines)
    }
}
