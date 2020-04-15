package main

import "fmt"

func getType(v interface{}) string {
	switch t := v.(type) {
	case string:
		return fmt.Sprintf("Input is of type %T", t)
	case int, int32:
		return fmt.Sprintf("Input is of type %T", t)
	case float32, float64:
		return fmt.Sprintf("Input is of type %T", t)
	case bool:
		return fmt.Sprintf("Input is of type %T", t)
	default:
		return "UNKNOWN"
	}
}

func main() {
	fmt.Println(getType(5))
	fmt.Println(getType("hey"))
	fmt.Println(getType(true))
	fmt.Println(getType(99.99))
}
