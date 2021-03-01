package timeline

import (
    "testing"
)

func TestTimelineManagerSpawn(t *testing.T) {
    t.Run("spawn when no timelines present", func(t *testing.T) {
        manager := &TimelineManager{}

        manager.Spawn()

        want := []int{1}
        AssertTimelines(t, manager.timelines, want)
    })

    t.Run("spawn two timelines", func(t *testing.T) {
        manager := &TimelineManager{}

        manager.Spawn()
        manager.Spawn()

        want := []int{1, 2}
        AssertTimelines(t, manager.timelines, want)
    })

    t.Run("spawn one when there's a hole between timelines", func(t *testing.T) {
        manager := &TimelineManager{
            timelines: []int{1, 2, 0, 4},
            lastTimelineID: 4,
        }

        manager.Spawn()

        want := []int{1, 2, 5, 4}
        AssertTimelines(t, manager.timelines, want)
    })
}

func TestTimelineManagerDestroy(t *testing.T) {
    t.Run("spawn, destroy", func (t *testing.T) {
        manager := &TimelineManager{}

        manager.Spawn()
        manager.destroy(1)

        AssertTimelines(t, manager.timelines, []int{})
    })

    t.Run("contract holes that are exposed to the right end", func (t *testing.T) {
        manager := &TimelineManager{}

        manager.Spawn()
        manager.Spawn()
        manager.Spawn()
        manager.destroy(2)

        AssertTimelines(t, manager.timelines, []int{1, 0, 3})

        manager.destroy(3)

        AssertTimelines(t, manager.timelines, []int{1})
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
