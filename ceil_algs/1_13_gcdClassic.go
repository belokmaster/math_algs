package ceil_algs

import (
	"fmt"
	"time"
)

// gcd вычисляет наибольший общий делитель (НОД) двух чисел a и b
func GcdClassic(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

// Пример функции для выполнения классического алгоритма Евклида
func ExecuteGcdClassic(a, b int) {
	start := time.Now()
	result := GcdClassic(a, b)
	duration := time.Since(start)

	fmt.Printf("1. Время выполнения: %v\n", duration)
	fmt.Printf("2. НОД чисел %d и %d: %d\n", a, b, result)
}
