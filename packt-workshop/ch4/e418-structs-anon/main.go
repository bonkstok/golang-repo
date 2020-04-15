package main

import "fmt"

type point struct {
	x int
	y int
}

func compare() (bool, bool) {
	//first anonymous struct
	point1 := struct {
		x int
		y int
	}{
		10,
		10,
	}

	point2 := struct {
		x int
		y int
	}{}
	point2.x = 10
	point2.y = 5

	//named struct
	point3 := point{10, 10}

	return point1 == point2, point1 == point3 //false, true
}

func main() {
	a, b := compare()
	fmt.Println(a)
	fmt.Println(b)
}
