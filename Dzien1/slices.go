// slices - coś na styl listy
package main

import "fmt"

func main() {
	a := make([]string, 3, 5)
	a[0], a[1], a[2] = "A", "B", "C"
	fmt.Println(a)

	b := new([50]int)[0:10]
	fmt.Printf("%v, len:%v cap:%v \n", b, len(b), cap(b))

	// dodanie wartosci do slices "a"
	a = append(a, "D", "E", "F")
	fmt.Println(a)

	// a klonuje do c
	c := make([]string, len(a))
	copy(c, a)
	fmt.Println(c)

	//iteracja
	for index, value := range c {
		fmt.Println(index, " --> ", value)
	}

}
