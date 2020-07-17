package main

import (
	"fmt"
	"regexp"
)

func main() {
	input := "I am Leo"
	println("input: ", input)
	// case insensitive with (?i)
	re := regexp.MustCompile("(?i)(i[' a]+m) (.*)")
	if re.MatchString(input) {
		match := re.ReplaceAllString(input, "How do you, $2!")
		fmt.Println("output: " + match)
	} else {
		fmt.Println("no match")
	}
}
