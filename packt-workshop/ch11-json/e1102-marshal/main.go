package main

import (
	"encoding/json"
	"fmt"
)

type student struct {
	StudentId     int      `json:"id"`
	LastName      string   `json:"lname"`
	MiddleInitial string   `json:"minitial,omitempty"`
	FirstName     string   `json:"fname"`
	IsEnrolled    bool     `json:"enrolled"`
	Courses       []course `json:"classes"`
}

type course struct {
	Name   string `json:"coursename"`
	Number int    `json:"coursenum"`
	Hours  int    `json:"coursehours"`
}

func main() {

	s1 := student{
		StudentId:     3141,
		LastName:      "Bond",
		MiddleInitial: "",
		FirstName:     "Pietje",
		IsEnrolled:    true,
		Courses: []course{
			{
				Name: "English",
			}, {
				Name: "Dutch",
			},
		},
	}

	s2 := student{
		StudentId:     3155,
		LastName:      "Evil",
		MiddleInitial: "Von",
		FirstName:     "Gunther",
	}
	s2course := course{Name: "English", Number: 4123, Hours: 22}
	s2.Courses = append(s2.Courses, s2course)
	s2course = course{Name: "Dutch", Number: 4923, Hours: 22}
	s2.Courses = append(s2.Courses, s2course)

	st, err := json.Marshal(s1)
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Println(string(st))
	fmt.Printf("%s\n", st)

	st, err = json.Marshal(s2)
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Printf("%s\n", st)

	st, err = json.MarshalIndent(s2, "", " ")
	if err != nil {
		fmt.Println("Error:", err)
	}
	//fmt.Printf("%s\n", st)
}
