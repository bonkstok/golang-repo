package main

import "fmt"

func itemsSold() {
	items := map[string]int{}
	items["John"] = 41
	items["Celina"] = 109
	items["Micah"] = 24
	items["Johnny"] = 1

	for k, v := range items {
		fmt.Printf("%s sold %d %s and %s\n", k, v, singularOrPlurar(v), result(v))

	}
}

func singularOrPlurar(input int) string {
	if input > 1 {
		return "items"
	} else {
		return "item"
	}
}

func result(input int) string {
	if input < 40 {
		return "is below expectations."
	} else if input > 40 && input <= 100 {
		return "meets expectations."
	} else {
		return "exceeded expectations."
	}
}

func main() {
	itemsSold()
}
