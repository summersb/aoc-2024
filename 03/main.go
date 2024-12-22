package main

import (
	"fmt"
	"strings"

	"github.com/summersb/support"
)

var debug = false

func main() {
	// read input
	reports, err := support.ReadFile("input.txt")
	if err != nil {
		return
	}
	oneLine := strings.Join(reports[:], "")
	Part1(oneLine)
	Part2(oneLine)
}

func Part1(reports string) {
	sum := ProcessRow(reports, false)
	fmt.Println("Part 1 Total", sum)
}

func Part2(reports string) {
	sum := ProcessRow(reports, true)
	fmt.Println("Part 2 Total", sum)
}

func ProcessRow(data string, control bool) int {
	sum := 0
	enabled := true
	for pos := 0; pos < len(data); pos++ {
		mulIdx := support.IndexOf(data, "mul(", pos)
		if mulIdx == -1 {
			// if no mul is found end
			break
		}
		if control {
			// if do is found, enable
			doIdx := support.IndexOf(data, "do()", pos)
			dontIdx := support.IndexOf(data, "don't()", pos)
			if debug {
				fmt.Println(mulIdx, doIdx, dontIdx)
			}
			if (dontIdx != -1 && doIdx < dontIdx) && doIdx < mulIdx && doIdx != -1 {
				enabled = true
				if debug {
					fmt.Println("Enabled", doIdx)
				}
			}
			// if don't is found, disable
			if dontIdx < mulIdx && dontIdx != -1 {
				enabled = false
				if debug {
					fmt.Println("Disabled", dontIdx)
				}
			}
		}

		// find ) then process
		endIdx := support.IndexOf(data, ")", mulIdx)
		if endIdx == -1 {
			if debug {
				fmt.Println("Error parsing mul")
			}
			continue
		}
		//subString := data[idx : endIdx+1]
		commaIdx := support.IndexOf(data, ",", mulIdx)
		if commaIdx == -1 {
			if debug {
				fmt.Println("Error parsing mul")
			}
			continue
		}
		endIdx = support.IndexOf(data, ")", commaIdx)
		if endIdx == -1 {
			if debug {
				fmt.Println("Error parsing mul")
			}
			continue
		}
		first, err := support.Num(data[mulIdx+4 : commaIdx])
		if err != nil {
			if debug {
				fmt.Println("Error parsing mul")
			}
			continue
		}
		second, err := support.Num(data[commaIdx+1 : endIdx])
		if err != nil {
			if debug {
				fmt.Println("Error parsing mul")
			}
			continue
		}
		if debug {
			fmt.Println(enabled, pos, first, "*", second)
		}
		if (control && enabled) || !control {
			sum += (first * second)
		}
		pos = endIdx
	}
	return sum
}
