package ceil_algs

import (
	"fmt"
	"strings"
	"time"
)

// Функция для сдвига числа в произвольной системе счисления на k разрядов влево
func ShiftPBaseNumberLeft(num string, k int, p int) string {
	// Создаем новое число, добавив k нулей в конец исходного числа
	return num + strings.Repeat("0", k)
}

// Функция для выполнения алгоритма "сдвиг числа влево на k разрядов в с.с. p"
func ExecuteShiftPBaseNumberLeft(num string, k int, p int) {
	start := time.Now()
	result := ShiftPBaseNumberLeft(num, k, p)
	duration := time.Since(start)

	fmt.Printf("1. Время выполнения: %v\n", duration)
	fmt.Printf("2. Число %s после сдвига на %d разрядов влево в %d-ой с.с: %s\n", num, k, p, result)
}
