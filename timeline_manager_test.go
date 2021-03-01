package timeline

import (
    "testing"
    "reflect"
)

func TestTimelineManagerSpawn(t *testing.T) {
    t.Run("spawn when no timelines present", func(t *testing.T) {
        manager := &TimelineManager{}

        manager.Spawn()

        want := []int{1}
        if !reflect.DeepEqual(manager.timelines, want) {
            t.Errorf("got timelines %v, want %v", manager.timelines, want)
        }
    })

    t.Run("spawn two timelines", func(t *testing.T) {
        manager := &TimelineManager{}

        manager.Spawn()
        manager.Spawn()

        want := []int{1, 2}
        if !reflect.DeepEqual(manager.timelines, want) {
            t.Errorf("got timelines %v, want %v", manager.timelines, want)
        }
    })

    t.Run("spawn one when there's a hole between timelines", func(t *testing.T) {
        manager := &TimelineManager{
            timelines: []int{1, 2, 0, 4},
            lastTimelineID: 4,
        }

        manager.Spawn()

        want := []int{1, 2, 5, 4}
        if !reflect.DeepEqual(manager.timelines, want) {
            t.Errorf("got timelines %v, want %v", manager.timelines, want)
        }
    })
}

func TestTimelineManagerDestroy(t *testing.T) {
    t.Run("spawn, destroy", func (t *testing.T) {
        manager := &TimelineManager{}

        manager.Spawn()
        manager.destroy(1)

        want := []int{}
        if !reflect.DeepEqual(manager.timelines, want) {
            t.Errorf("got timelines %v, want %v", manager.timelines, want)
        }
    })
}

func TestTimelineManagerFacets(t *testing.T) {
    /*
    style := TimelineStyle{
        Pipe: "|",
        CrossedPipe: "+",
        WideMinus: "-",
    }

        gotIdle, gotCrossed = manager.Facets()
        wantIdle, wantCrossed = "||||", "++++"

        if gotIdle != wantIdle {
            t.Errorf("got idle facet %q, want %q", gotIdle, wantIdle)
        }

        if gotCrossed != wantCrossed {
            t.Errorf("got crossed facet %q, want %q", gotCrossed, wantCrossed)
        }
        */
}

/*
func TestTimelineManagerColumnGaps(t *testing.T) {
}
*/
