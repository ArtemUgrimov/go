package main

import (
	"fmt"
	"time"
)

type InitFunc func(string)
type FindFunc func(string) []int

func measure(initFunc InitFunc, initArg string, findFastFunc FindFunc, findSlowFunc FindFunc) {
	startHashAllTime := time.Now()
	initFunc(initArg)
	startHashFindTime := time.Now()
	_ = findFastFunc("hash")
	elapsedHashAllTime := time.Now()

	// fmt.Println("Word `hash` occures in the next lines :")
	// fmt.Println(linesWithHash)

	// fmt.Println("==========================================")

	startLinearFindTime := time.Now()
	_ = findSlowFunc("hash")
	endLinearFindTime := time.Now()
	// fmt.Println("Word `hash` occures in the next lines :")
	// fmt.Println(linesWithHash)

	// fmt.Println("==========================================")

	fmt.Printf("Hashing full cycle spent %s but the search alone spent %s\n",
		elapsedHashAllTime.Sub(startHashAllTime), elapsedHashAllTime.Sub(startHashFindTime))

	fmt.Printf("Linear search spent %s\n", endLinearFindTime.Sub(startLinearFindTime))

}

func main() {
	text :=
		`Hash map data structures use a hash function, which turns a key into an index within an underlying array.
The hash function can be used to access an index when inserting a value or retrieving a value from a hash map.
Hash map underlying data structure
Hash maps are built on top of an underlying array data structure using an indexing system.
Each index in the array can store one key-value pair.
If the hash map is implemented using chaining for collision resolution, each index can store another data structure such as a linked list, which stores all values for multiple keys that hash to the same index.
Each Hash Map key can be paired with only one value. However, different keys can be paired with the same value.
Also I would like to add a few lines to this text to have exactly 10 lines of text.
By the way, did you know about creator of go, our god...
Oh this is a very long story. Anyway, goodbye.`

	textEditor := TextEditor{}

	fmt.Println("Text as a variable (10 lines)========")
	measure(textEditor.initWithString, text, textEditor.findWord, textEditor.findWordSlow)
	fmt.Println("==========================================")
	fmt.Println("Text as a 100 lines from file =======")
	measure(textEditor.initWithFile, "100.txt", textEditor.findWord, textEditor.findWordSlow)
	fmt.Println("==========================================")
	fmt.Println("Text as a 1000 lines from file ======")
	measure(textEditor.initWithFile, "1000.txt", textEditor.findWord, textEditor.findWordSlow)
	fmt.Println("==========================================")
	fmt.Println("Text as a 10k lines from file =======")
	measure(textEditor.initWithFile, "10k.txt", textEditor.findWord, textEditor.findWordSlow)
	fmt.Println("==========================================")
	// textEditor.countWords()
}
