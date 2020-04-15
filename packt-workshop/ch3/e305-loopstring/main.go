package main

import "fmt"

func main() {
	logLevel := "デバッグ"

	for index, runeVal := range logLevel {
		fmt.Println(index, string(runeVal))
	}
	fmt.Println(len(logLevel))
	for i := 0; i < len(logLevel); i++ {
		fmt.Println(logLevel[i])
	}
}
