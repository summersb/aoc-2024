package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func abs(a, b int) int {
	val := a - b
	if val < 0 {
		val = -val
	}
	return val
}

func main() {
	fmt.Println("AOC Day 1")
	filename := "input.txt"
	stack1, stack2, err := readList(filename)

	if err != nil {
		fmt.Println("Error loading input file", err)
		return
	}

	// read list 1, sort ascending
	sort.Slice(stack1, func(i, j int) bool {
		return stack1[i] < stack1[j]
	})
	// read list 2, sort ascending
	sort.Slice(stack2, func(i, j int) bool {
		return stack2[i] < stack2[j]
	})

	sum := 0
	for i := 0; i < len(stack1); i++ {
		fmt.Printf("%d %d %d\n", stack1[i], stack2[i], abs(stack1[i], stack2[i]))
		sum = sum + abs(stack1[i], stack2[i])
	}

	fmt.Println("Sum ", sum)
	// walk both lists, add difference between list1[i] and list2[i]
	// return sum
}

func readList(fileName string) ([]int, []int, error) {

	file, err := os.Open(fileName)
	if err != nil {
		return nil, nil, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	stack1 := []int{}
	stack2 := []int{}

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Fields(line)
		i, err := strconv.Atoi(parts[0])
		if err != nil {
			return nil, nil, err
		}
		stack1 = append(stack1, i)
		i, err = strconv.Atoi(parts[1])
		if err != nil {
			return nil, nil, err
		}
		stack2 = append(stack2, i)
	}

	if err := scanner.Err(); err != nil {
		return nil, nil, err
	}
	return stack1, stack2, nil
}
