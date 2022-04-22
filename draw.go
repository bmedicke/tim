package main

import "github.com/gdamore/tcell/v2"

func drawText(s tcell.Screen, x1, y1, x2, y2 int, text string) {
	col := x1
	row := y1
	defStyle := tcell.StyleDefault.Background(tcell.ColorReset).
		Foreground(tcell.ColorWhite)

	for _, r := range []rune(text) {
		s.SetContent(col, row, r, nil, defStyle)
		col++
		if col >= x2 {
			row++    // jump to next row.
			col = x1 // carriage return.
		}
		if row > y2 {
			break
		}
	}
}
