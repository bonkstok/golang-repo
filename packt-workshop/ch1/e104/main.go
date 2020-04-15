package main

import (
	"fmt"
	"time"
)

var (
	Debug       bool // will be assigned bool's zero value - which is false
	LogLevel    = "info"
	startUpTime = time.Now()
)

func main() {
	fmt.Println(Debug)
}
