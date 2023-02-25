package main

import (
	"fmt"
	"time"
)

func main() {

	now := time.Now()
	fmt.Println(now)
	dayOfMonth := 20 //now.Day()
	switch dayOfMonth {
	case 10:
		fmt.Println("Zapłac ZUS")
	case 20:
		fmt.Println("Zapłac PIT")
		fmt.Println("Idę na koszykówkę")
	case 25:
		fmt.Println("Zapłac VAT")
	case 7, 14:
		fmt.Println("Idę na koszykówkę")
	default:
		fmt.Println("Dzis nie ma nic do zapłaty")
	}

}
