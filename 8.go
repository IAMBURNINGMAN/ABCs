package main

import "fmt"

func eight() {
	menu := map[string]float64{
		"govno":           160,
		"govno na lopate": 260,
		"govno v rot":     360,
	}

	menu["ssaki"] = 100

	fmt.Println(menu)

	value, ok := menu["салат"]
	fmt.Println(value, ok)
	menu["салат"] = 1200

	delete(menu, "govno v rot")

	for index, price := range menu {
		fmt.Println(index, "стоит", price, "рэ")
	}
}
