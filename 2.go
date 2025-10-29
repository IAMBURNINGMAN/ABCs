package main

import "fmt"

func vtoraia() {
	IsSunny := false
	IsWeekend := false
	switch {
	case IsSunny && IsWeekend:
		fmt.Println("Идеальный день для прогулки")
	case IsSunny && !IsWeekend:
		fmt.Println("Погода хорошая но нужно поработать")
	case !IsSunny && IsWeekend:
		fmt.Println("Можно остаться дома и отдохнуть")
	default:
		fmt.Println("Рабочий день с плохой погодой")
	}

}
