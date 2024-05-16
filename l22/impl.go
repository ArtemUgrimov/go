package main

import (
	"fmt"
	"reflect"
	"slices"
	"strings"
)

func Serialize(v interface{}) string {
	var resultItems []string

	tp := reflect.TypeOf(v)
	val := reflect.ValueOf(v)

	for i := 0; i < tp.NumField(); i++ {
		field := tp.Field(i)
		tagValue := field.Tag.Get("param")
		fieldValue := val.Field(i)

		if tagValue == "-" {
			continue
		}

		resultItems = append(resultItems, fmt.Sprintf("%s=%v", tagValue, fieldValue))
	}
	slices.Sort(resultItems)

	return strings.Join(resultItems, " ")
}
