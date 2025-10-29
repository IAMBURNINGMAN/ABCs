package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Order struct {
	Id       int
	Quantity int
}

func Suplier(ordersChain chan<- Order) {
	for i := 1; i <= 5; i = i + 1 {
		order := Order{Id: i, Quantity: rand.Intn(100)}
		ordersChain <- order
		time.Sleep(300 * time.Millisecond)
	}
	close(ordersChain)
}

func wareHouse(ordersChan <-chan Order, resultchain chan<- string) {
	ItemsLeft := 500
	for order := range ordersChan {
		time.Sleep(500 * time.Millisecond)
		if order.Quantity < ItemsLeft {
			ItemsLeft = ItemsLeft - order.Quantity
			resultchain <- fmt.Sprintf("Номер операциии %d, число товаров в операции: %d, остаток на складе: %d", order.Id, order.Quantity, ItemsLeft)
		} else {
			resultchain <- fmt.Sprintf("Нихуя товара нет")
		}

	}
	close(resultchain)
}

func twelwe() {
	ordersChan := make(chan Order)
	resultChan := make(chan string)
	go Suplier(ordersChan)
	go wareHouse(ordersChan, resultChan)
	for msg := range resultChan {
		fmt.Println(msg)
	}
	fmt.Println("All orders accepted")
}
