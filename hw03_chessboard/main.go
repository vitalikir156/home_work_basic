package main

import "fmt"

func main() {
	phase := false
	howmuch := 8
	var gm string
	symone := "#"
	symtwo := " "
	fmt.Printf("Enter length: ")
	fmt.Scanln(&howmuch)
	fmt.Printf("Use doublespaces? ")
	fmt.Scanln(&gm)
	if gm == "y" || gm == "Y" || gm == "yes" || gm == "true" {
		symone = "# "
		symtwo = "  "
	}
	i := 1
	for i <= howmuch {
		horizontal(howmuch, phase, symone, symtwo)
		fmt.Println()
		if phase {
			phase = false
		} else {
			phase = true
		}
		i++
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
