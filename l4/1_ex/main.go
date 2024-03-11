package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Editor struct {
	rows []string
}

func (editor *Editor) find(sub string) []string {
	var result []string
	for _, row := range editor.rows {
		if strings.Contains(row, sub) {
			result = append(result, row)
		}
	}
	return result
}

func (editor *Editor) fromConsole() {
	fmt.Println("Please, enter some strings one by one, divided by enter : (type 'end' to end the input)")
	editor.rows = []string{}
	reader := bufio.NewReader(os.Stdin)
	for {
		statement, err := reader.ReadString('\n')
		if err != nil {
			break
		}
		statement = strings.Trim(statement, "\r\n")
		if statement == "end" {
			break
		}
		editor.rows = append(editor.rows, statement)
	}
}

func (editor *Editor) fromFile(filename string) {
	editor.rows = []string{}
	file, err := os.Open(filename)

	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		editor.rows = append(editor.rows, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}
}

func main() {
	var editor = Editor{}
	editor.fromConsole()
	result := editor.find("hell")
	for _, item := range result {
		fmt.Println(item)
	}

	editor.fromFile("sample.txt")
	result = editor.find("hell")
	for _, item := range result {
		fmt.Println(item)
	}
}
