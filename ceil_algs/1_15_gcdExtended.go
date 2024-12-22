package ceil_algs

import (
	"fmt"
	"time"
)

// extendedGCD возвращает НОД двух чисел a и b, а также коэффициенты x и y такие, что ax + by = НОД(a, b)
func GcdExtended(a, b int) (int, int, int) {
	if a == 0 {
		return b, 0, 1
	}

	gcd, x1, y1 := GcdExtended(b%a, a)
	x := y1 - (b/a)*x1
	y := x1

	return gcd, x, y
}

// Пример функции для выполнения классического алгоритма Евклида
func ExecuteGcdExtendedc(a, b int) {
	start := time.Now()
	gcd, x, y := GcdExtended(a, b)
	duration := time.Since(start)

	fmt.Printf("1. Время выполнения: %v\n", duration)
	fmt.Printf("2. НОД(%d) = %d*%d + %d*%d\n", gcd, a, x, b, y)
}
