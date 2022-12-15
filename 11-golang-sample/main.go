package main

import (
	"fmt"
	"os"
)

func main() {
	num := []int(nil)
	fn(num)
}

func Log(msg string, level int) {
	fmt.Println(msg)
	if level == 10 {
		os.Exit(1)
	}
}

func Fatal(msg string) {
	Log(msg, 10)
}

func fn(x []int) {
	if x == nil {
		Fatal("unexpected nil pointer")
	}
	fmt.Println(x)
}
