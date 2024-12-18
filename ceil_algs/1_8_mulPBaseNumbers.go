package ceil_algs

import (
	"fmt"
	"time"
)

// Функция для умножения чисел в произвольной системе счисления с использованием сдвига
func MulPBaseNumbers(num1, num2 string, p int) string {
	// Результат начнется с "0" (это строка, представляющая число в произвольной системе счисления)
	result := "0"

	// Проходим по каждой цифре второго числа (начиная с младших разрядов)
	// Для этого используем цикл от последней цифры к первой
	for i := len(num2) - 1; i >= 0; i-- {
		// Преобразуем текущую цифру второго числа (из символа в число)
		digit2 := charToInt(num2[i])

		// Умножаем первое число на цифру второго числа (получаем промежуточное произведение)
		partialProduct := MulPBaseNumberByDigit(num1, digit2, p)

		// Сдвигаем результат на i разрядов влево (i — это позиция цифры во втором числе)
		// Каждое произведение нужно сдвигать влево в зависимости от позиции цифры во втором числе
		partialProduct = ShiftPBaseNumberLeft(partialProduct, len(num2)-i-1, p)

		// Складываем текущий результат с частным произведением
		// Используем функцию сложения чисел в произвольной системе счисления
		result = AddPBaseNumbers(result, partialProduct, p)
	}

	// Возвращаем итоговое произведение
	return result
}

// Функция для выполнения алгоритма умножения чисел в произвольной системе счисления
func ExecuteMulPBaseNumbers(num1, num2 string, p int) {
	start := time.Now()
	result := MulPBaseNumbers(num1, num2, p)
	duration := time.Since(start)

	fmt.Printf("1. Время выполнения: %v\n", duration)
	fmt.Printf("2. Результат умножения чисел %s и %s в %d-ой с.с: %s\n", num1, num2, p, result)
}
