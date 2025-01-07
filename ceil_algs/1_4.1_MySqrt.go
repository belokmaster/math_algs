package ceil_algs

import (
	"fmt"
	"time"
)

func mySqrt(n int) int {
	left := 0
	right := n + 1
	mid := (left + right) / 2

	for left <= right {
		if mid*mid == n {
			return mid
		}

		// Если квадрат текущего mid больше n,
		// то правильный корень находится левее, уменьшаем правую границу
		if mid*mid > n {
			right = mid - 1
		} else {
			left = mid + 1
		}

		mid = (left + right) / 2
	}

	return mid
}

// Функция для выполнения алгоритма вычисления квадратного корня через бинарный поиск
func ExecuteMySqrt(n int) {
	start := time.Now()
	root := mySqrt(n)
	duration := time.Since(start)

	fmt.Printf("1. Время выполнения: %v\n", duration)
	fmt.Printf("2. Квадратный корень числа %d: %d\n", n, root)
}
