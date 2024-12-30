package ceil_algs

import (
	"fmt"
	"time"
)

/*
	Алгоритм Евклида - классический алгоритм нахождения наибольшего общего делителя (НОД) двух чисел.
	GcdClassic вычисляет НОД двух целых чисел a и b.
	Сложность выполнения алгоритма: O(log(min(a, b))), т.к. на каждом шаге уменьшается одно из чисел.
	Инвариант: После каждой итерации одно из чисел уменьшается, приближая их к НОД.
*/

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
