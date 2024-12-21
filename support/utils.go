package support

import "strconv"

func Num(data string) int {
	i, err := strconv.Atoi(data)
	if err != nil {
		panic(err)
	}
	return i
}

func Abs(a int) int {
	if a < 0 {
		a = -a
	}
	return a
}
