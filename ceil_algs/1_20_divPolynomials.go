package ceil_algs

import (
	"errors"
	"fmt"
	"time"
)

// Функция для деления двух многочленов с проверкой делителя
func DividePolynomials(dividend, divisor []int) ([]int, []int, error) {
	// Проверка на деление на нулевой многочлен
	if len(divisor) == 0 || divisor[len(divisor)-1] == 0 {
		return nil, nil, errors.New("делитель не может быть нулевым многочленом")
	}

	// Определяем степень делимого и делителя
	degDividend := len(dividend) - 1
	degDivisor := len(divisor) - 1

	// Если степень делителя больше степени делимого, результат деления - нулевой многочлен
	if degDividend < degDivisor {
		return []int{0}, dividend, nil
	}

	// Результирующий многочлен для частного
	quotient := make([]int, degDividend-degDivisor+1)

	// Копируем делимое, чтобы работать с остатком
	remainder := make([]int, len(dividend))
	copy(remainder, dividend)

	// Пока степень остатка больше или равна степени делителя
	for i := degDividend; i >= degDivisor; i-- {
		// Коэффициент для текущей степени частного
		quotient[i-degDivisor] = remainder[i] / divisor[degDivisor]

		// Сдвигаем делитель и умножаем на коэффициент
		for j := 0; j <= degDivisor; j++ {
			remainder[i-j] -= quotient[i-degDivisor] * divisor[degDivisor-j]
		}
	}

	// Удаляем старшие нули в остатке
	for len(remainder) > 0 && remainder[0] == 0 {
		remainder = remainder[1:]
	}

	// Возвращаем частное и остаток
	return quotient, remainder, nil
}

// Пример функции для выполнения деления двух многочленов
func ExecuteDividePolynomials(a, b []int) {
	start := time.Now()
	quotient, remainder, err := DividePolynomials(a, b)
	if err != nil {
		fmt.Println("Ошибка:", err)
		return
	}
	duration := time.Since(start)

	fmt.Printf("1. Время выполнения: %v\n", duration)
	fmt.Printf("Многочлен 1 (делимое): %v\n", a)
	fmt.Printf("Многочлен 2 (делитель): %v\n", b)
	fmt.Printf("Частное: %v\n", quotient)
	fmt.Printf("Остаток: %v\n", remainder)
}
