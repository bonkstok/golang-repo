package main

import (
	"fmt"
)

func main() {

	var1, var2, var3 := "sheep", 100, 500
	var1, var2, var3 = "sheep", var3, 50
	fmt.Println(var1, var2, var3)
}
