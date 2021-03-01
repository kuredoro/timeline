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

/*
func TestTimelineManagerColumnGaps(t *testing.T) {
}
*/
