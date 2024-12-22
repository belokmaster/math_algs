package ceil_algs

import (
	"fmt"
	"time"
)

// gcdBinary вычисляет наибольший общий делитель (НОД) двух чисел a и b
// с использованием бинарного алгоритма Евклида
func GcdBinary(a, b int) int {
	// Обработка крайних случаев
	if a == 0 {
		return b
	}
	if b == 0 {
		return a
	}

	// Определяем количество общих множителей 2
	shift := 0
	for (a|b)&1 == 0 { // Пока оба числа четные
		a >>= 1
		b >>= 1
		shift++
	}

	// Приводим a к нечетному числу
	for a&1 == 0 {
		a >>= 1
	}

	for b != 0 {
		// Приводим b к нечетному числу
		for b&1 == 0 {
			b >>= 1
		}

		// Уменьшаем большее из чисел
		if a > b {
			a, b = b, a
		}
		b -= a
	}

	// Восстанавливаем общий множитель 2
	return a << shift
}

// Пример функции для выполнения классического алгоритма Евклида
func ExecuteGcdBinary(a, b int) {
	start := time.Now()
	result := GcdBinary(a, b)
	duration := time.Since(start)

	fmt.Printf("1. Время выполнения: %v\n", duration)
	fmt.Printf("2. НОД чисел %d и %d: %d\n", a, b, result)
}
