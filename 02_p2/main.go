package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	reports, err := readFile("../02/input.txt")
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
	fmt.Println("Safe reports", safe)

}

func num(data string) int {
	i, err := strconv.Atoi(data)
	if err != nil {
		panic(err)
	}
	return i
}

func abs(a int) int {
	if a < 0 {
		a = -a
	}
	return a
}

func isSafe(data []string) bool {
	dir := 0
	if num(data[0]) > num(data[1]) {
		dir = 1
	}

	if num(data[0]) < num(data[1]) {
		dir = -1
	}
	errorsFound := 0
	if dir == 0 {
		errorsFound += 1
	}
	for i := 1; i < len(data); i++ {
		change := num(data[i-1]) - num(data[i])

		if abs(change) > 3 {
			errorsFound += 1
			continue
		}
		if change == 0 {
			errorsFound += 1
			continue
		}

		if dir == -1 {
			if change > 0 {
				errorsFound += 1
				continue
			}
		}
		if dir == 1 {
			if change < 0 {
				errorsFound += 1
				continue
			}
		}
	}
	fmt.Println(data, errorsFound)
	if errorsFound > 1 {
		return false
	}
	return true

}
func readFile(fileName string) ([]string, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}

	defer file.Close()
	scanner := bufio.NewScanner(file)

	data := []string{}
	for scanner.Scan() {
		line := scanner.Text()
		data = append(data, line)
	}

	return data, nil
}
