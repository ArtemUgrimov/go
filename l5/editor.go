package main

import (
	"fmt"
	"os"
	"strings"
)

type TextEditor struct {
	text  []string
	index map[string][]int
}

func (editor *TextEditor) createIndex() map[string][]int {
	editor.index = make(map[string][]int)
	for index, phrase := range editor.text {
		words := strings.Split(phrase, " ")
		for _, word := range words {
			// I am removing a dot at the end of the sentence
			// since it doesnt trim while doing the split
			trimmedWord := strings.ToLower(strings.Trim(word, ".,"))
			rowsWithWord, ok := editor.index[trimmedWord]
			if !ok {
				editor.index[trimmedWord] = []int{index}
			} else {
				if rowsWithWord[len(rowsWithWord)-1] != index {
					rowsWithWord = append(rowsWithWord, index)
					editor.index[trimmedWord] = rowsWithWord
				}
			}
		}
	}
	return editor.index
}

func (editor *TextEditor) initWithString(text string) {
	editor.text = strings.Split(text, "\n")
	editor.createIndex()
}

func (editor *TextEditor) initWithFile(path string) {
	dat, err := os.ReadFile(path)
	if err != nil {
		fmt.Printf("TextEditor cannot open a file with path %s. Error : %s", path, err)
		os.Exit(1)
	}
	editor.text = strings.Split(string(dat), "\n")
	editor.createIndex()
}

func (editor *TextEditor) findWord(word string) []int {
	rows, ok := editor.index[strings.ToLower(word)]
	if ok {
		return rows
	} else {
		return []int{}
	}
}

func (editor *TextEditor) findWordSlow(word string) []int {
	wordToFind := strings.ToLower(word)
	// make an array with the size of 0 and capacity of max available rows count
	result := make([]int, 0, len(editor.text))
	// iterating through the rows of text
	for row, str := range editor.text {
		// iterating through the words of the row
		for _, word := range strings.Split(str, " ") {
			if strings.ToLower(strings.Trim(word, ".,")) == wordToFind {
				// check that we don't have such row in the result
				if len(result) == 0 || result[len(result)-1] != row {
					result = append(result, row)
				}
			}
		}
	}
	return result
}

func (editor *TextEditor) countWords() {
	for word, counter := range editor.index {
		fmt.Println(word, "occures in the next lines", counter)
	}
}
