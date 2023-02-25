// maps - słowniki
package main

import (
	"fmt"
	"sort"
)

func findKey(key string, dict map[string]int) bool {
	for k := range dict {
		if k == key {
			return true
		}
	}
	return false
}

func main() {

	var regPlates = map[string]string{"WI": "Warszawa", "GD": "Gdańsk", "PO": "Poznań"}
	emptyMaps := map[string]int{}
	a := make(map[string]int)
	a["key1"] = 10
	a["key2"] = 20

	fmt.Println(regPlates)
	fmt.Println(emptyMaps)
	fmt.Println(a, len(a))
	if findKey("key3", a) {
		fmt.Println(a["key3"])
	} else {
		fmt.Println("Brak klucza")
	}

	// sortowanie słownika na podstawie kluczy
	keys := make([]string, 0, len(regPlates))
	for k := range regPlates {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	fmt.Println(keys)
	for _, k := range keys {
		fmt.Println(k, regPlates[k])
	}

}
