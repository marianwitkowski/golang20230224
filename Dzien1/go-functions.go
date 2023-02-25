// stosowanie funkcji w Golang

package main

import (
	"fmt"
	"reflect"
)

func dummyFunc() {
	fmt.Println("XYZ")
}

// deklaracja funkcji przyjmujacej parametry i zwracającej wartośc
func add(x int, y int) int {
	var result int = x + y
	return result
}

func cube(x int) (result int) {
	result = x * x * x
	return
}

func rect(a int, b int) (area int, circuit int) {
	area = a * b
	circuit = 2*a + 2*b
	return
}

// przekazywanie parametrów przez wskaźnik
func updateByPointer(a *int, s *string) {
	*a += 20
	*s += "...chyba Ty"
	return
}

var (
	area = func(a int, b int) int {
		return a * b
	}
)

func anyParams(s ...string) {
	fmt.Println(s)
}

func variadicFunction(s string, numbers ...int) {
	fmt.Println(s)
	fmt.Println(numbers)
}

func anyVariadicFunction(x ...interface{}) {
	for _, v := range x {
		fmt.Println(" --> ", v, reflect.TypeOf(v))
	}
}

// domknięcia w Golang
func intGenerator() func() int {
	counter := 0
	return func() int {
		counter++
		return counter
	}
}

// ===============================================
func main() {
	sum := add(20, 30)
	fmt.Println(sum)
	fmt.Println(cube(5))
	var a, c int
	a, c = rect(2, 3)
	fmt.Printf("Pole=%v Obwod=%v \n", a, c)

	var x = 20
	var s = "Chce mi się pic"
	fmt.Println("Przed:", x, s)
	updateByPointer(&x, &s)
	fmt.Println("Po:", x, s)

	fmt.Println("Wywołanie f. anonimowej ", area(4, 5))

	anyParams("Hello")
	anyParams("Hello", "world", "!")

	variadicFunction("Hello", 1, 2, 4)
	variadicFunction("Hello", 1, 2, 4, 10, 40, 50)

	anyVariadicFunction(1, "hello", true, 10.5, []string{"A", "B", "C"})

	genCounter := intGenerator()
	fmt.Println("================")
	fmt.Println(genCounter())
	fmt.Println(genCounter())
	fmt.Println(genCounter())
}
