package main

import (
	"fmt"

	"github.com/summersb/support"
)

func main() {
	lines, err := support.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	support.SetDebug(true)
	Part1(lines)
}

func Part1(lines []string) {
	// Part 1
	for _, row := range lines {
		fmt.Println(row)
	}

	count := 0
	//loop over rows
	for y, row := range lines {
		//loop over columns
		for x, col := range row {
			if string(col) == "X" {
				//if row,col == X, check all 8 directions for M
				foundCount := checkForWord(lines, x, y, "XMAS")
				count += foundCount
			}
		}
	}
	fmt.Println("Count", count)
	//if M, check next char for A in same dir
	//if A, check next char for S in same dir
}

func checkForWord(table []string, col int, row int, word string) int {
	//check for word in all 8 directions
	//up
	count := 0
	found := true
	for i := 0; i < len(word); i++ {
		if row-i < 0 {
			found = false
			break
		}
		if table[row-i][col] != word[i] {
			found = false
			break
		}
	}
	if found {
		fmt.Println("word at (up)", row, col, found)
		count++
	}
	// down
	found = true
	for i := 0; i < len(word); i++ {
		if row+i > len(table)-1 {
			found = false
			break
		}
		if table[row+i][col] != word[i] {
			found = false
			break
		}
	}
	if found {
		fmt.Println("word at (down)", row, col, found)
		count++
	}
	// right
	found = true
	for i := 0; i < len(word); i++ {
		if col+i > len(table[0])-1 {
			found = false
			break
		}
		if table[row][col+i] != word[i] {
			found = false
			break
		}
	}
	if found {
		fmt.Println("word at (right)", row, col, found)
		count++
	}
	// left
	found = true
	for i := 0; i < len(word); i++ {
		if col-i < 0 {
			found = false
			break
		}
		if table[row][col-i] != word[i] {
			found = false
			break
		}
	}
	if found {
		fmt.Println("word at (left)", row, col, found)
		count++
	}
	// up-right
	found = true
	for i := 0; i < len(word); i++ {
		if row-i < 0 {
			found = false
			break
		}
		if col+i > len(table[0])-1 {
			found = false
			break
		}
		if table[row-i][col+i] != word[i] {
			found = false
			break
		}
	}
	if found {
		fmt.Println("word at (up-right)", row, col, found)
		count++
	}
	// down-right
	found = true
	for i := 0; i < len(word); i++ {
		if row+i > len(table)-1 {
			found = false
			break
		}
		if col+i > len(table[0])-1 {
			found = false
			break
		}
		if table[row+i][col+i] != word[i] {
			found = false
			break
		}
	}
	if found {
		fmt.Println("word at (down-right)", row, col, found)
		count++
	}
	// down-left
	found = true
	for i := 0; i < len(word); i++ {
		if row+i > len(table)-1 {
			found = false
			break
		}
		if col-i < 0 {
			found = false
			break
		}
		if table[row+i][col-i] != word[i] {
			found = false
			break
		}
	}
	if found {
		fmt.Println("word at (down-left)", row, col, found)
		count++
	}
	// up-left
	found = true
	for i := 0; i < len(word); i++ {
		if row-i < 0 {
			found = false
			break
		}
		if col-i < 0 {
			found = false
			break
		}
		if table[row-i][col-i] != word[i] {
			found = false
			break
		}
	}
	if found {
		fmt.Println("word at (up-left)", row, col, found)
		count++
	}
	return count
}
