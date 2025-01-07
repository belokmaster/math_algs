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

// Функция для чтения ввода и преобразования строки коэффициентов в срез целых чисел
func readPolynomialInput(prompt string) ([]int, error) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(prompt)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input) // Убираем лишние пробелы и символы новой строки

	// Разделяем строку по пробелам и конвертируем в срез целых чисел
	coeffStrs := strings.Split(input, " ")
	coeffs := make([]int, len(coeffStrs))
	for i, coeffStr := range coeffStrs {
		n, err := strconv.Atoi(coeffStr)
		if err != nil {
			return nil, fmt.Errorf("ошибка при преобразовании строки в число: %v", err)
		}
		coeffs[i] = n
	}

	return coeffs, nil
}

// Функция для чтения вещественного числа с ввода
func readFloatInput(prompt string) (float64, error) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(prompt)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input) // Убираем лишние пробелы и символы новой строки

	f, err := strconv.ParseFloat(input, 64)
	if err != nil {
		return 0, fmt.Errorf("ошибка при преобразовании строки в число: %v", err)
	}

	return f, nil
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
	case "1.1":
		n, err = readIntInput("Введите значение для n: ")
		if err != nil {
			fmt.Println(err)
			return
		}
		// Выполнение алгоритма "Решето Эратосфена"
		ceil_algs.ExecuteSimpleSieveOfEratosthenes(n)
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
	case "4.1":
		n, err = readIntInput("Введите значение для n: ")
		if err != nil {
			fmt.Println(err)
			return
		}
		// Выполнение алгоритма вычисления квадратного корня через бинарный поиск
		ceil_algs.ExecuteMySqrt(n)
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
	case "5.1":
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
		// Выполнение алгоритма "вычитание двух чисел в с.с. p"
		ceil_algs.ExecuteSubPBaseNumbers(num1, num2, p)
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
	case "9":
		num1, err := readStringInput("Введите число: ")
		if err != nil {
			fmt.Println(err)
			return
		}
		p, err := readIntInput("Введите основание системы счисления (p): ")
		if err != nil {
			fmt.Println(err)
			return
		}
		q, err := readIntInput("Введите основание системы счисления (q): ")
		if err != nil {
			fmt.Println(err)
			return
		}
		// Выполнение алгоритма "преобразования числа из системы счисления с основанием p в систему с основанием q"
		ceil_algs.ExecutePBaseToQBase(num1, p, q)
	case "11":
		num1, err := readIntInput("Введите число для возведения в степень: ")
		if err != nil {
			fmt.Println(err)
			return
		}
		num2, err := readIntInput("Введите число для степени: ")
		if err != nil {
			fmt.Println(err)
			return
		}
		// Выполнение алгоритма "возведения числа в натуральную степень"
		ceil_algs.ExecutePow(num1, num2)
	case "12":
		num1, err := readFloatInput("Введите число для возведения в степень: ")
		if err != nil {
			fmt.Println(err)
			return
		}
		num2, err := readIntInput("Введите число для степени: ")
		if err != nil {
			fmt.Println(err)
			return
		}
		// Выполнение алгоритма "возведения числа в степень методом быстрого возведения в степень"
		ceil_algs.ExecuteBinPow(num1, num2)
	case "13":
		num1, err := readIntInput("Введите первое число: ")
		if err != nil {
			fmt.Println(err)
			return
		}
		num2, err := readIntInput("Введите второе число: ")
		if err != nil {
			fmt.Println(err)
			return
		}
		// Выполнение классического алгоритма Евклида
		ceil_algs.ExecuteGcdClassic(num1, num2)
	case "14":
		num1, err := readIntInput("Введите первое число: ")
		if err != nil {
			fmt.Println(err)
			return
		}
		num2, err := readIntInput("Введите второе число: ")
		if err != nil {
			fmt.Println(err)
			return
		}
		// Выполнение бинарного алгоритма Евклида
		ceil_algs.ExecuteGcdBinary(num1, num2)
	case "15":
		num1, err := readIntInput("Введите первое число: ")
		if err != nil {
			fmt.Println(err)
			return
		}
		num2, err := readIntInput("Введите второе число: ")
		if err != nil {
			fmt.Println(err)
			return
		}
		// Выполнение обобщенного алгоритма Евклида
		ceil_algs.ExecuteGcdExtendedc(num1, num2)
	case "16":
		a, err := readIntInput("Введите коэффициент a: ")
		if err != nil {
			fmt.Println(err)
			return
		}
		b, err := readIntInput("Введите коэффициент b: ")
		if err != nil {
			fmt.Println(err)
			return
		}
		c, err := readIntInput("Введите коэффициент c: ")
		if err != nil {
			fmt.Println(err)
			return
		}
		// Выполнение алгоритма решения диофантового уравнения
		ceil_algs.ExecuteSolveDiophantine(a, b, c)
	case "17":
		a, err := readIntInput("Введите коэффициент a: ")
		if err != nil {
			fmt.Println(err)
			return
		}
		b, err := readIntInput("Введите коэффициент b: ")
		if err != nil {
			fmt.Println(err)
			return
		}
		m, err := readIntInput("Введите коэффициент m: ")
		if err != nil {
			fmt.Println(err)
			return
		}
		// Выполнение линейного сравнения
		ceil_algs.ExecuteSolveLinearCongruence(a, b, m)
	case "18":
		a, err := readPolynomialInput("Введите первый полином: ")
		if err != nil {
			fmt.Println(err)
			return
		}
		b, err := readPolynomialInput("Введите второй полином: ")
		if err != nil {
			fmt.Println(err)
			return
		}
		// Выполнение алгоритма сложения двух полиномов
		ceil_algs.ExecuteAddPolynomials(a, b)
	case "18.1":
		a, err := readPolynomialInput("Введите первый полином: ")
		if err != nil {
			fmt.Println(err)
			return
		}
		b, err := readPolynomialInput("Введите второй полином: ")
		if err != nil {
			fmt.Println(err)
			return
		}
		// Выполнение алгоритма вычитания двух полиномов
		ceil_algs.ExecuteSubtractPolynomials(a, b)
	case "19":
		a, err := readPolynomialInput("Введите первый полином: ")
		if err != nil {
			fmt.Println(err)
			return
		}
		b, err := readPolynomialInput("Введите второй полином: ")
		if err != nil {
			fmt.Println(err)
			return
		}
		// Выполнение алгоритма умножения двух полиномов
		ceil_algs.ExecuteMultiplyPolynomials(a, b)
	case "20":
		a, err := readPolynomialInput("Введите первый полином: ")
		if err != nil {
			fmt.Println(err)
			return
		}
		b, err := readPolynomialInput("Введите второй полином: ")
		if err != nil {
			fmt.Println(err)
			return
		}
		// Выполнение алгоритма умножения двух полиномов
		ceil_algs.ExecuteDividePolynomials(a, b)
	}
}
