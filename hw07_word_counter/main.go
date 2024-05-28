package main

import (
	"fmt"
	"regexp"
	"strings"
)

func main() {
	fmt.Println(CountWords("Раз раз    ,;        -089u0 ad [0o9t\rshb86574]   раз два три два ТрИ \nТри рАрз раз"))
}

func CountWords(instring string) map[string]int {
	var instringnew string
	instring = strings.ToLower(instring)
	instring = strings.ReplaceAll(instring, "\n", " ")
	instring = strings.ReplaceAll(instring, "—", " ")
	regexp := regexp.MustCompile(`[[:punct:]]|[[:cntrl:]]|[[:digit:]]`)
	instring = regexp.ReplaceAllString(instring, "")
	for instring != instringnew { // этот быдлокод призван убрать пробелы
		instringnew = instring
		instring = strings.ReplaceAll(instring, "  ", " ")
	}
	stringmap := map[string]int{}
	splitstring := strings.Split(instringnew, " ")
	for _, word := range splitstring {
		val, ok := stringmap[word]
		if ok {
			stringmap[word] = val + 1
		} else {
			stringmap[word] = 1
		}
	}
	return stringmap
}
