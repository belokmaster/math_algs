package main

import (
	"bufio"
	"fmt"
	"math_algs/ceil_algs"
	"os"
	"strconv"
	"strings"
)

// Функция для чтения целого числа с ввода
func readIntInput(prompt string) (int, error) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(prompt)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input) // Убираем лишние пробелы и символы новой строки

	n, err := strconv.Atoi(input)
	if err != nil {
		return 0, fmt.Errorf("ошибка при преобразовании строки в число: %v", err)
	}

	return n, nil
}

// Функция для чтения строки ввода
func readStringInput(prompt string) (string, error) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(prompt)
	text, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(text), nil
}

func main() {
	// Читаем команду с ввода
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Введите команду: ")
	command, _ := reader.ReadString('\n')
	command = strings.TrimSpace(command) // Убираем лишние пробелы и символы новой строки

	var n int
	var err error

	switch command {
	case "1":
		n, err = readIntInput("Введите значение для n: ")
		if err != nil {
			fmt.Println(err)
			return
		}
		// Выполнение алгоритма "Решето Эратосфена"
		ceil_algs.ExecuteSieveOfEratosthenes(n)
	case "2":
		n, err = readIntInput("Введите значение для n: ")
		if err != nil {
			fmt.Println(err)
			return
		}
		// Выполнение алгоритма "Метод пробных делителей"
		ceil_algs.ExecuteTrialDivision(n)
	case "3":
		n, err = readIntInput("Введите значение для n: ")
		if err != nil {
			fmt.Println(err)
			return
		}
		// Выполнение алгоритма "Метод факторизации Ферма"
		ceil_algs.ExecuteFermatFactorization(n)
	case "4":
		n, err = readIntInput("Введите значение для n: ")
		if err != nil {
			fmt.Println(err)
			return
		}
		// Выполнение алгоритма "итерационная формула Герона"
		ceil_algs.ExecuteHeronSqrt(n)
	case "5":
		num1, err := readStringInput("Введите первое число: ")
		if err != nil {
			fmt.Println(err)
			return
		}
		num2, err := readStringInput("Введите второе число: ")
		if err != nil {
			fmt.Println(err)
			return
		}
		p, err := readIntInput("Введите основание системы счисления (p): ")
		if err != nil {
			fmt.Println(err)
			return
		}
		// Выполнение алгоритма "сложение двух чисел в с.с. p"
		ceil_algs.ExecuteAddPBaseNumbers(num1, num2, p)
	case "6":
		num1, err := readStringInput("Введите первое число: ")
		if err != nil {
			fmt.Println(err)
			return
		}
		num2, err := readIntInput("Введите второе число: ")
		if err != nil {
			fmt.Println(err)
			return
		}
		p, err := readIntInput("Введите основание системы счисления (p): ")
		if err != nil {
			fmt.Println(err)
			return
		}
		// Выполнение алгоритма "умножение числа на цифру в с.с. p"
		ceil_algs.ExecuteMulPBaseNumberByDigit(num1, num2, p)
	case "7":
		num, err := readStringInput("Введите число для сдвига: ")
		if err != nil {
			fmt.Println(err)
			return
		}
		k, err := readIntInput("Введите количество разрядов для сдвига: ")
		if err != nil {
			fmt.Println(err)
			return
		}
		p, err := readIntInput("Введите основание системы счисления (p): ")
		if err != nil {
			fmt.Println(err)
			return
		}
		// Выполнение алгоритма "сдвиг числа влево на k разрядов в с.с. p"
		ceil_algs.ExecuteShiftPBaseNumberLeft(num, k, p)
	case "8":
		num1, err := readStringInput("Введите первое число: ")
		if err != nil {
			fmt.Println(err)
			return
		}
		num2, err := readStringInput("Введите второе число: ")
		if err != nil {
			fmt.Println(err)
			return
		}
		p, err := readIntInput("Введите основание системы счисления (p): ")
		if err != nil {
			fmt.Println(err)
			return
		}
		// Выполнение алгоритма "умножение чисел в с.с. p"
		ceil_algs.ExecuteMulPBaseNumbers(num1, num2, p)
	}
}
