package main

import (
	"bufio"
	"fmt"
	"math_algs/ceil_algs"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	// Считываем строку с ввода
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Введите команду: ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input) // Убираем лишние пробелы и символы новой строки

	// Преобразуем строку в целое число
	n, err := strconv.Atoi(input)
	if err != nil {
		fmt.Println("Ошибка при преобразовании строки в число:", err)
		return
	}

	// Используем switch case для выбора действия
	switch input {
	case "1":
		// После выбора команды 1, запрашиваем новое значение для n
		fmt.Print("Введите значение для n: ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input) // Убираем лишние пробелы и символы новой строки

		// Преобразуем строку в целое число
		n, err = strconv.Atoi(input)
		if err != nil {
			fmt.Println("Ошибка при преобразовании строки в число:", err)
			return
		}

		// Выполняем функцию SieveOfEratosthenes для числа n
		// Сохраняем текущее время перед началом выполнения алгоритма
		start := time.Now()
		primes := ceil_algs.SieveOfEratosthenes(n)
		duration := time.Since(start)
		// Выводим время выполнения
		fmt.Printf("Время выполнения: %v\n", duration)

		// Также можно вывести количество простых чисел для проверки
		fmt.Printf("Количество простых чисел: %d\n", len(primes))
		fmt.Printf("Простые числа до %d: %v\n", n, primes)
	default:
		fmt.Println("Неизвестная команда")
	}
}
