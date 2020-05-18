package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	nowNewYork, _ := time.LoadLocation("America/New_York")
	nowLA, _ := time.LoadLocation("America/Los_Angeles")

	fmt.Println("Current time:", now.Format(time.ANSIC))
	fmt.Println("Time in NY:", now.In(nowNewYork).Format(time.ANSIC))
	fmt.Println("Time in LA:", now.In(nowLA).Format(time.ANSIC))

	t := time.Date(2019, time.November, 3, 13, 0, 0, 0, time.UTC)
	chLoc, _ := time.LoadLocation("America/Chicago")
	fmt.Println(t.In(chLoc))
}
