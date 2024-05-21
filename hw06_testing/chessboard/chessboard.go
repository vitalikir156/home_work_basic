package chessboard

import "fmt"

func Chessboard() {
	phase := false
	var howmuch int
	var gm string
	symone := "#"
	symtwo := " "
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
	} else if gm == "y" || gm == "Y" || gm == "yes" || gm == "true" {
		symone = "# "
		symtwo = "  "
	}

	for i := 0; i < howmuch; i++ {
		horizontal(howmuch, phase, symone, symtwo)
		fmt.Println()
		phase = !phase
	}
}

func horizontal(howmuch int, f bool, symone string, symtwo string) {
	i := 1
	for i <= howmuch {
		if f {
			f = false
			fmt.Printf("%s", symtwo)
		} else {
			fmt.Printf("%s", symone)
			f = true
		}
		i++
	}
}
