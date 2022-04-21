package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gdamore/tcell"
)

func main() {
	// keep track of start time:
	start := time.Now()
	stopwatch(start)
}

func stopwatch(start time.Time) {
	// create and initialize new screen:
	s, err := tcell.NewScreen()
	if err != nil {
		log.Fatalf("%+v", err)
	}
	if err := s.Init(); err != nil {
		log.Fatalf("%+v", err)
	}

	// start event loop:
	for {
		s.Show()            // update screen.
		ev := s.PollEvent() // fetch event.

		// handle event:
		switch ev := ev.(type) {
		case *tcell.EventKey:
			if ev.Rune() == 'q' || ev.Key() == tcell.KeyEsc {
				// stop screen so the log message persist:
				s.Fini()
				// calculate and log timedelta:
				fmt.Println(time.Now().Sub(start))
				// terminate the program:
				os.Exit(0)
			}
		case *tcell.EventResize:
			s.Sync()
			// print instructions:
			instruction := "stopwatch running, press escape or q to quit"
			fmt.Println(instruction)
		}
	}
}
