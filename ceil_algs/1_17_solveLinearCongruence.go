package ceil_algs

import (
	"fmt"
	"time"
)

// Функция для решения линейного сравнения ax ≡ b (mod m)
func SolveLinearCongruence(a, b, m int) (bool, int, int) {
	gcd, x0, _ := GcdExtended(a, m)

	if b%gcd != 0 {
		return false, 0, 0 // Решений нет
	}

	// Находим частное решение
	x := (x0 * (b / gcd)) % (m / gcd)

	if x < 0 {
		x += m / gcd
	}

	return true, x, m / gcd
}

// Пример функции для выполнения алгоритма решения линейного сравнения
func ExecuteSolveLinearCongruence(a, b, m int) {
	start := time.Now()
	flag, x, step := SolveLinearCongruence(a, b, m)
	duration := time.Since(start)

	fmt.Printf("1. Время выполнения: %v\n", duration)
	if flag {
		fmt.Printf("2. Решение линейного сравнения %dx ≡ %d (mod %d):\n", a, b, m)
		fmt.Printf("3. x ≡ %d (mod %d)\n", x, step)
		fmt.Println("4. Все решения можно найти по формуле: x = x0 + k * step, где k — целое число.")
	} else {
		fmt.Println("2. Линейное сравнение не имеет решений")
	}
}
