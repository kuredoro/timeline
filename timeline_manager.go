package timeline

// TimeManager will have sinks of type FormattedSink
// It will have a func print(header, message)
// It will be able to text wrap message how ever it likes
// The sinks will also have Colorizers that given the column and the id of the timeline
// will output corresponding color for the bar...
// I.e., transform the symbol for bar...
//

type TimelineManager struct {
    timelines []int
    lastTimelineID int
}

func (tm *TimelineManager) Spawn() bool {
    firstZero := 0
    for ; firstZero < len(tm.timelines); firstZero++ {
        if tm.timelines[firstZero] == 0 {
            break
        }
    }

    tm.lastTimelineID++

    if firstZero != len(tm.timelines) {
        tm.timelines[firstZero] = tm.lastTimelineID
    } else {
        tm.timelines = append(tm.timelines, tm.lastTimelineID)
    }

    return false
}
