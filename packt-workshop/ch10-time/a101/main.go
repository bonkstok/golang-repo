package main

import (
	"fmt"
	"time"
)

func main() {
	timeNow := time.Now()
	second := timeNow.Second()
	minute := timeNow.Minute()
	hour := timeNow.Hour()
	day := timeNow.Day()
	month := timeNow.Month()
	year := timeNow.Year()

	fmt.Printf("%v:%v:%v %v/%v/%v\n", hour, minute, second, year, month, day)
}
