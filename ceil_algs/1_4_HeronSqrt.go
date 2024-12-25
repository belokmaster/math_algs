package ceil_algs

import (
	"fmt"
	"time"
)

/*
	Итерационная формула Герона (Heron's method).
	Метод вычисления приближенного значения квадратного корня числа.
	HeronSqrt использует итерационную формулу Герона для вычисления целого квадратного корня числа n.
	Сложность выполнения алгоритма: O(log(n)), т.к.:
		1 - каждый шаг итерации уменьшает ошибку вдвое.
	Возвращает целое значение приближенного квадратного корня числа n.
	Инвариант: На каждом шаге значение x приближается к истинному значению квадратного корня числа n.
*/

func HeronSqrt(n int) int {
	// Если число отрицательное, возвращаем -1, так как квадратный корень не определен для отрицательных чисел
	if n < 0 {
		return -1
	}

	// Если число равно 0, возвращаем 0, так как квадратный корень из 0 равен 0
	if n == 0 {
		return 0
	}

	// Начальное приближение для квадратного корня. Обычно используем само число.
	x := n

	// Переменная для хранения предыдущего значения x, чтобы контролировать остановку итерации
	prevX := 0

	// Итерация продолжается до тех пор, пока текущее значение x не стабилизируется
	// Условие завершения цикла - совпадение текущего и предыдущего значения x
	for x != prevX {
		// Сохраняем текущее значение x в prevX перед обновлением
		prevX = x

		// Итерационная формула Герона для обновления значения x: x = (x + n/x) / 2
		// Мы берем среднее арифметическое между текущим значением x и n, деленным на x.
		x = (x + n/x) / 2
	}

	// Когда цикл завершится, x будет содержать приближенное целое значение квадратного корня числа n
	return x
}

// Функция для выполнения алгоритма "итерационной формулы Герона"
func ExecuteHeronSqrt(n int) {
	start := time.Now()
	root := HeronSqrt(n)
	duration := time.Since(start)

	fmt.Printf("1. Время выполнения: %v\n", duration)
	fmt.Printf("2. Квадратный корень числа %d: %d\n", n, root)
}
