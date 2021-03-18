package timeline

type TimelineStyle struct {
    Pipe string
    Dash string
    Space string

    DashedPipe map[TimelineState]string
    Tick map[TimelineState]string
    Postfix map[TimelineState]string
}

var StyleASCII = &TimelineStyle{
    Pipe: "|",
    Dash: "-",
    Space: " ",
    DashedPipe: map[TimelineState]string{
        START: "+",
        INPROGRESS: "+",
        LASTWORDS: "+",
    },
    Tick: map[TimelineState]string{
        START: "/",
        INPROGRESS: "}",
        LASTWORDS: "\\",
    },
    Postfix: map[TimelineState]string{
        START: "-- ",
        INPROGRESS: "- ",
        LASTWORDS: "__ ",
    },
}

var StyleDefault = &TimelineStyle{
    Pipe: "│",
    Dash: "─",
    Space: " ",
    DashedPipe: map[TimelineState]string{
        START: "┼",
        INPROGRESS: "┼",
        LASTWORDS: "┼",
    },
    Tick: map[TimelineState]string{
        START: "┌",
        INPROGRESS: "├",
        LASTWORDS: "└",
    },
    Postfix: map[TimelineState]string{
        START: "─ ",
        INPROGRESS: "╴ ",
        LASTWORDS: "─ ",
    },
}

var StyleEdgesBold = &TimelineStyle{
    Pipe: "│",
    Dash: "─",
    Space: " ",
    DashedPipe: map[TimelineState]string{
        START: "╪",
        INPROGRESS: "┼",
        LASTWORDS: "╪",
    },
    Tick: map[TimelineState]string{
        START: "╒",
        INPROGRESS: "├",
        LASTWORDS: "╘",
    },
    Postfix: map[TimelineState]string{
        START: "═",
        INPROGRESS: "╴ ",
        LASTWORDS: "═",
    },
}

var StyleRounded = &TimelineStyle{
    Pipe: "│",
    Dash: "─",
    Space: " ",
    DashedPipe: map[TimelineState]string{
        START: "┼",
        INPROGRESS: "┼",
        LASTWORDS: "┼",
    },
    Tick: map[TimelineState]string{
        START: "╭",
        INPROGRESS: "├",
        LASTWORDS: "╰",
    },
    Postfix: map[TimelineState]string{
        START: "─ ",
        INPROGRESS: "╴ ",
        LASTWORDS: "─ ",
    },
}
