package ceil_algs

import (
	"fmt"
	"time"
)

// Функция для умножения числа в произвольной системе счисления на цифру
func MultiplyPBaseNumberByDigit(num string, digit int, p int) string {
	// Создаем срез для хранения результата, длина которого будет на 1 больше, чем у исходных чисел (для переноса)
	result := make([]byte, len(num)+1)
	carry := 0 // Переменная для хранения переноса

	// Умножение числа на цифру
	for i := len(num) - 1; i >= 0; i-- {
		// Преобразуем символ цифры числа в целое число
		digit1 := charToInt(num[i])

		// Умножаем цифру на число и добавляем перенос
		product := digit1*digit + carry

		// Записываем остаток от деления на основание системы счисления в результат
		result[i+1] = intToChar(product % p)

		// Обновляем перенос для следующего разряда
		carry = product / p
	}

	// Если есть перенос, записываем его в старший разряд
	if carry > 0 {
		result[0] = intToChar(carry)
	} else {
		// Если переноса нет, удаляем старший разряд
		result = result[1:]
	}

	// Возвращаем результат как строку
	return string(result)
}

// Функция для выполнения алгоритма "умножение числа на цифру в с.с. p"
func ExecuteMultiplyPBaseNumberByDigit(num string, digit int, p int) {
	start := time.Now()
	result := MultiplyPBaseNumberByDigit(num, digit, p)
	duration := time.Since(start)

	fmt.Printf("1. Время выполнения: %v\n", duration)
	fmt.Printf("2. Результат умножения числа %s на цифру %d в %d-ой с.с: %s\n", num, digit, p, result)
}
