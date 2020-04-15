package main

import "fmt"

type name string //custom string type

//create a struct location with 2 int fields

type location struct {
	x int
	y int
}

type size struct {
	width  int
	height int
}

//struct dot embeds each of the precending structs in it
type dot struct {
	name
	location
	size
}

func getDots() []dot {
	var dot1 dot    //all have zero values
	dot2 := dot{}   //all have zero values
	dot2.name = "A" //using the type name
	//for size and location use promoted fields to set their value
	dot2.x = 5
	dot2.y = 6
	dot2.width = 10
	dot2.height = 20

	//when initializing embedded tpyes you cant use promotion
	//for name, the result is the sane but for location and size smth different

	dot3 := dot{
		name: "B",
		location: location{
			x: 13,
			y: 27,
		},
		size: size{
			width:  5,
			height: 7,
		},
	}

	//use the type names to set data
	dot4 := dot{}
	dot4.name = "C"
	dot4.location.x = 101
	dot4.location.y = 209
	dot4.size.width = 87
	dot4.size.height = 43

	return []dot{dot1, dot2, dot3, dot4}
}

func main() {
	myDots := getDots()

	for index, element := range myDots {
		fmt.Println(element.name)
		fmt.Printf("dot%v: %#v\n", index+1, element)
	}

}
