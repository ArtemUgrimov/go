package main

import (
	"fmt"

	"github.com/ArtemUgrimov/findutils"
)

func main() {
	result, error := findutils.FindNumbersInFile("numbers.txt")
	if error != nil {
		fmt.Println(error)
	} else {
		fmt.Println(result)
	}
}
