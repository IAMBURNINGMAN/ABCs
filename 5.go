package main

import "fmt"

type Car3 struct {
	Model string
	Year  int
}

func yearapdate(p Car3) {
	p.Year++
	fmt.Println(p.Model, p.Year)
}

func five() {
	car := Car3{Model: "Kaho", Year: 2009}
	yearapdate(car)
	fmt.Println(car.Year)
}
