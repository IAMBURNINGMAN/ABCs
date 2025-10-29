package main

import (
	"fmt"
	"strconv"
)

type Movable interface {
	Move() string
}

type Caring struct {
	Brand string
	Model string
	Speed int
}

func (c Caring) Move() string {
	return string(c.Brand) + " " + c.Model + " едет со скоростью " + strconv.Itoa(c.Speed)
}

type Person struct {
	Speed int
	Name  string
	Age   int
}

func (p Person) Move() string {
	return string(p.Name) + " " + " идет со скоростью " + strconv.Itoa(p.Speed)
}

func Printmovement(m Movable) {
	fmt.Println(m.Move())
}

func movable(m Movable) {
	fmt.Println(m.Move())
}

func nine() {
	slais := []Movable{
		Person{
			Speed: 20,
			Age:   20,
			Name:  "Chel",
		},
		Caring{
			Speed: 100,
			Brand: "Chel",
			Model: "Cheliks",
		},
	}
	for _, m := range slais {
		movable(m)
	}
}
