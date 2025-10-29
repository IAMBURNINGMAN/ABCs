package main

import (
	con "Basic/Converter"
	matem "Basic/mathpack"
	"fmt"
)

func eleven() {
	fmt.Println(matem.IsPrime(7), matem.IsPrime(8))
	fmt.Println(matem.Pow(123, 1), matem.Pow(2, 5))

	fmt.Println(con.CelsiusToFahrenheit(36), con.KilometersToMiles(100))
}
