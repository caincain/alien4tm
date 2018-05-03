package main

import (
	"fmt"

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
	// figure out what the screensize is:
	w, h := termbox.Size()
	const coldef = termbox.ColorDefault

	for i := 6; i < w-6; i++ {
		for j := 6; j < h-6; j++ {
			termbox.SetCell(i, j, ' ', termbox.ColorYellow, termbox.ColorYellow)
		}
	}
	termbox.SetCursor(1, 1)
	//
	for yy := max_y; yy >= min_y; yy-- {
		for xx := min_x; xx <= max_x; xx++ {
			if _, ok := worldMap[coordinates{xx, yy}]; ok {
				termbox.SetCell(8+xx, 8+yy, '[', termbox.ColorBlue, termbox.ColorWhite)
				termbox.SetCell(8+xx+1, 8+yy, ' ', termbox.ColorBlue, termbox.ColorWhite)
				termbox.SetCell(8+xx+2, 8+yy, ' ', termbox.ColorBlue, termbox.ColorWhite)
				termbox.SetCell(8+xx+3, 8+yy, ' ', termbox.ColorBlue, termbox.ColorWhite)
				termbox.SetCell(8+xx+4, 8+yy, ']', termbox.ColorBlue, termbox.ColorWhite)
			} else {

			}
		}
		fmt.Println()
	}

	termbox.Flush()
	return
}
