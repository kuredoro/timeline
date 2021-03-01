package timeline

type TimelineStyle struct {
    Pipe string
    DashedPipe string
    Dash string
    Space string
    StartTick string
    InterTick string
    FinalTick string
}

var StyleASCII = &TimelineStyle{
    Pipe: "|",
    DashedPipe: "+",
    Dash: "-",
    Space: " ",
    StartTick: "/",
    InterTick: "}",
    FinalTick: "\\",
}
