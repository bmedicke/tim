package main

import (
	"os"
	"time"

	"github.com/gdamore/tcell/v2"
)

func countdown(duration *time.Duration) {
	screen := setupScreen()
	redrawTicker := time.NewTicker(redrawDelay)

	go handleCountdownEvents(screen)
	time.AfterFunc(*duration, func() { exitCountdownWith(0, screen) })

	for range redrawTicker.C {
		if !*quietFlag {
			screen.Show() // update screen.
			durationLeft := (*duration - time.Now().Sub(start)).String()
			drawText(screen, 0, 1, 30, 2, durationLeft)
			drawText(screen, 0, 0, 30, 1, instructions)
		}
	}
}

func handleCountdownEvents(screen tcell.Screen) {
	for {
		ev := screen.PollEvent() // wait for and fetch event.

		// process event:
		switch ev := ev.(type) {
		case *tcell.EventKey:
			if ev.Rune() == 'q' {
				if *quitWithSuccessFlag {
					exitCountdownWith(0, screen)
				} else {
					exitCountdownWith(1, screen)
				}
			}
			if ev.Rune() == 'd' {
				*quietFlag = !*quietFlag
				screen.Clear()
				screen.Sync()
			}
			if ev.Key() == tcell.KeyCtrlC {
				exitCountdownWith(127, screen)
			}
		case *tcell.EventResize:
			screen.Sync()
		}
	}
}

func exitCountdownWith(code int, screen tcell.Screen) {
	screen.Fini()
	os.Exit(code)
}
