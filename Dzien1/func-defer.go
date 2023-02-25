// opoznianie wykonania kodu

package main

import (
	"fmt"
	"os"
)

func createFile(fn string) *os.File {
	fmt.Println("Otwieram...")
	fd, err := os.Create(fn)
	if err != nil {
		panic(err)
	}
	return fd
}

func writeFile(fd *os.File, text string) {
	fmt.Println("Zapisuję...")
	fmt.Fprintln(fd, text)
}

func closeFile(fd *os.File) {
	fmt.Println("Zamykam...")
	err := fd.Close()
	if err != nil {
		fmt.Fprintf(os.Stderr, "wystapil błąd : %v\n", err)
		os.Exit(1)
	}
}

func main() {
	/*defer fmt.Println("Linia 1")
	defer fmt.Println("Linia 2")
	fmt.Println("Linia 3")*/

	fd := createFile("test.txt")
	defer closeFile(fd)
	writeFile(fd, "Linia1")

}
