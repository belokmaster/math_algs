package ceil_algs

import (
	"fmt"
	"math"
	"time"
)

// Решето Эратосфена. Получение всех простых чисел до вводимого значения n.

// SieveOfEratosthenes возвращает срез, содержащий все простые числа до n
func SieveOfEratosthenes(n int) []int {
	// Если n меньше 2, возвращаем пустой срез
	if n < 2 {
		return []int{}
	}

	// Создаем булев срез для отметки составных чисел, начиная с 3
	// Число нечетных чисел до n можно вычислить как (n-1)/2 + 1.
	// Это потому, что в интервале от 1 до n половина чисел являются нечетными.
	isComposite := make([]bool, (n-1)/2+1)

	// Инициализируем простые числа, проходя только по нечетным
	// Массив isComposite предназначен, чтобы помечать составные числа среди нечётных чисел.
	// индекс 0 соответствует числу 1, индекс 1 — числу 3,
	// индекс 2 — числу 5, индекс 3 — числу 7, индекс 4 — числу 9 и т. д.
	for i := 1; i <= (int(math.Sqrt(float64(n)))-1)/2; i++ {
		if !isComposite[i] {
			for j := (2 * i * (i + 1)); j <= (n-1)/2; j += 2*i + 1 {
				isComposite[j] = true
			}
		}
	}

	// Собираем простые числа
	primes := []int{2} // 2 - это единственное четное простое число
	for i := 1; i <= (n-1)/2; i++ {
		if !isComposite[i] {
			primes = append(primes, 2*i+1)
		}
	}

	return primes
}

// Функция для выполнения алгоритма "Решето Эратосфена"
func ExecuteSieveOfEratosthenes(n int) {
	start := time.Now()
	primes := SieveOfEratosthenes(n)
	duration := time.Since(start)

	fmt.Printf("1. Время выполнения: %v\n", duration)
	fmt.Printf("2. Количество простых чисел до %d: %d\n", n, len(primes))
	fmt.Printf("3. Простые числа до %d: %v\n", n, primes)
}
