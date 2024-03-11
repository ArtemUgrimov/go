package main

import (
	"fmt"
	"sort"
)

type SuperSus struct {
	ID int
}

func isPresent(input []SuperSus, element SuperSus) bool {
	for _, item := range input {
		if item.ID == element.ID {
			return true
		}
	}
	return false
}

func uniqSorted(input []SuperSus) []SuperSus {
	var result []SuperSus
	for _, item := range input {
		if !isPresent(result, item) {
			result = append(result, item)
		}
	}
	sort.Slice(result, func(i, j int) bool {
		return result[i].ID < result[j].ID
	})

	return result
}

func main() {
	var elements = []SuperSus{
		{1}, {3}, {5}, {10}, {1}, {5}, {100}, {-1},
	}
	result := uniqSorted(elements)
	fmt.Println(result)
}
