package main

import (
	"fmt"
	"strings"

	"github.com/summersb/support"
)

func main() {

	reports, err := support.ReadFile("input.txt")
	if err != nil {
		return
	}

	safe := 0
	for i := 0; i < len(reports); i++ {
		data := strings.Fields(reports[i])
		if isSafe(data) {
			safe += 1
		}
	}

	fmt.Println("Safe reports ", safe)
}

func isSafe(data []string) bool {
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
