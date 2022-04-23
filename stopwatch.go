package main

import (
	"fmt"
	"os"
	"time"

	"github.com/gdamore/tcell/v2"
)

func stopwatch() {
	const stopwatchInstructions = "q: quit, d: toggle display"
	screen := setupScreen()

	go handleStopwatchEvents(screen)

	for {
		if !*quietFlag {
			screen.Show() // update screen.
			duration := fmt.Sprintf("%s", time.Now().Sub(start))
			drawText(screen, 0, 1, 30, 2, duration)
			drawText(screen, 0, 0, 30, 1, stopwatchInstructions)
		}
		time.Sleep(time.Millisecond * 50)
	}
}

func handleStopwatchEvents(screen tcell.Screen) {
	for {
		ev := screen.PollEvent() // wait for and fetch event.

		// process event:
		switch ev := ev.(type) {
		case *tcell.EventKey:
			if ev.Rune() == 'q' {
				exitStopwatchWith(0, screen)
			}
			if ev.Rune() == 'd' {
				*quietFlag = !*quietFlag
				screen.Clear()
				screen.Sync()
			}
			if ev.Key() == tcell.KeyCtrlC {
				exitStopwatchWith(127, screen)
			}
		case *tcell.EventResize:
			screen.Sync()
		}
	}
}

func exitStopwatchWith(code int, screen tcell.Screen) {
	// stop screen first so the log message persists:
	screen.Fini()
	// calculate and log final timedelta:
	fmt.Println(time.Now().Sub(start))
	os.Exit(code)
}
