package main

import (
	"time"

	termbox "github.com/nsf/termbox-go"
)

// bunch of helper functions
func fill(x, y, w, h int, cell termbox.Cell) {
	for ly := 0; ly < h; ly++ {
		for lx := 0; lx < w; lx++ {
			termbox.SetCell(x+lx, y+ly, cell.Ch, cell.Fg, cell.Bg)
		}
	}
}

// printMapGraphic prints a world map graphically
func printMapGraphic(worldMap map[coordinates]*city, min_x, min_y, max_y, max_x int) {
	// figure out what info about terminal
	w, h := termbox.Size()
	margin := 10 * h / 100 // margin: 10% of width

	// we want the cursor
	termbox.SetCursor(1, 1)

	//
	wCell := (w/2 - 2*margin) / max_x
	hCell := (h/2 - 2*margin) / max_y

	// set cities or non-cities
	for yy := max_y; yy >= min_y; yy-- {
		for xx := min_x; xx <= max_x; xx++ {
			if _, ok := worldMap[coordinates{xx, yy}]; ok {
				// func fill(x, y, w, h int, cell termbox.Cell) {
				fill(margin+(xx*wCell), margin+(yy*hCell), wCell, hCell, termbox.Cell{Ch: 'X', Bg: termbox.ColorYellow})
			} else {
				fill(margin+(xx*wCell), margin+(yy*hCell), wCell, hCell, termbox.Cell{Bg: termbox.ColorBlue})
			}
			termbox.Flush()
			time.Sleep(50 * time.Millisecond)
		}
	}

	termbox.Flush()
}
