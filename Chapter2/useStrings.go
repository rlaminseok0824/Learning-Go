package main

import (
	"fmt"
	s "strings"
)

var f = fmt.Printf

func main() {
	upper := s.ToUpper("Hello there!")
	f("To upper: %s\n", upper)
	f("To Lower: %s\n", s.ToLower(upper))
}
