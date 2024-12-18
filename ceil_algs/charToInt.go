package ceil_algs

// Вспомогательная функция для преобразования символа в соответствующее целое число
func charToInt(c byte) int {
	if c >= '0' && c <= '9' {
		return int(c - '0')
	} else if c >= 'A' && c <= 'F' {
		return int(c - 'A' + 10)
	}
	return -1 // Ошибка, если символ не является допустимой цифрой
}
