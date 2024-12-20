package ceil_algs

import (
	"fmt"
	"time"
)

/*
	Pow возводит число base в степень exponent итеративным методом.
	Сложность выполнения алгоритма: O(n)
	Ничего такого..
*/

func Pow(base, exponent int) int {
	result := 1
	for i := 0; i < exponent; i++ {
		result *= base
	}
	return result
}

// Пример функции для выполнения возведения числа в степень.
func ExecutePow(base, exponent int) {
	start := time.Now()
	result := Pow(base, exponent)
	duration := time.Since(start)

	fmt.Printf("1. Время выполнения: %v\n", duration)
	fmt.Printf("2. Результат возведения числа %d в %d-ую в степень: %d\n", base, exponent, result)
}
