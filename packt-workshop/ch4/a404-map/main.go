package main

import "fmt"

func main() {
	mySlice := []string{
		"Monday",
		"Tuesday",
		"Wednesday",
		"Thursday",
		"Friday",
		"Saturday",
		"Sunday",
	}
	//var tmpSlice [7]string
	//	var tmpSlice = make([]string, 7)

	//fmt.Println(mySlice[0 : len(mySlice)-1])
	//tmpSlice = mySlice[0 : len(mySlice)-1]
	//mySlice[0] = mySlice[len(mySlice)-1]
	//mySlice[4:] = append(mySlice, tmpSlice...)
	/*
		for i := 0; i < len(mySlice); i++ {
			//fmt.Println(mySlice[i])
			if i == (len(mySlice) - 1) {
				tmpSlice[0] = mySlice[i]
			} else {
				tmpSlice[i+1] = mySlice[i]
			}
		}
		mySlice = tmpSlice
		fmt.Println(mySlice)
	*/
	tmpSlice := []string{}

	tmpSlice = append(tmpSlice, mySlice[(len(mySlice)-1)])
	tmpSlice = append(tmpSlice, mySlice[0:(len(mySlice)-1)]...)
	fmt.Println(tmpSlice)

	week := []string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}
	week = append(week[6:], week[:6]...) //easier since append does not need a list as first argument.. guess you make a new slice with sunday and append rest to it

}
