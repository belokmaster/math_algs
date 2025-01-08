package ceil_algs

import (
	"fmt"
	"time"
)

// Функция для разложения многочлена на множители свободные от квадратов
func FactorizeFreeFromSquares(polynomial []int) ([][]int, error) {
	// Если степень многочлена 0, то он уже является свободным от квадратов
	if len(polynomial) == 1 {
		return [][]int{polynomial}, nil
	}

	var factors [][]int
	currPolynomial := append([]int(nil), polynomial...)

	for {
		if len(currPolynomial) == 0 {
			break
		}

		// Выполняем деление
		quotient, remainder, err := DividePolynomials(currPolynomial, []int{1, 1})
		if err != nil {
			return nil, err
		}

		if len(remainder) == len(currPolynomial) {
			break
		}

		if len(remainder) > 1 || (len(remainder) == 1 && remainder[0] != 0) {
			factors = append(factors, quotient)
			currPolynomial = remainder
		} else {
			break
		}
	}

	return factors, nil
}

// Пример использования функции для выполнения разложения многочлена
func ExecuteFactorizeFreeFromSquares(polynomial []int) {
	start := time.Now()
	factors, err := FactorizeFreeFromSquares(polynomial)
	if err != nil {
		fmt.Println("Ошибка:", err)
		return
	}
	duration := time.Since(start)

	fmt.Printf("1. Время выполнения: %v\n", duration)
	fmt.Printf("Исходный многочлен: %v\n", polynomial)
	fmt.Printf("Разложение на свободные от квадратов множители: %v\n", factors)
}
