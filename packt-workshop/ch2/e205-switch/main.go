package main

import (
	"fmt"
	"time"
)

func main() {
	dayBorn := time.Sunday
	//fmt.Println(dayBorn)
	switch dayBorn {
	case time.Monday:
		fmt.Println("Monday child")
	case time.Tuesday:
		fmt.Println("Tuesday child)")
	case time.Wednesday:
		fmt.Println("Wednesday child")
	case time.Thursday:
		fmt.Println("Thursday child")
	case time.Friday:
		fmt.Println("Friday child")
	case time.Saturday:
		fmt.Println("Saturday child")
	case time.Sunday:
		fmt.Println("Sunday child")
	default:
		fmt.Println("Day is not known! error")
	}

}
