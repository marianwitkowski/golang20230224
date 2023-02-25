// operatory w Golang
package main

import "fmt"

// arytmetyczne
func main() {

	x, y := 12, 10
	fmt.Println(x + y)
	fmt.Println(x - y)
	fmt.Println(x * y)
	fmt.Println(x / y)
	a, b := 13, 3
	fmt.Println(a / b)
	fmt.Println(a % b)
	a++
	fmt.Println(a)
	a--
	fmt.Println(a)

	x = 100
	x += 2 // x = x + 2
	x -= 2 // x = x - 2
	x *= 3 // x = x * 3
	x /= 3 // x = x / 3
	x %= 3 // x = x % 3

	/*
	 operatory komparacji
	 == != < <= > >=
	*/
	z := x > y
	fmt.Println(z)

	// x > y i jednoczenie a <= 10
	z = (x > y) && a <= 10
	// x > y lub a <= 10
	z = (x > y) || a <= 10
	// negacja x>y
	z = !(x > y)

	// operatory bitow
	var m uint = 0b1010 // 111
	var n uint = 0b0110
	var p uint

	p = m & n
	fmt.Printf("%b\n", p)
	p = m | n
	fmt.Printf("%b\n", p)
	p = m ^ n
	fmt.Printf("%b\n", p)
	p = m << 1
	fmt.Printf("%b\n", p)
	p = m >> 1
	fmt.Printf("%b\n", p)
}
