package timeline

type TimelineStyle struct {
    Pipe string
    DashedPipe string
    Dash string
    Space string
    StartTick string
    InterTick string
    FinalTick string

    Tick map[TimelineState]string
    Postfix map[TimelineState]string
}

var StyleASCII = &TimelineStyle{
    Pipe: "|",
    DashedPipe: "+",
    Dash: "-",
    Space: " ",
    Tick: map[TimelineState]string{
        START: "/",
        INPROGRESS: "}",
        LASTWORDS: "\\",
    },
    Postfix: map[TimelineState]string{
        START: "- ",
        INPROGRESS: "- ",
        LASTWORDS: "- ",
    },
}
