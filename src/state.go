package main

import (
	"errors"
	"image"
)

type Move struct {
	X, Y   int
	Player bool
}

var (
	GameOngoing int8 = 0
	OisWinner   int8 = 1
	XisWinner   int8 = 2
	Draw        int8 = 3
)

type State struct {
	StateBoard     [9][9]int8  // StateBoard: 0 for empty, 1 for O, 2 for X, init with all 0
	ExpectedPlayer bool        // who played last, init with O == false
	ExpectedSquare image.Point // What point ({0..2}, {0..2}) it is expected the next move to be in
	Winner         int8        // init 0, 1 for O, 2 for X, 3 for draw?
	overview       [3][3]int8  // board overview tracks the small tic tacs
}

func (state *State) Init() {
	for i := range state.StateBoard {
		for j := range state.StateBoard[i] {
			state.StateBoard[i][j] = 0
		}
	}

	state.ExpectedPlayer = false
	state.ExpectedSquare = image.Point{-1, -1}
	state.Winner = GameOngoing

	for i := range state.overview {
		for j := range state.overview {
			state.overview[i][j] = 0
		}
	}
}

var (
	ErrOutOfRange  = errors.New("move out of range")
	ErrWrongTurn   = errors.New("wrong player's turn")
	ErrWrongSquare = errors.New("invalid move square")
	ErrGameOver    = errors.New("game over")
)

func TicTacWon3x3(ttt3x3 [3][3]int8) int8 {
	// check rows
	for i := range 3 {
		if ttt3x3[i][0] == ttt3x3[i][1] && ttt3x3[i][1] == ttt3x3[i][2] {
			return ttt3x3[0][i]
		}
	}

	// check columns
	for j := range 3 {
		if ttt3x3[0][j] == ttt3x3[1][j] && ttt3x3[1][j] == ttt3x3[2][j] {
			return ttt3x3[0][j]
		}
	}

	//check diagonals
	if ttt3x3[1][1] == ttt3x3[0][2] && ttt3x3[1][1] == ttt3x3[2][0] {
		return ttt3x3[1][1]
	}
	if ttt3x3[1][1] == ttt3x3[0][0] && ttt3x3[1][1] == ttt3x3[2][2] {
		return ttt3x3[1][1]
	}

	filled := true
	for i := range 3 {
		for j := range 3 {
			if ttt3x3[i][j] == 0 {
				filled = false
				break
			}
		}

		if !filled {
			break
		}
	}

	if filled {
		return Draw
	}
	return GameOngoing
}

func makeSubBoard(board [9][9]int8, from image.Point) [3][3]int8 {
	var ttt3x3 [3][3]int8
	for i := range 3 {
		for j := range 3 {
			ttt3x3[i][j] = board[i+from.X][j+from.Y]
		}
	}

	return ttt3x3
}

func (s *State) updateOverview() {
	var overview [3][3]int8 // We build this and check for winner there
	for i := range 3 {
		for j := range 3 {
			subboard := makeSubBoard(s.StateBoard, image.Point{X: 3 * i, Y: 3 * j})
			subResult := TicTacWon3x3(subboard)
			if subResult != OisWinner && subResult != XisWinner {
				// remove carry over info  like draw and game ongoing
				subResult = 0
			}
			overview[i][j] = TicTacWon3x3(subboard)
		}
	}
	s.overview = overview
}

func (s *State) IsLegalMove(mv Move) error {
	if mv.X < 0 || mv.Y < 0 || mv.X > 8 || mv.Y > 8 { //check if out of range
		return ErrOutOfRange
	}
	if mv.Player != s.ExpectedPlayer {
		return ErrWrongTurn
	}
	// mv.X / 3 gives the X in square 0..2 or if square is occupied
	if mv.X/3 != s.ExpectedSquare.X || mv.Y/3 != s.ExpectedSquare.Y ||
		s.StateBoard[mv.X][mv.Y] != 0 {
		return ErrWrongSquare
	}
	if s.Winner != 0 {
		return ErrGameOver
	}
	return nil
}

func (s *State) updateWinner() {
	s.Winner = TicTacWon3x3(s.overview)
}

func (s *State) PlayMove(mv Move) error {
	s.StateBoard[mv.X][mv.Y] = 1 // by default O
	if mv.Player == true {
		s.StateBoard[mv.X][mv.Y]++
	}

	s.ExpectedPlayer = !s.ExpectedPlayer // Update expected player
	s.ExpectedSquare = image.Point{X: mv.X / 3, Y: mv.Y / 3}
	s.updateOverview()
	s.updateWinner()

	return s.IsLegalMove(mv)
}
