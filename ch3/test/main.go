package main

import (
	"strings"
	"unicode/utf8"
)

func main() {
	f := func(c rune) bool {
		return c == '+' || c == '-'
	}
	r := strings.Split("5+6", utf8.DecodeRune(f))

}
