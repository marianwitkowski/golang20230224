// Konwersja typów danych w Golang

package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func main() {
	x := "100"
	// konwersja str to int
	x_int, err := strconv.Atoi(x)
	fmt.Println(x_int, err, reflect.TypeOf(x_int))

	y := "3.14159"
	y_flot, err := strconv.ParseFloat(y, 64)
	fmt.Println(y_flot, err, reflect.TypeOf(y_flot))

	// konwersja wewnątrz jednego typu danych (int)
	var i int = 123
	fmt.Println(reflect.TypeOf(i))

	var j int64
	j = int64(i)
	fmt.Println(j, reflect.TypeOf(j))

	var k int16
	k = int16(i)
	fmt.Println(k, reflect.TypeOf(k))

	var a float32 = 123.456
	var b float64
	b = float64(a)
	fmt.Println(b, reflect.TypeOf(b))

	fmt.Printf("wartosc a=%v, wartosc b=%v, type=%v\n", a, b, reflect.TypeOf(b))

	num1 := 123
	strNum1 := fmt.Sprintf("|%10d|", num1)
	strNum2 := fmt.Sprintf("|%-10d|", num1)
	strNum3 := fmt.Sprintf("|%b|", num1)

	fmt.Println(strNum1)
	fmt.Println(strNum2)
	fmt.Println(strNum3)
}
