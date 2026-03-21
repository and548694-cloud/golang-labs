package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var y float64
	x := rand.Float64()*70 - 10 // [-10; 60)

	if x >= 1 && x <= 50 {
		if x*x-4 == 0 {
			fmt.Println("Помилка: ділення на нуль")
			return
		}
		y = x / (x*x - 4)
	} else {
		y = 2 * x * x
	}

	fmt.Printf("x: %.2f\n", x)
	fmt.Printf("Result: %.4f\n", y)
}
