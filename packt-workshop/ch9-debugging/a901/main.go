package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

var (
	ErrInvalidSSNLength  = errors.New("invalid SSN length")
	ErrInvalidSSNNumbers = errors.New("invalid SSN numbers")
	ErrInvalidSSNPrefix  = errors.New("SSN has 000 as prefix")
	ErrInvalidDigitPlace = errors.New("SSNs that start wwith 9 require 7 or 9 in fourth place")
)

func checkSSNLength(input string) error {
	if len(input) != 9 {
		return fmt.Errorf("value of %s caused error: %v", input, ErrInvalidSSNLength)
	}
	return nil
}

func checkChars(input string) error {
	for _, k := range input {
		if string(k) == "-" {
			continue
		}
		if !isInt(string(k)) {
			return fmt.Errorf("%s - expected type int got %T", ErrInvalidSSNNumbers, string(k))
		}
	}
	return nil
}

func checkPrefix(input string) error {
	if strings.HasPrefix(input, "000") {
		return ErrInvalidSSNPrefix
	}
	return nil
}

func checkFirstNumber(input string) error {
	if string(input[0]) == "9" && (string(input[3]) != "7" && string(input[3]) != "9") {
		return ErrInvalidDigitPlace
	}
	return nil
}

func isInt(input string) bool {
	_, err := strconv.ParseInt(input, 0, 64)
	return err == nil
}

func main() {
	validateSSN := []string{"123-45-6789", "012-8-678", "000-12-0962", "999-33-3333", "087-65-4321", "123-45-zzzz", "000-00000"}

	for _, v := range validateSSN {

		err := checkChars(v)
		if err != nil {
			fmt.Println(err)
			continue
		}
		err = checkSSNLength(v)
		if err != nil {
			fmt.Println(err)
			continue
		}
		err = checkPrefix(v)
		if err != nil {
			fmt.Println(err)
			continue
		}
		err = checkFirstNumber(v)
		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Println(v)
	}

}
