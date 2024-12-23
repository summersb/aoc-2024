package support

import "fmt"

var debug = false

func SetDebug(d bool) {
	debug = d
}

func Log(msg ...any) {
	if debug {
		fmt.Print(msg...)
	}
}

func Logln(msg ...any) {
	if debug {
		fmt.Println(msg...)
	}
}
