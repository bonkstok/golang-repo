package main

import (
	"fmt"
	"strings"
)

func main() {
	hdr := []string{"empid", "employee", "address", "hours worked", "hourly rate", "manager"}
	result := csvHdrCol(hdr)
	fmt.Println("Result:")
	fmt.Println(result)
}

func csvHdrCol(hdr []string) map[int]string {
	csvIdxToCol := make(map[int]string)

	for i, v := range hdr {
		v = strings.TrimSpace(v)
		switch strings.ToLower(v) {
		case "employee":
			csvIdxToCol[i] = v
		case "hours worked":
			csvIdxToCol[i] = v
		case "hourly rate":
			csvIdxToCol[i] = v
		}
	}
	return csvIdxToCol
}
