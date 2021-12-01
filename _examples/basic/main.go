package main

import (
	"fmt"
	"io"
	"os"

	"github.com/kuredoro/timeline"
)

func main() {
	manager := &timeline.TimelineManager{
		Style: timeline.StyleWide,
		Sinks: []io.Writer{os.Stdout},
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

	fmt.Println("")

    // ---

	first = manager.Spawn()
	first.Println("a")

	second = manager.Spawn()
	second.Println("b")

	third = manager.Spawn()
	third.Println("c")

	first.Println("d")

	second.End()
	second.Println("e")

	first.Println("f")
	third.Println("z")

	second2 := manager.Spawn()
	second2.Println("g")

	first.End()
	first.Println("h")

	second2.Println("i")

	third.End()
	third.Println("j")

	second2.End()
	second2.Println("k")
}
