package ceil_algs

// Вспомогательная функция для получения максимального из двух чисел.
func max(a, b int) int {
	if a > b {
		return a // Возвращаем большее значение
	}
	return b // Возвращаем меньшее или равное значение
}