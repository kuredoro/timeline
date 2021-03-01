package timeline_test

import (
	"bytes"
	"io"
	"testing"

	"github.com/kuredoro/timeline"
)

func TestTimelineUseCases(t *testing.T) {
    t.Run("spawn, print something, die", func(t *testing.T) {
        buf := &bytes.Buffer{}

        manager := &timeline.TimelineManager{
            Style: timeline.StyleASCII,
            Sinks: []io.Writer{buf},
        }

        tl := manager.Spawn()

        tl.Println("start", 1)
        timeline.AssertSink(t, buf, "/- start 1\n")
        buf.Reset()

        tl.Println("middle", 2)
        timeline.AssertSink(t, buf, "}- middle 2\n")
        buf.Reset()

        tl.Println("another middle", 3)
        timeline.AssertSink(t, buf, "}- another middle 3\n")
        buf.Reset()

        tl.End()
        tl.Println("last words")
        timeline.AssertSink(t, buf, "\\- last words\n")
        buf.Reset()

        tl.Println("ignore me")
        timeline.AssertSink(t, buf, "")

        tl.End()
        tl.Println("ignore me also, hehe")
        timeline.AssertSink(t, buf, "")
    })
}
