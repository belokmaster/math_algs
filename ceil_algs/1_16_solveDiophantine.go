package ceil_algs

import (
	"fmt"
	"time"
)

/*
	Решение диофантова уравнения. Диофантово уравнение имеет вид ax + by = c.
	SolveDiophantine выполняет нахождение целочисленных решений для уравнения.
	Метод основан на расширенном алгоритме Евклида.
	Возвращает флаг наличия решения и значения x и y, удовлетворяющие уравнению.
	Инвариант: ax + by = c.
*/

func SolveDiophantine(a, b, c int) (bool, int, int) {
	// Вычисляем наибольший общий делитель (НОД) и коэффициенты x0 и y0
	d, x0, y0 := GcdExtended(a, b)

	// Если c не делится на d, то уравнение не имеет целочисленных решений
	if c%d != 0 {
		return false, 0, 0
	}

	// Вычисляем частное решение
	x := x0 * (c / d)
	y := y0 * (c / d)
	return true, x, y
}

// Пример функции для выполнения алгоритма решения диофантового уравнения
func ExecuteSolveDiophantine(a, b, c int) {
	start := time.Now()
	flag, x, y := SolveDiophantine(a, b, c)
	duration := time.Since(start)

	fmt.Printf("1. Время выполнения: %v\n", duration)
	if flag {
		fmt.Printf("2. Решение диофантова уравнения %dx + %dy = %d:\n", a, b, c)
		fmt.Printf("2. x = %d, y = %d\n", x, y)
	} else {
		fmt.Println("2. Уравнение не имеет целочисленных решений")
	}
}
