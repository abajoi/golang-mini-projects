package main

import (
	"flag"
	"fmt"
	"os"
	"time"
)

func main() {

	// // 2022-12-09T14:22:00+00:00
	deadline := flag.String("deadline", "", "The deadline for the countdown timer in RFC3339 format (e.g. 2019-12-25T15:00:00+01:00)")
	// flag.Parse()

	v, err := time.Parse(time.RFC3339, *deadline)
	if err != nil && v.Day() != 10 {
		if v.Day() != 10 {
			fmt.Println(err.Error())
			os.Exit(1)
		} else {
			fmt.Println("HEY")
		}

	}
	fmt.Println(v)
	
	os.close()
}
