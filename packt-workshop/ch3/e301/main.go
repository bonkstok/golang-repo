package main

import (
	"fmt"
	"unicode"
)

func main() {
	fmt.Println(passwordChecker("asdf2asd@!Dfas"))
}

func passwordChecker(pw string) bool {
	pwR := []rune(pw) // convert password string itno rune -> type that is safe for multi-byte utf8 chars
	fmt.Println(pwR)

	if len(pwR) < 8 {
		return false
	}
	hasUpper, hasLower, hasNumber, hasSymbol := false, false, false, false

	for _, v := range pwR {

		if unicode.IsUpper(v) {
			hasUpper = true
		}
		if unicode.IsNumber(v) {
			hasNumber = true
		}
		if unicode.IsLower(v) {
			hasLower = true
		}
		if unicode.IsSymbol(v) || unicode.IsPunct(v) {
			hasSymbol = true
		}
	}
	return hasUpper && hasLower && hasNumber && hasSymbol //if all is true, return true else false
}
