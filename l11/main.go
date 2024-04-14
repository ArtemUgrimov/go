package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
)

func task1() {
	file, error := os.Open("1689007675141_numbers.txt")
	if error != nil {
		fmt.Print(error)
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	telNumbers := make([]string, 0)
	for {
		lineBytes, _, err := reader.ReadLine()
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Print(err)
			return
		}
		line := string(lineBytes)
		telNumbers = append(telNumbers, line)
	}

	// may be open bracket(or not) then 3 digits then closed bracket, space or dot
	// then may be space and three digits
	// then may be dash, dot or space(or not)
	// then should be 4 digits
	telNumbersRegEx := regexp.MustCompile(`\(?\d{3}[\)-. ]? ?\d{3}[-. ]?\d{4}`)
	for _, tn := range telNumbers {
		found := telNumbersRegEx.FindAllStringIndex(tn, -1)
		fmt.Println(found)
	}
}

func task2() {
	file, error := os.Open("1689007676028_text.txt")
	if error != nil {
		fmt.Print(error)
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	contentBytes := make([]byte, 2048)
	_, error = reader.Read(contentBytes)
	if error != nil && error != io.EOF {
		fmt.Println(error)
		return
	}
	content := string(contentBytes)

	findPatterns := []*regexp.Regexp{
		regexp.MustCompile(`[йцкнгшщзхфвпрлджчмтбЙЦКНГШЩЗХФВПРЛДЖЧМТБ]\S*[йцкнгшщзхфвпрлджчмтб][^ ,.]*`),
		regexp.MustCompile(``),
		regexp.MustCompile(``),
	}
	for _, pattern := range findPatterns {
		result := pattern.FindAllStringIndex(content, -1)
		fmt.Println(result)
	}
}

func main() {
	task1()
	task2()
}
