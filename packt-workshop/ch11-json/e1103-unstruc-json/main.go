package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	jsonData := []byte(`
	{
		"id": 0,
		"lname": "Utheren",
		"fname": "Kareltje",
		"isEnrolled": true,
		"grades": [100,86,33.96],
		"class":
		{
			"coursename": "History",
			"coursenum": 101,
			"coursehours": 3
		}
	}
`)
	if !json.Valid(jsonData) {
		fmt.Println("Json is not valid")
		os.Exit(1)
	}
	x := map[string]interface{}{}
	err := json.Unmarshal(jsonData, &x)
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Println(x)
	checkTypes(x)
	var v interface{}
	err = json.Unmarshal(jsonData, &v)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	fmt.Println(v)

}

func checkTypes(input map[string]interface{}) {
	for k, v := range input {
		switch value := v.(type) {
		case string:
			fmt.Printf("(string) - %v => %v\n", k, value)
		case int:
			fmt.Printf("(integer) - %v => %v\n", k, value)
		case bool:
			fmt.Printf("(bool) - %v => %v\n", k, value)
		case []interface{}:
			for i, j := range value {
				fmt.Println(i, j)
			}
		default:
			fmt.Printf("(unknown) - %v => %v -- %T\n", k, value, value)
		}
	}
}
