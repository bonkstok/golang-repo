package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()

	sum := 0

	for i := 0; i < 255; i++ {
		sum += i
		fmt.Println(sum)
	}
	end := time.Now()
	duration := end.Sub(start)

	fmt.Println(duration) //displayes in nanoseconds

	fmt.Println("Task took", duration.Hours(), "hours to complete")
	fmt.Println("Task took", duration.Minutes(), "hours to complete")
	fmt.Println("Task took", duration.Seconds(), "seconds to complete")
	fmt.Println("Task took", duration.Nanoseconds(), "ns to complete")

	fmt.Println("The script completed at: ", end)
	fmt.Println("The task took", duration.Hours(), "hour(s) to complete!")
	fmt.Println("The task took", duration.Minutes(), "minutes(s) to complete!")
	fmt.Println("The task took", duration.Seconds(), "seconds(s) to complete!")
	fmt.Println("The task took", duration.Nanoseconds(), "nanosecond(s) to complete!")
}
