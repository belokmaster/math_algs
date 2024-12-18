package ceil_algs

import "strings"

// Вспомогательная функция для преобразования массива цифр в строку
func intArrayToString(arr []byte) string {
	var sb strings.Builder
	for _, b := range arr {
		sb.WriteByte(intToChar(int(b)))
	}
	return sb.String()
}
