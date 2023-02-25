// Struktury w Golang
package main

import (
	"encoding/json"
	"fmt"
)

type rectangle struct {
	a     float32
	b     float32
	color string
}

type Employee struct {
	Firstname string `json:firstname`
	Lastname  string `json:lastname`
	Location  string `json:location`
}

func main() {
	var rect1 = rectangle{10, 5, "red"}
	fmt.Println(rect1)

	var rect2 = new(rectangle)
	rect2.a = 5
	rect2.b = 10
	rect2.color = "green"
	fmt.Println(rect2)

	var rect3 = rectangle{a: 3, b: 12}
	fmt.Println(rect3)

	jsonString := `
	{
		"firstname" : "Jan",
		"lastname" : "Kowalski",
		"location" : "Warszawa"
	}
	`
	emp := new(Employee)
	json.Unmarshal([]byte(jsonString), emp)
	fmt.Println(emp)

	emp2 := new(Employee)
	emp2.Firstname, emp2.Lastname, emp2.Location = "Zenon", "M", "Białystok"
	jsonStr, _ := json.Marshal(emp2)
	fmt.Printf("%s\n", jsonStr)
}
