package main

import (
	"fmt"
	"lab3/calc"
)

func main() {

	fmt.Println("=== Використання функцій пакету ===")

	fmt.Println("Сума:", calc.Sum(1, 2, 3, 4))
	fmt.Println("Максимум:", calc.Max(10, 5, 8))
	fmt.Println("Мінімум:", calc.Min(10, 5, 8))

	result, err := calc.Divide(10, 2)
	if err != nil {
		fmt.Println("Помилка:", err)
	} else {
		fmt.Println("Ділення:", result)
	}

	_, err = calc.Divide(10, 0)
	if err != nil {
		fmt.Println("Помилка:", err)
	}

	fmt.Println("\n=== Використання інтерфейсу Calculator ===")

	var c calc.Calculator = calc.Calc{}

	fmt.Println("Сума:", c.Sum(2, 4, 6))
	fmt.Println("Максимум:", c.Max(7, 3, 9))
	fmt.Println("Мінімум:", c.Min(7, 3, 9))

	result, err = c.Divide(20, 5)
	if err != nil {
		fmt.Println("Помилка:", err)
	} else {
		fmt.Println("Ділення:", result)
	}
}
