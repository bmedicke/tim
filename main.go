package main

import (
	"flag"
	"fmt"
	"log"
	"time"

	"github.com/gdamore/tcell/v2"
)

// keep track of start time:
var start = time.Now()

func main() {
	// register command line flags:
	stopwatchMode := flag.Bool("s", false, "stopwatch mode")
	countdownMode := flag.Bool("c", false, "countdown mode")
	flag.Parse()

	if *stopwatchMode {
		stopwatch()
	} else if *countdownMode {
		fmt.Println("countdown mode not yet implemented")
	} else {
		flag.PrintDefaults()
	}
}

// create and initialize new screen:
func setupScreen() tcell.Screen {
	screen, err := tcell.NewScreen()
	if err != nil {
		log.Fatalf("%+v", err)
	}
	if err := screen.Init(); err != nil {
		log.Fatalf("%+v", err)
	}
	return screen
}
