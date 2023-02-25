// Tablice w Golang
package main

import "fmt"

func main() {

	var intArray [3]int
	intArray[0] = 1
	intArray[1] = 10
	intArray[2] = 100

	a := [5]int{10, 20, 30, 40, 50}
	var b [5]int = [5]int{10, 20, 30, 40, 50}
	c := [...]int{10, 20, 30, 40, 50}
	d := [10]int{0: 1, 2: 1} // zmiana dla indeksu 0 i 2 na wartosc 1
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)

	for i := 0; i < len(c); i++ {
		fmt.Print(c[i], "|")
	}
	fmt.Println()

	for index, value := range c {
		fmt.Println(index, " -> ", value)
	}

	copyA := a
	fmt.Println(copyA)
	copyA[0] = -10
	fmt.Println(a, "|", copyA)

	copyAPointer := &a // przekazanie adresu/referencji do tablicy a
	copyAPointer[0] = -100
	fmt.Println(a, "|", copyAPointer)

	// wybieranie na raz większej liczby elementów z tablicy
	// 10, 20, 30, 40, 50
	fmt.Println("2 pierwsze elementy B: ", b[:2])
	fmt.Println("2-gi, 3 i 4-ty element B: ", b[1:4])
	fmt.Println("Od wart. 30 do konca ", b[2:])
	fmt.Println("Dwa ostanie elementy", b[len(b)-2:])

}
