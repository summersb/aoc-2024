package main

import (
	"fmt"
	"strings"

	"github.com/summersb/support"
)

func main() {

	reports, err := support.ReadFile("../02/input.txt")
	if err != nil {
		return
	}
	safe := 0
	for i := 0; i < len(reports); i++ {
		data := strings.Fields(reports[i])
		fmt.Println(data)
		// test each line by removing one element
		if IsSafe(data) {
			safe += 1
			continue
		}
		for j := 0; j < len(data); j++ {
			newData := make([]string, len(data))
			copy(newData, data)
			newData = append(newData[:j], newData[j+1:]...)
			if IsSafe(newData) {
				safe += 1
				break
			}
		}
		fmt.Print(data, safe)
		fmt.Println()
	}
	fmt.Println("Safe reports", safe)
}

func IsSafe(data []string) bool {
	dir := 0
	if support.Num(data[0]) > support.Num(data[1]) {
		dir = 1
	}

	if support.Num(data[0]) < support.Num(data[1]) {
		dir = -1
	}
	if dir == 0 {
		return false
	}
	for i := 1; i < len(data); i++ {

		change := support.Num(data[i-1]) - support.Num(data[i])

		if support.Abs(change) > 3 {
			return false
		}
		if change == 0 {
			return false
		}

		if dir == -1 {
			if change > 0 {
				return false
			}
		}
		if dir == 1 {
			if change < 0 {
				return false
			}
		}
	}
	return true
}
