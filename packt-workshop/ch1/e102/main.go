package main

import "fmt"

func main() {
	person_info := [4]string{"John", "Backname", "34", "me@ma.local"}
	for i := 0; i < len(person_info); i++ {
		fmt.Println(person_info[i])
	}

}
