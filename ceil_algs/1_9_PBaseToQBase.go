package ceil_algs

import (
	"fmt"
	"time"
)

// Функция для преобразования числа из системы счисления с основанием p в систему с основанием q
func PBaseToQBase(num string, p, q int) string {
	// Шаг 1: Преобразуем число из системы счисления p в массив цифр
	// Создадим срез для хранения числа в обратном порядке
	var digits []int
	for i := len(num) - 1; i >= 0; i-- {
		digits = append(digits, charToInt(num[i]))
	}

	// Шаг 2: Преобразуем в систему счисления q
	var result []byte
	for len(digits) > 0 {
		// Получаем остаток от деления на q
		remainder := 0
		for i := len(digits) - 1; i >= 0; i-- {
			// Умножаем текущую цифру на основание p
			digits[i] += remainder * p
			// Получаем новую цифру и остаток от деления на q
			remainder = digits[i] % q
			digits[i] = digits[i] / q
		}
		// Добавляем остаток в результат
		if remainder < 10 {
			result = append(result, byte(remainder+'0'))
		} else {
			result = append(result, byte(remainder-10+'A'))
		}
		// Удаляем ведущие нули
		for len(digits) > 0 && digits[len(digits)-1] == 0 {
			digits = digits[:len(digits)-1]
		}
	}

	// Шаг 3: Переворачиваем результат, так как он был построен в обратном порядке
	if len(result) == 0 {
		return "0"
	}

	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}

	return string(result)
}

// Пример функции для выполнения перевода с замером времени
func ExecutePBaseToQBase(num string, p, q int) {
	start := time.Now()
	result := PBaseToQBase(num, p, q)
	duration := time.Since(start)

	fmt.Printf("1. Время выполнения: %v\n", duration)
	fmt.Printf("2. Результат перевода числа %s из системы счисления с основанием %d в систему счисления с основанием %d: %s\n", num, p, q, result)
}
