package main

import "fmt"

type person struct {
	lname  string
	age    int
	salary float64
}

func main() {
	p1 := person{
		lname:  "van Achter",
		age:    88,
		salary: 2500.12,
	}

	fname := "Joe"
	grades := []int{100, 87, 67}
	states := map[string]string{"NH": "Noord-Holland", "ZH": "Zuid-Holland", "GER": "Limburg"}

	gatherAll := []interface{}{p1, fname, grades, states}
	fmt.Println("Formating to GO representation")
	for _, v := range gatherAll {
		fmt.Printf("%#v\n", v)
	}

}
