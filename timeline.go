package timeline

import "fmt"

type TimelineState int

const (
    START TimelineState = iota
    INPROGRESS
    LASTWORDS
    FINISHED
)

type Timeline struct {
    manager *TimelineManager
    style *TimelineStyle

    column int
    state TimelineState
}

func (t *Timeline) Println(args ...interface{}) {
    if t.state == FINISHED {
        return
    }

    tick := t.style.Tick[t.state]
    postfix := t.style.Postfix[t.state]

    facets := t.manager.Facets()

    header := make([]string, t.column)
    copy(header, facets.idle[:t.column])
    header = append(header, tick)
    header = append(header, facets.dashed[t.state][t.column+1:]...)

    msg := fmt.Sprintln(args...)

    t.manager.print(header, postfix + msg)

    if t.state == START {
        t.state = INPROGRESS
    } else if t.state == LASTWORDS {
        t.state = FINISHED
        t.manager.destroy(t.column)
    }
}

func (t *Timeline) End() {
    if t.state == FINISHED {
        return
    }

    t.state = LASTWORDS
}
