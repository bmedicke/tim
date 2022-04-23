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

var (
	quietFlag           = flag.Bool("q", false, helpQuietFlag)
	quitWithSuccessFlag = flag.Bool("x", false, helpQuitWithSuccess)
)

func main() {
	// register command line flags for modes (exclusive):
	countdownMode := flag.Bool("c", false, helpCountdownMode)
	stopwatchMode := flag.Bool("s", false, helpStopwatchMode)
	timerMode := flag.Bool("t", false, helpTimerMode)
	helpMode := flag.Bool("h", false, helpHelpMode)

	flag.Parse()

	if *helpMode {
		flag.PrintDefaults()
	} else if *stopwatchMode {
		stopwatch()
	} else if *countdownMode {
		fmt.Println("countdown mode not yet implemented")
	} else if *timerMode {
		fmt.Println("timer mode not yet implemented")
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
