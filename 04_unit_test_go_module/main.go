// create go module:
//	go mod init demo.org/hello
// run:
//	go run main.go
// test:
//	go test

package main

import (
	"fmt"
)

func main() {
	fmt.Println(Hello())
}

// Hello function to be unit-tested
func Hello() string {
	return "Hello, world."
}
