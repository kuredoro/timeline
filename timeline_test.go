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

    t.Run("2 timelines", func(t *testing.T) {
        buf := &bytes.Buffer{}

        manager := &timeline.TimelineManager{
            Style: timeline.StyleASCII,
            Sinks: []io.Writer{buf},
        }

        first := manager.Spawn()
        first.Println("hello")

        second := manager.Spawn()
        second.Println("Hi")

        first.Println("how's it goin'?")
        second.Println("Awesome! Go is awesome!")

        first.End()
        first.Println("i hate go. bye.")

        second.Println("Hmph.... ...What?")
        second.Println("If they cease communication with people based on the language they like....")
        second.Println("Then I'm blessed I won't have to deal with them.")

        third := manager.Spawn()
        third.Println("ugh... u stil here?")

        second.End()
        second.Println("No.")

        third.Println("damn...")

        third.End()
        third.Println("okay")

        want :=
`/- hello
|/- Hi
}+- how's it goin'?
|}- Awesome! Go is awesome!
\+- i hate go. bye.
 }- Hmph.... ...What?
 }- If they cease communication with people based on the language they like....
 }- Then I'm blessed I won't have to deal with them.
/+- ugh... u stil here?
|\- No.
}- damn...
\- okay
`

        timeline.AssertSink(t, buf, want)
    })
}
