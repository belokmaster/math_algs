package ceil_algs

import (
	"fmt"
	"strings"
	"time"
)

// Вспомогательная функция для получения максимального из двух чисел. Я ЖЕ НА ГОЛАНГЕ
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// Функция для сложения двух чисел в произвольной системе счисления
func AddPBaseNumbers(num1, num2 string, p int) string {
	// Находим длину самой длинной строки, чтобы выровнять числа
	maxLen := max(len(num1), len(num2))

	// Дополняем более короткое число нулями слева, чтобы длина обоих чисел стала одинаковой
	num1 = strings.Repeat("0", maxLen-len(num1)) + num1
	num2 = strings.Repeat("0", maxLen-len(num2)) + num2

	// Создаем срез для хранения результата, длина которого будет на 1 больше, чем у исходных чисел (для переноса)
	result := make([]byte, maxLen+1)
	carry := 0 // Переменная для хранения переноса

	// Сложение чисел с учетом переноса
	for i := maxLen - 1; i >= 0; i-- {
		// Преобразуем символы чисел в целые числа
		digit1 := int(num1[i] - '0') // Преобразуем символ в число
		digit2 := int(num2[i] - '0')

		// Суммируем цифры, добавляем перенос из предыдущего разряда
		sum := digit1 + digit2 + carry
		// Записываем результат в соответствующую позицию среза, используя модуль p для получения цифры в данной системе счисления
		result[i+1] = byte(sum%p) + '0'
		// Обновляем перенос для следующей цифры
		carry = sum / p
	}

	// Учет возможного переноса в старший разряд
	if carry > 0 {
		result[0] = byte(carry) + '0' // Если перенос существует, записываем его в старший разряд
	} else {
		// Если переноса нет, удаляем первый элемент среза (старший разряд)
		result = result[1:]
	}

	return string(result)
}

// Функция для выполнения алгоритма "сложение двух чисел в с.с. p"
func ExecuteAddPBaseNumbers(num1, num2 string, p int) {
	start := time.Now()
	result := AddPBaseNumbers(num1, num2, p)
	duration := time.Since(start)

	fmt.Printf("1. Время выполнения: %v\n", duration)
	fmt.Printf("2. Сумма чисел %s и %s в %d-ой с.с: %s\n", num1, num2, p, result)
}
