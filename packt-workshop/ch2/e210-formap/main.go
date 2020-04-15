package main

import "fmt"

func main() {
	config := map[string]string{
		"debug":    "1",
		"logLevel": "warn",
		"version":  "1.2.1",
	}

	for key, value := range config { //order is random
		fmt.Println(key, "=", value)
	}
}
