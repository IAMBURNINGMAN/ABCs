package main

import (
	"fmt"
	"strings"
)

func Dss(numbers []int) []int {
	countofpos := 0
	sumnegat := 0
	for _, value := range numbers {
		if value > 0 {
			countofpos++
		} else {
			sumnegat += value
		}
	}
	res := []int{countofpos, sumnegat}
	return res
}
func main() {
	NumbersRes := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, -11, -12, -13, -14, -15}
	fmt.Println(Dss(NumbersRes))
	strings.Repeat()
}
