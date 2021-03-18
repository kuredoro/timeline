package timeline

import (
    "fmt"
    "io"
)

type FacetSet struct {
    idle []string
    dashed map[TimelineState][]string
}

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

    Style *TimelineStyle
    facets FacetSet

    Sinks []io.Writer
}

func (tm *TimelineManager) Spawn() *Timeline {
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

    tm.generateFacets()

    return &Timeline{
        manager: tm,
        style: tm.Style,
        column: firstZero,
    }
}

func (tm *TimelineManager) destroy(column int) {
    tm.timelines[column] = 0

    lastNonZero := len(tm.timelines) - 1
    for lastNonZero >= 0 && tm.timelines[lastNonZero] == 0 {
        lastNonZero--
    }

    tm.timelines = tm.timelines[:lastNonZero+1]

    tm.generateFacets()
}

func (tm *TimelineManager) Facets() FacetSet {
    return tm.facets
}

func (tm *TimelineManager) generateFacet(present, absent string) []string {
    str := make([]string, len(tm.timelines))

    for i, id := range tm.timelines {
        if id == 0 {
            str[i] = absent
        } else {
            str[i] = present
        }
    }

    return str
}

func (tm *TimelineManager) generateFacets() {
    if tm.Style == nil {
        return
    }

    tm.facets.idle = tm.generateFacet(tm.Style.Pipe, tm.Style.Space)

    if tm.facets.dashed == nil {
        tm.facets.dashed = make(map[TimelineState][]string)
    }

    for k, v := range tm.Style.DashedPipe {
        tm.facets.dashed[k] = tm.generateFacet(v, tm.Style.Dash)
    }
}

func (tm *TimelineManager) print(timelineHeader []string, msg string) {
    for _, sink := range tm.Sinks {
        for _, header := range timelineHeader {
            fmt.Fprint(sink, header)
        }

        fmt.Fprintf(sink, msg)
    }
}
