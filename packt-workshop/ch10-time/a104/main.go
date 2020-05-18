package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	fmt.Println(now.Format(time.ANSIC))
	SSS := time.Duration(21966 * time.Second)
	Future := now.Add(SSS)
	fmt.Println(Future.Format(time.ANSIC))
	//time.Duration()
}
