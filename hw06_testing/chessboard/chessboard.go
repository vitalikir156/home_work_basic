package chessboard

import (
	"fmt"
)

var fullchessboard string

func Manual() {
	var howmuch int
	var gm bool
	fmt.Printf("Enter length: ")
	n, err := fmt.Scanln(&howmuch)
	if err != nil {
		fmt.Println(n, err)
		fmt.Println("Using default value (8)")
		howmuch = 8
	}
	fmt.Printf("Use doublespaces? ")
	n, err = fmt.Scanln(&gm)
	if err != nil {
		fmt.Println(n, err)
		fmt.Println("Using default value (no)")
		_, _ = Auto(howmuch, gm)
	}
}

func Auto(howmuch int, gm bool) (string, error) {
	fullchessboard = ""
	if howmuch < 2 {
		return "", fmt.Errorf("too small integer: %v", howmuch)
	}
	phase := false
	symone := "#"
	symtwo := " "
	if gm {
		symone = "# "
		symtwo = "  "
	}
	for i := 0; i < howmuch; i++ {
		horizontal(howmuch, phase, symone, symtwo)
		fullchessboard += "\n"
		phase = !phase
	}
	fmt.Println(fullchessboard)
	return fullchessboard, nil
}

func horizontal(howmuch int, f bool, symone string, symtwo string) {
	i := 1
	for i <= howmuch {
		if f {
			f = false
			fullchessboard += symtwo
		} else {
			fullchessboard += symone
			f = true
		}
		i++
	}
}
