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

    state TimelineState
}

func (t *Timeline) Println(args ...interface{}) {
    if t.state == FINISHED {
        return
    }

    tick := t.style.Tick[t.state]
    postfix := t.style.Postfix[t.state]

    msg := fmt.Sprintln(args...)

    t.manager.print([]string{tick}, postfix + msg)

    if t.state == START {
        t.state = INPROGRESS
    } else if t.state == LASTWORDS {
        t.state = FINISHED
    }
}

func (t *Timeline) End() {
    if t.state == FINISHED {
        return
    }

    t.state = LASTWORDS
}
