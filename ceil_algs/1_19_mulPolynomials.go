package ceil_algs

import (
	"fmt"
	"time"
)

// Функция для умножения двух многочленов
func MultiplyPolynomials(poly1, poly2 []int) []int {
	len1 := len(poly1)
	len2 := len(poly2)
	resultLen := len1 + len2 - 1

	// Создаем результирующий многочлен с нулевыми коэффициентами
	result := make([]int, resultLen)

	// Умножаем коэффициенты многочленов
	for i := 0; i < len1; i++ {
		for j := 0; j < len2; j++ {
			result[i+j] += poly1[i] * poly2[j]
		}
	}

	return result
}

// Пример функции для выполнения умножения двух многочленов
func ExecuteMultiplyPolynomials(a, b []int) {
	start := time.Now()
	result := MultiplyPolynomials(a, b)
	duration := time.Since(start)

	fmt.Printf("1. Время выполнения: %v\n", duration)
	fmt.Printf("Многочлен 1: %v\n", a)
	fmt.Printf("Многочлен 2: %v\n", b)
	fmt.Printf("Результат умножения: %v\n", result)
}
