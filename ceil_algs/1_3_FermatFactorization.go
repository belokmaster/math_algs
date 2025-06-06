package ceil_algs

import (
	"fmt"
	"math"
	"time"
)

/*
	Метод Ферма (Fermat's factorization method).
	Алгоритм для факторизации чисел, который основан на представлении числа в виде разности квадратов.
	FermatFactorization выполняет факторизацию числа n методом Ферма.
	Сложность выполнения алгоритма: O(n^1/4), т.к.:
		1 - начальное значение x выбирается как ближайшее целое число, которое больше или равно квадратному корню из n.
		2 - пока y^2 не является точным квадратом, увеличиваем x и пересчитываем y.
	Возвращает два множителя числа n.
	Инвариант: На каждом шаге, если y^2 = x^2 - n, то (x - y) * (x + y) = n.
*/

func FermatFactorization(n int) (int, int) {
	// Если число четное, сразу возвращаем 2 и n/2
	if n%2 == 0 {
		return 2, n / 2
	}

	// Начальное значение для x выбирается как ближайшее целое число,
	// которое больше или равно квадратному корню из n.
	x := int(math.Ceil(math.Sqrt(float64(n))))
	y2 := x*x - n                    // Вычисляем y^2 = x^2 - n
	y := int(math.Sqrt(float64(y2))) // Находим целую часть квадратного корня из y^2

	// Пока y^2 не является точным квадратом, увеличиваем x и пересчитываем y.
	for y*y != y2 {
		x++                             // Увеличиваем x на 1
		y2 = x*x - n                    // Пересчитываем y^2 для нового значения x
		y = int(math.Sqrt(float64(y2))) // Находим новое значение y
	}

	// Возвращаем найденные множители, которые являются (x - y) и (x + y).
	return x - y, x + y
}

// Функция для выполнения алгоритма "Метода факторизации Ферма"
func ExecuteFermatFactorization(n int) {
	start := time.Now()
	divisor_a, divisor_b := FermatFactorization(n)
	duration := time.Since(start)

	fmt.Printf("1. Время выполнения: %v\n", duration)
	fmt.Printf("2. Делители числа %d: %d, %d\n", n, divisor_a, divisor_b)
}
