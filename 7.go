package main

import "fmt"

type Laptop struct {
	Brand string
	Price int
}

func seven() {
	laptop := []Laptop{
		{Brand: "Apple", Price: 10},
		{Brand: "Lenovo", Price: 20},
		{Brand: "ROG", Price: 30},
	}
	for index, laptop := range laptop {
		fmt.Println(index+1, laptop.Brand, laptop.Price)
	}

	laptop[1].Price++

	fmt.Println()
	fmt.Println(laptop[1].Brand, laptop[1].Price)

	laptop = append(laptop, Laptop{Brand: "ASUS", Price: 5})

	for index, item := range laptop {
		fmt.Println(index+1, ":", item.Brand, item.Price)
	}
}
