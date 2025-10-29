package main

import (
	"errors"
	"fmt"
)

func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("Cannot divide by zero")
	}
	return a / b, nil
}

func ValidateAge(a int) error {
	if a <= 0 || a > 120 {
		return fmt.Errorf("Ошибка: %w", ErrValidateAge)
	}
	return nil

}

var ErrValidateAge = errors.New("Возраст некорректен")

func ten() {
	err := ValidateAge(12)
	if errors.Is(err, ErrValidateAge) {
		fmt.Println(err)
	} else {
		fmt.Println("Возраст корректен")
	}

	result, err := divide(10, 0)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}

}
