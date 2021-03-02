package timeline_test

import (
	"bytes"
	"io"
	"testing"

	"github.com/kuredoro/timeline"
)

func TestTimelineUseCases(t *testing.T) {
    t.Run("1 timeline", func(t *testing.T) {
        buf := &bytes.Buffer{}

        manager := &timeline.TimelineManager{
            Style: timeline.StyleASCII,
            Sinks: []io.Writer{buf},
        }

        tl := manager.Spawn()

        tl.Println("start", 1)
        tl.Println("middle", 2)
        tl.Println("another middle", 1, 2, 3)

        tl.End()
        tl.Println("last words")

        tl.Println("ignore me")

        tl.End()
        tl.Println("ignore me also, hehe")

        want :=
`/- start 1
}- middle 2
}- another middle 1 2 3
\- last words
`

        timeline.AssertSink(t, buf, want)
    })
}
