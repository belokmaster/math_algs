package ceil_algs

import (
	"fmt"
	"math"
	"time"
)

/*
	Метод пробных делителей. Метод факторизации, получение простых делителей числа.
	TrialDivision выполняет разложение числа n на простые множители.
	Сложность выполнения алгоритма: O(n^1/2), т.к.:
		1 - проверка делимости числа на 2. В худшем случае деление дает O(logn)
		2 -
	Алгоритм работает по принципу исключения составных чисел.
	Возвращает слайс простых множителей числа n.
*/

func TrialDivision(n int) []int {
	// Инициализация пустого слайса для хранения множителей
	ans := []int{}

	// Если число меньше 2, возвращаем пустой слайс, так как числа меньше 2 не имеют простых множителей
	if n < 2 {
		return ans
	}

	// Проверяем делимость числа n на 2
	// Если n делится на 2, добавляем 2 в список множителей и делим n на 2 до тех пор, пока это возможно
	for n%2 == 0 {
		ans = append(ans, 2) // Добавляем 2 в список множителей
		n /= 2               // Делим n на 2
	}

	// Проверяем делимость числа n на нечётные числа, начиная с 3 и до квадратного корня из n
	// Шаг увеличивается на 2, чтобы пропускать чётные числа, которые уже были проверены
	for i := 3; i <= int(math.Sqrt(float64(n))); i += 2 {
		// Пока n делится на текущее нечётное число i, добавляем i в список множителей и делим n на i
		for n%i == 0 {
			ans = append(ans, i) // Добавляем i в список множителей
			n /= i               // Делим n на i
		}
	}

	// Если после всех делений n всё ещё больше 2, это значит, что n само является простым числом
	// Добавляем n в список множителей
	if n > 2 {
		ans = append(ans, n)
	}

	// Возвращаем список простых множителей
	return ans
}

// Функция для выполнения алгоритма "Метод пробных делителей"
func ExecuteTrialDivision(n int) {
	start := time.Now()
	divisors := TrialDivision(n)
	duration := time.Since(start)

	fmt.Printf("1. Время выполнения: %v\n", duration)
	fmt.Printf("2. Количество делителей: %d\n", len(divisors))
	fmt.Printf("3. Делители числа %d: %v\n", n, divisors)
}
