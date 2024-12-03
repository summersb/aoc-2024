package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {

	s1, s2, err := readList("../01/input.txt")
	if err != nil {
		return
	}

	sort.Slice(s1, func(i, j int) bool {
		return s1[i] < s1[j]
	})
	sort.Slice(s2, func(i, j int) bool {
		return s2[i] < s2[j]
	})

	sum := 0
	for i := 0; i < len(s1); i++ {
		temp := 0
		for j := 0; j < len(s2); j++ {
			if s1[i] == s2[j] {
				temp += 1
			}
		}
		sum += s1[i] * temp
	}

	fmt.Println("Total score", sum)
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
