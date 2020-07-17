package main

import (
	"fmt"
	"regexp"
	"strings"
)

func main() {
	fmt.Println("convert camel case to snake case")
	s := "ThisIsASample"
	fmt.Println(s)

	// expected: "this_is_a_sample"
	r := convert(s)
	fmt.Printf("%q", r)
}

func convert(s string) string {
	var pHighLow = regexp.MustCompile(`(.)([A-Z][a-z0-9])`)
	var pLowHigh = regexp.MustCompile(`([a-z])([A-Z])`)

	snake := pHighLow.ReplaceAllString(s, "${1}_${2}")
	snake = pLowHigh.ReplaceAllString(snake, "${1}_${2}")

	return strings.ToLower(snake)

}
