package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	date := time.Date(2019, 31, 31, 02, 49, 21, 0, time.UTC)
	day := strconv.Itoa(date.Day())
	month := strconv.Itoa(int(date.Month()))
	year := strconv.Itoa(date.Year())
	second := strconv.Itoa(date.Second())
	minute := strconv.Itoa(date.Minute())
	hour := strconv.Itoa(date.Hour())
	fmt.Println(date.Format(time.ANSIC))
	fmt.Printf("Type of day %T\n", day)
	fmt.Printf("type: %T\n", time.January)

	fmt.Printf("%v:%v:%v %v/%v/%v\n", hour, minute, second, year, month, day)
}
