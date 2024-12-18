package ceil_algs

import (
	"fmt"
	"strings"
	"time"
)

// Функция для вычитания двух чисел в произвольной системе счисления
func SubtractPBaseNumbers(num1, num2 string, p int) string {
	// Находим длину самой длинной строки, чтобы выровнять числа
	maxLen := max(len(num1), len(num2))

	// Дополняем более короткое число нулями слева, чтобы длина обоих чисел стала одинаковой
	num1 = strings.Repeat("0", maxLen-len(num1)) + num1
	num2 = strings.Repeat("0", maxLen-len(num2)) + num2

	// Создаем срез для хранения результата, длина которого будет такой же, как у исходных чисел
	result := make([]byte, maxLen)
	borrow := 0 // Переменная для хранения заимствования (borrow)

	// Вычитание чисел с учетом заимствования
	for i := maxLen - 1; i >= 0; i-- {
		// Преобразуем символы чисел в целые числа
		digit1 := charToInt(num1[i]) // Преобразуем символ в число
		digit2 := charToInt(num2[i])

		// Вычитаем цифры, учитывая заимствование из предыдущего разряда
		diff := digit1 - digit2 - borrow
		if diff < 0 {
			// Если результат отрицателен, учитываем заимствование
			diff += p
			borrow = 1
		} else {
			borrow = 0
		}

		// Записываем результат в соответствующую позицию среза
		result[i] = intToChar(diff)
	}

	// Убираем ведущие нули, если они есть
	return strings.TrimLeft(string(result), "0")
}

// Функция для выполнения алгоритма "вычитание двух чисел в с.с. p"
func ExecuteSubPBaseNumbers(num1, num2 string, p int) {
	start := time.Now()
	result := SubtractPBaseNumbers(num1, num2, p)
	duration := time.Since(start)

	fmt.Printf("1. Время выполнения: %v\n", duration)
	fmt.Printf("2. Разность чисел %s и %s в %d-ой с.с: %s\n", num1, num2, p, result)
}
