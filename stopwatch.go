package main

import (
	"fmt"
	"os"
	"time"

	"github.com/gdamore/tcell/v2"
)

const stopwatchInstructions = "press q or esc to stop"

func stopwatch(quietFlag bool) {
	screen := setupScreen()

	for {
		if screen.HasPendingEvent() {
			handleStopwatchEvent(screen)
		} else {
			if !quietFlag {
				screen.Show() // update screen.
				duration := fmt.Sprintf("%s", time.Now().Sub(start))
				drawText(screen, 0, 1, 30, 2, duration)
				drawText(screen, 0, 0, 30, 1, stopwatchInstructions)
			}
		}
		time.Sleep(time.Millisecond * 50)
	}
}

func handleStopwatchEvent(screen tcell.Screen) {
	ev := screen.PollEvent() // fetch event.

	// process event:
	switch ev := ev.(type) {
	case *tcell.EventKey:
		if ev.Rune() == 'q' || ev.Key() == tcell.KeyEsc {
			exitStopwatchWith(0, screen)
		}
		if ev.Key() == tcell.KeyCtrlC {
			exitStopwatchWith(127, screen)
		}
	case *tcell.EventResize:
		screen.Sync()
	}
}

func exitStopwatchWith(code int, screen tcell.Screen) {
	// stop screen first so the log message persists:
	screen.Fini()
	// calculate and log final timedelta:
	fmt.Println(time.Now().Sub(start))
	os.Exit(code)
}
