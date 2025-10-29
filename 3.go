package main

import (
	"fmt"
	"time"
)

func tri() {
	for i := 1; i < 11; i++ {
		fmt.Print(i, " ")
	}
	fmt.Println()
	j := 0
	for {
		fmt.Print(j, " ")
		time.Sleep(500 * time.Millisecond)
		j++
		if j == 20 {
			break
		}
	}

}
