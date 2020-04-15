package main

import (
	"fmt"
	"time"
)

func getConfig() (bool, string, time.Time) {
	return false, "info", time.Now()
}

func main() {
	//one line declaration of variables
	//usually not used due to readability for declaring vars, it is used when an function returns multiple values
	//Debug, LogLevel, startUpTime := false, "info", time.Now()

	//function getConfig returns 3 values
	Debug, LogLevel, startUpTime := getConfig()
	fmt.Println(Debug, LogLevel, startUpTime)
}
