package timeline

import (
	"fmt"
	"io"
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

func AssertFacets(t *testing.T, manager *TimelineManager, want FacetSet) {
    t.Helper()

    got := manager.Facets()

    if got, want := got.idle, want.idle; !reflect.DeepEqual(got, want) {
        t.Errorf("got idle facet %#v, want %#v. timelines=%v", got, want, manager.timelines)
    }

    if got, want := got.dashed[START], want.dashed[START]; !reflect.DeepEqual(got, want) {
        t.Errorf("got dashed START facet %#v, want %#v. timelines=%v", got, want, manager.timelines)
    }

    if got, want := got.dashed[INPROGRESS], want.dashed[INPROGRESS]; !reflect.DeepEqual(got, want) {
        t.Errorf("got dashed INPROGRESS facet %#v, want %#v. timelines=%v", got, want, manager.timelines)
    }

    if got, want := got.dashed[LASTWORDS], want.dashed[LASTWORDS]; !reflect.DeepEqual(got, want) {
        t.Errorf("got dashed LASTWORDS facet %#v, want %#v. timelines=%v", got, want, manager.timelines)
    }
}

func AssertSink(t *testing.T, sink fmt.Stringer, want string) {
    t.Helper()

    got := sink.String()

    gotLines := strings.Split(got, "\n")
    wantLines := strings.Split(want, "\n")

    commonSize := len(gotLines)
    if len(wantLines) < commonSize {
        commonSize = len(wantLines)
    }

    for i := 0; i < commonSize; i++ {
        if gotLines[i] != wantLines[i] {
            t.Errorf("got sink's line %d %q, want %q", i + 1, gotLines[i], wantLines[i])
        }
    }

    for i := commonSize; i < len(gotLines); i++ {
        t.Errorf("got sink's line %d %q, want none", i + 1, gotLines[i])
    }

    for i := commonSize; i < len(wantLines); i++ {
        t.Errorf("expected sink's line %d %q, got none", i + 1, wantLines[i])
    }
}

func AssertSinks(t *testing.T, sinks []io.Writer, wants ...string) {
    t.Helper()

    for i, sink := range sinks {
        buf, ok := sink.(fmt.Stringer)
        if !ok {
            t.Errorf("sink #%d's output cannot be retrieved", i)
            continue
        }

        got := buf.String()
        if got != wants[i] {
            t.Errorf("got sink #%d's output %q, want %q", i, got, wants[i])
        }
    }
}
