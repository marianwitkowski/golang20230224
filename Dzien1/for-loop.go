package main

import "fmt"

func main() {

	count := 10
	for i := 0; i < count; i++ {
		fmt.Println(i)
	}

	fmt.Println("=================")
	k := 1
	for ; k <= 10; k++ {
		fmt.Println(k)
	}

	fmt.Println("=================")
	k = 1
	for k <= 10 {
		fmt.Println(k)
		k++
	}

	fmt.Println("=================")
	for k = 1; ; k++ {
		fmt.Println(k)
		if k == 10 {
			break
		}
	}

	fmt.Println("=================")
	k = 1
	for {
		fmt.Println(k)
		if k++; k == 11 {
			break
		}
	}

	fmt.Println("=================")
	// iterowanie po słowniku
	countries := map[string]string{"PL": "Polska", "DE": "Niemcy", "US": "USA"}
	for abbrev, country := range countries {
		fmt.Printf("klucz=%v \t wartosc=%v \n", abbrev, country)
	}

	for _, country := range countries {
		fmt.Printf("wartosc=%v \n", country)
	}

}
