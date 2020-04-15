package main

import (
	"fmt"
	"runtime"
)

func main() {
	var list []int // 403 MiB - each is 64 bit / 8 byte -- 64bit  * 10 million - will not be 640miB
	//var list []int8 // 53MiB - each is 8 bit / 1 byte

	for i := 0; i < 10000000; i++ {
		list = append(list, 9223372036854775807)
	}
	//to set memory to 1 Gigibyte GiB - you need 1024 MiB
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("TotalAlloc (Heap) = %v MiB\n", m.TotalAlloc/1024/1024) //mebibytes
	fmt.Printf("TotalAlloc (Heap) = %v \n", m.TotalAlloc)              //prints in bytes
	fmt.Println(len(list))
}
