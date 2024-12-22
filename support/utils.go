package support

import (
	"strconv"
	"strings"
)

func Num(data string) (int, error) {
	i, err := strconv.Atoi(data)
	if err != nil {
		return 0, err
	}
	return i, nil
}

func Abs(a int) int {
	if a < 0 {
		a = -a
	}
	return a
}

func IndexOf(data string, find string, pos int) int {
	idx := strings.Index(data[pos:], find)
	if idx > -1 {
		idx += pos
	}

	return idx
}
