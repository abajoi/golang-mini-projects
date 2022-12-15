package main

import (
	"fmt"
	"strings"
)

func main() {
	num := ""
	list := strings.Split(num, ",")

	getData(list)

	// fn(num)
}

func getData(list []string) {
	for i := 0; i < len(list)+1; i++ {
		fmt.Println(list[i])
	}
}
