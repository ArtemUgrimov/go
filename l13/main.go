package main

import (
	"fmt"

	findutils "github.com/ArtemUgrimov/findutils/pkg"
)

func main() {
	fmt.Println(findutils.FindNumbersInFile("numbers.txt"))
}
