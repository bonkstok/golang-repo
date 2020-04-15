package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	r := rand.Intn(5) + 1
	//var rr int = 999
	//fmt.Println(r," ",rr)
	stars := strings.Repeat("*", r)

	fmt.Println(stars)

}
