package timeline

import "fmt"

type TimelineState int

const (
    c_start TimelineState = iota
    c_inprogress
    c_lastwords
    c_finished
)

type Timeline struct {
    manager *TimelineManager
    style *TimelineStyle

    state TimelineState
}

func (t *Timeline) Println(args ...interface{}) {
    msg := fmt.Sprintln(args...)

    if t.state == c_start {
        t.manager.print([]string{t.style.StartTick, t.style.Dash, t.style.Space}, msg)
        t.state = c_inprogress
    } else if t.state == c_inprogress {
        t.manager.print([]string{t.style.InterTick, t.style.Dash, t.style.Space}, msg)
    } else if t.state == c_lastwords {
        t.manager.print([]string{t.style.FinalTick, t.style.Dash, t.style.Space}, msg)
        t.state = c_finished
    }
}

func (t *Timeline) End() {
    if t.state == c_finished {
        return
    }

    t.state = c_lastwords
}
