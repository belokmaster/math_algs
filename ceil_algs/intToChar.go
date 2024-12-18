package ceil_algs

// Вспомогательная функция для преобразования цифры в строку в соответствующий символ в системе счисления
func intToChar(digit int) byte {
	if digit < 10 {
		return byte(digit + '0')
	} else {
		return byte(digit - 10 + 'A') // Для чисел 10-15 возвращаем символы A-F
	}
}
