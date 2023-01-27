package main

import (
	"fmt"
	"unicode"
	"unicode/utf8"
)

type validator func(s string) bool

func anyCharCheck(s string, charCheck func(ch rune) bool) bool {
	for _, char := range s {
		if charCheck(char) {
			return true
		}
	}
	return false
}

func digits(s string) bool {
	return anyCharCheck(s, unicode.IsDigit)
}

func letters(s string) bool {
	return anyCharCheck(s, unicode.IsLetter)
}

func minlen(length int) validator {
	return func(s string) bool {
		if utf8.RuneCountInString(s) >= length {
			return true
		}
		return false
	}
}

func and(funcs ...validator) validator {
	return func(s string) bool {
		for _, fn := range funcs {
			if !fn(s) {
				return false
			}
		}
		return true
	}
}

func or(funcs ...validator) validator {
	return func(s string) bool {
		for _, fn := range funcs {
			if fn(s) {
				return true
			}
		}
		return false
	}
}

type password struct {
	value string
	validator
}

func (p *password) isValid() bool {
	return p.validator(p.value)
}

func main() {
	var s string
	fmt.Scan(&s)
	validator := or(and(digits, letters), minlen(10))
	p := password{s, validator}
	fmt.Println(p.isValid())
}
