package main

import "fmt"

type Car2 struct {
	Model string
	Year  int
}

func six() {
	incr := func(car *Car2) { car.Year++ }
	car := Car2{Model: "Kaho", Year: 2009}
	fmt.Println(car.Year)
	incr(&car)
	fmt.Println(car.Year)
}
