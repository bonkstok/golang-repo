package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	time.Sleep(time.Second * 2)
	end := time.Now()
	elapsed := end.Sub(start)
	fmt.Println(elapsed.Seconds())
}
