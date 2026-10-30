package main

import (
	"image"
	"strings"

	tea "charm.land/bubbletea/v2"
)

// for bubble tea
type model struct {
	choices  []string         // items on the to-do list
	cursor   int              // which to-do list item our cursor is pointing at
	selected map[int]struct{} // which to-do items are selected
}

// aux functions
func strToRunes2D(s string) [][]rune {
	lines := strings.Split(s, "\n")
	grid := make([][]rune, len(lines))
	for j, line := range lines {
		grid[j] = []rune(line)
	}

	return grid
}

func runes2DToStr(grid [][]rune) string {
	var sb strings.Builder
	for i, row := range grid {
		if i > 0 {
			sb.WriteByte('\n')
		}
		sb.WriteString(string(row))
	}
	return sb.String()
}

type Interface struct {
	// UNIT of offset chars per box in big board
	offsetUnit    image.Point
	userInterface string
	RenderedBoard string
	smallBoard    [][]rune
	oWinBoard     [][]rune
	xWinBoard     [][]rune
	bigBoard      [][]rune
}

func (in *Interface) Init(small, owin, xwin, big string) {
	in.offsetUnit = image.Point{X: 6, Y: 14}
	in.smallBoard = strToRunes2D(small)
	in.oWinBoard = strToRunes2D(owin)
	in.xWinBoard = strToRunes2D(xwin)
	in.bigBoard = strToRunes2D(big)

	in.RenderedBoard = ""
}

func (in *Interface) placeOnBig(small [][]rune, place image.Point) {
	offset := image.Point{X: place.X * in.offsetUnit.X, Y: place.Y * in.offsetUnit.Y}
	for i := range small {
		for j := range small[i] {
			in.bigBoard[i+offset.X][j+1+offset.Y] = in.smallBoard[i][j]
		}
	}
}

func (in *Interface) renderBoard(state State) {
	// Place all boards on big board according to state overview
	for i := range state.overview {
		for j := range state.overview[i] {
			if state.overview[i][j] == OisWinner {
				in.placeOnBig(in.oWinBoard, image.Point{X: i, Y: j})
			} else if state.overview[i][j] == XisWinner {
				in.placeOnBig(in.xWinBoard, image.Point{X: i, Y: j})
			} else {
				in.placeOnBig(in.smallBoard, image.Point{X: i, Y: j})
			}
		}
	}

	// Place O and X or leave empty according to game state
	// if for ev
	for i := range state.StateBoard {
		for j := range state.StateBoard[i] {

		}
	}

	in.RenderedBoard = runes2DToStr(in.bigBoard)
}
