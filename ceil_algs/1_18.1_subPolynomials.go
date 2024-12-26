package ceil_algs

import (
	"fmt"
	"math"
	"time"
)

// Функция для вычитания двух многочленов
func SubtractPolynomials(poly1, poly2 []int) []int {
	len1 := len(poly1)
	len2 := len(poly2)
	maxLen := int(math.Max(float64(len1), float64(len2)))

	// Выровняем длины многочленов, добавив нули к началу более короткого
	if len1 < maxLen {
		padding := make([]int, maxLen-len1)
		poly1 = append(padding, poly1...)
	} else if len2 < maxLen {
		padding := make([]int, maxLen-len2)
		poly2 = append(padding, poly2...)
	}

	// Создаем результирующий многочлен с нулевыми коэффициентами
	result := make([]int, maxLen)

	// Вычитаем коэффициенты многочленов
	for i := 0; i < maxLen; i++ {
		result[i] = poly1[i] - poly2[i]
	}

	return result
}

// Пример функции для выполнения вычитания двух многочленов
func ExecuteSubtractPolynomials(a, b []int) {
	start := time.Now()
	result := SubtractPolynomials(a, b)
	duration := time.Since(start)

	fmt.Printf("1. Время выполнения: %v\n", duration)
	fmt.Printf("Многочлен 1: %v\n", a)
	fmt.Printf("Многочлен 2: %v\n", b)
	fmt.Printf("Результат вычитания: %v\n", result)
}
