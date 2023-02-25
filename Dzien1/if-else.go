// instrukcja warunkowa
package main

import (
	"fmt"
	"strconv"
)

func main() {

	var age = 20
	if age >= 18 {
		fmt.Println("Jesteś pełnoletni")
	} else {
		fmt.Println("Nie jesteś pełnoletni")
	}

	var x = 16
	if (x%5 == 0) && (x%3 == 0) {
		fmt.Println("Liczba podzielna przez 5 i 3 jednocześnie")
	} else if x%5 == 0 {
		fmt.Println("Liczba podzielna przez 5")
	} else if x%3 == 0 {
		fmt.Println("Liczba podzielna przez 3")
	} else {
		fmt.Println("Ani to, ani tamto")
	}

	if x = 30; x == 30 {
		fmt.Println("Liczba ma wartośc 30")
	}

	v := "ABC"
	if s, err := strconv.Atoi(v); err == nil {
		fmt.Printf("%T, %v", s, s)
	} else {
		fmt.Println("Konwersja nie powiodła się")
	}

}
