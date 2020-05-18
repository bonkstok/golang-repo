package main

import (
	"encoding/json"
	"fmt"
)

type student struct {
	StudentID     int      `json:"id"`
	LastName      string   `json:"lname"`
	MiddleInitial string   `json:"minitial"`
	FirstName     string   `json:"fname"`
	IsEnrolled    bool     `json:"enrolled"`
	Courses       []course `json:"classes"`
}

type course struct {
	Name   string `json:"coursename"`
	Number int    `json:"coursenum"`
	Hours  int    `json:"coursehours"`
}

type simple struct {
	Name     string
	Password string
	Age      int
}

func main() {
	data := []byte(` 
	{ 
	  "id": 123, 
	  "lname": "Smith", 
	  "minitial": null, 
	  "fname": "John", 
	  "enrolled": true, 
	  "classes": [{ 
		"coursename": "Intro to Golang", 
		"coursenum": 101, 
		"coursehours": 4 
	  }, 
	  { 
		"coursename": "English Lit", 
		"coursenum": 101, 
		"coursehours": 3
	  }, 
	  { 
		"coursename": "World History", 
		"coursenum": 101, 
		"coursehours": 3 
	  } 
	] 
	 } 
  `)
	jsonValid := json.Valid(data)
	fmt.Println("Is json data valid:", jsonValid)
	var s student
	err := json.Unmarshal(data, &s)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(s)
	fmt.Println(s.FirstName)
	for _, v := range s.Courses {
		fmt.Println(v.Name)
	}
	fmt.Printf("%#v", s)
	dataSimple := []byte(`
	{
		"name": "Donaly",
		"age": 92,
		"password": "katrine"
	}
	`)
	fmt.Println("Is valid:", json.Valid(dataSimple))
	var s1 simple
	err = json.Unmarshal(dataSimple, &s1)
	if err != nil {
		fmt.Println("errror", err)
	}
	fmt.Println(s1)
}
