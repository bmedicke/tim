package main

const (
	instructions      = "q: quit, d: toggle display"
	helpCountdownMode = `countdown mode:
pressing q exits with 1,
letting the timer run out exits with 0`
	helpHelpMode  = "show this help text"
	helpQuietFlag = `quiet flag, can also be toggled
by pressing d during runtime`
	helpQuitWithSuccess = `for countdown and timer mode:
pressing q exits with 0 instead of 1,
ctrl-c still exits with 127`
	helpStopwatchMode = `stopwatch mode:
pressing q exits with 0`
	helpTimerMode = `timer mode (count up):
pressing q exits with 1,
letting the timer run out exits with 0`
	helpVerboseFlag = "show instructions while running"
	helpVersionMode = "show version"
)
