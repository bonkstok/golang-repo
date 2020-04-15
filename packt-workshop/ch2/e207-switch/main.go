package main

import (
	"fmt"
	"time"
)

func main() {
	//fmt.Println(dayBorn)
	switch dayBorn := time.Sunday; {
	case dayBorn == time.Sunday || dayBorn == time.Saturday:
		fmt.Println("WEEKEND")
	default:
		fmt.Println("NOT THE WEEKEND")
	}
}
