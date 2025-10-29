package mathpack

import "math"

func IsPrime(num int) bool {
	counter := 0
	for i := 2; i <= int(math.Sqrt(float64(num))); i++ {
		if num%i == 0 {
			counter++
		}
	}
	if counter >= 1 {
		return false
	} else {
		return true
	}
}

func Pow(base, exponent int) int {
	return int(math.Pow(float64(base), float64(exponent)))
}
