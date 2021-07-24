package main

import (
	"github.com/rover10/learn/research"
)

func main() {
	research.Hello()
}

// https://golang.org/doc/tutorial/call-module-code
// go mod init github.com/rover10/learn
// go mod edit -replace github.com/rover10/learn/=../
// go mod tidy
// go install
// go run .
