package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"reflect"
	"time"
)

//func gStateToBoard(gState [3][3]rune) (string, error) {
//	// temporary only works for fixed 3 by 3 boards
//	for i, j := range gState {
//
//	}
//}

func main() {
	cfg, err := loadConfig("config.toml")
	if err != nil {
		log.Fatal(err)
	}

	// Temporary debug log RM LATER
	_, val := reflect.TypeOf(cfg.Board.Offset), reflect.ValueOf(cfg.Board.Offset)
	fmt.Println("This is the first character offset: ", val)
	fmt.Printf("This is the small board path: %s\n", cfg.Board.PathS)
	fmt.Printf("This is the great board path: %s\n", cfg.Board.PathG)

	// Load the raw empty board from the text file into an array of bytes
	file, err := os.Open(cfg.Board.PathS)
	if err != nil {
		return
	}
	f := bufio.NewScanner(file)
	activeBoard, err := loadBoard(f, "", "Xwin")
	if err != nil {
		log.Fatal(err)
	}

	// DBG PRINT RM later

	fmt.Print("\x1b[2J\x1b[H") // clear screen, cursor to 1;1
	fmt.Print(activeBoard)
	time.Sleep(time.Second)

	activeBoardXwin, err := loadBoard(f, "Xwin", "Owin")
	fmt.Print("\x1b[H") // back to top-left
	fmt.Print(activeBoardXwin)
	time.Sleep(time.Second)

	activeBoardOwin, err := loadBoard(f, "Owin", "BIG")
	fmt.Print("\x1b[H") // back to top-left
	fmt.Print(activeBoardOwin)
	time.Sleep(time.Second)

	activeBoardBig, err := loadBoard(f, "BIG", "end")
	fmt.Print("\x1b[H") // back to top-left
	fmt.Print(activeBoardBig)
	time.Sleep(time.Second)

	var inter Interface
	var state State
	state.Init()
	inter.Init(
		activeBoard,
		activeBoardOwin,
		activeBoardXwin,
		activeBoardBig,
	)

	inter.renderBoard(state)
	fmt.Print("\x1b[H") // back to top-left
	fmt.Print(inter.RenderedBoard)

	fmt.Println() // move to a new line when done
}
