package main

import (
	"fmt"
	"time"
)

func main() {
	dayBorn := time.Sunday
	//fmt.Println(dayBorn)
	switch dayBorn {
	case time.Monday, time.Tuesday, time.Wednesday, time.Thursday, time.Friday:
		fmt.Println("Weekday child")
	case time.Saturday, time.Sunday:
		fmt.Println("Weekend child")
	default:
		fmt.Println("Day is not known! error")
	}

}
