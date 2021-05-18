package main

import (
	"fmt"
	"time"

	"github.com/DraculK/Sudoku/sudoku"
)

func main() {
	puzzle := "4.....8.5.3..........7......2.....6.....8.4......1.......6.3.7.5..2.....1.4......"

	inicio := time.Now()
	sudoku.PrettyDisplayInit(puzzle)
	fmt.Println("Resolvendo..")
	resultado, err := sudoku.Resolve(puzzle)
	if err != nil {
		fmt.Println(err)
	}
	final := time.Since(inicio)
	fmt.Println(final)
	sudoku.PrettyDisplay(resultado)
}
