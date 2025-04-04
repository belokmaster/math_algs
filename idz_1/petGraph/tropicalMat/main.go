package main

import (
	"bufio"
	"fmt"
	"image/color"
	"math"
	"os"
	"strconv"
	"strings"

	"github.com/fogleman/gg"
)

// TropicalMultiply выполняет тропическое умножение матриц A и B с учетом нулей на диагонали
func TropicalMultiply(A, B [][]float64) [][]float64 {
	n := len(A)
	C := make([][]float64, n)
	for i := range C {
		C[i] = make([]float64, n)
		for j := range C[i] {
			if i == j {
				C[i][j] = 0 // Главная диагональ всегда 0
			} else {
				C[i][j] = math.Inf(1)
			}
		}
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i == j {
				continue // Диагональ уже заполнена нулями
			}
			for k := 0; k < n; k++ {
				C[i][j] = math.Min(C[i][j], A[i][k]+B[k][j])
			}
		}
	}
	return C
}

// LoadMatrix загружает матрицу и проверяет нули на диагонали
func LoadMatrix(filename string) ([][]float64, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var matrix [][]float64
	scanner := bufio.NewScanner(file)
	lineNum := 0
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "|") || len(line) == 0 {
			continue
		}

		values := strings.Fields(line)
		var row []float64
		for colNum, value := range values {
			num, err := strconv.ParseFloat(value, 64)
			if err != nil {
				return nil, fmt.Errorf("ошибка парсинга числа в строке %d, столбец %d", lineNum+1, colNum+1)
			}
			// Проверяем главную диагональ
			if lineNum == colNum && num != 0 {
				return nil, fmt.Errorf("на главной диагонали (строка %d, столбец %d) должно быть 0", lineNum+1, colNum+1)
			}
			row = append(row, num)
		}
		matrix = append(matrix, row)
		lineNum++
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return matrix, nil
}

// PrintMatrix выводит матрицу с ∞ вместо Inf
func PrintMatrix(matrix [][]float64) {
	for _, row := range matrix {
		for _, value := range row {
			if math.IsInf(value, 1) {
				fmt.Print("∞ ")
			} else {
				fmt.Printf("%.0f ", value)
			}
		}
		fmt.Println()
	}
}

// MatricesEqual сравнивает матрицы, игнорируя диагональ
func MatricesEqual(A, B [][]float64) bool {
	if len(A) != len(B) {
		return false
	}
	for i := range A {
		if len(A[i]) != len(B[i]) {
			return false
		}
		for j := range A[i] {
			if i == j {
				continue // Диагональ не сравниваем (там всегда 0)
			}
			if A[i][j] != B[i][j] {
				return false
			}
		}
	}
	return true
}

// drawTropicalMatrix рисует матрицу с выделенной диагональю
func drawTropicalMatrix(matrix [][]float64, filename string) {
	const (
		cellSize   = 60
		labelSpace = 40
	)

	n := len(matrix)
	width := n*cellSize + 2*labelSpace
	height := n*cellSize + 2*labelSpace

	dc := gg.NewContext(width, height)
	dc.SetColor(color.White)
	dc.Clear()

	// Загружаем шрифт
	if err := dc.LoadFontFace("Arial.ttf", 14); err != nil {
		dc.LoadFontFace("/usr/share/fonts/truetype/dejavu/DejaVuSans.ttf", 14)
	}

	// Подписи вершин
	for i := 0; i < n; i++ {
		// Горизонтальные подписи
		x := float64(labelSpace + i*cellSize + cellSize/2)
		y := float64(labelSpace / 2)
		dc.DrawStringAnchored(fmt.Sprintf("V%d", i), x, y, 0.5, 0.5)

		// Вертикальные подписи
		x = float64(labelSpace / 2)
		y = float64(labelSpace + i*cellSize + cellSize/2)
		dc.DrawStringAnchored(fmt.Sprintf("V%d", i), x, y, 0.5, 0.5)
	}

	// Рисуем матрицу
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			// Цвет фона
			if i == j {
				dc.SetRGB(0.9, 0.9, 0.9) // Серый для диагонали
			} else if !math.IsInf(matrix[i][j], 1) {
				dc.SetRGB(0.8, 0.9, 1.0) // Голубой для значений
			} else {
				dc.SetRGB(1, 1, 1) // Белый для ∞
			}

			x := float64(labelSpace + j*cellSize)
			y := float64(labelSpace + i*cellSize)
			dc.DrawRectangle(x, y, cellSize, cellSize)
			dc.FillPreserve()

			// Границы
			dc.SetRGB(0, 0, 0)
			dc.SetLineWidth(1)
			dc.Stroke()

			// Текст
			value := "∞"
			if !math.IsInf(matrix[i][j], 1) {
				value = fmt.Sprintf("%.0f", matrix[i][j])
			}
			if i == j {
				value = "0" // Принудительно показываем 0 на диагонали
			}
			dc.DrawStringAnchored(value, x+cellSize/2, y+cellSize/2, 0.5, 0.5)
		}
	}

	// Заголовок
	dc.DrawStringAnchored("Тропическая матрица смежности", float64(width)/2, 20, 0.5, 0.5)

	// Сохраняем
	if err := dc.SavePNG(filename); err != nil {
		fmt.Println("Ошибка сохранения:", err)
	} else {
		fmt.Println("Сохранено в", filename)
	}
}

func main() {
	// Загрузка матрицы
	matrix, err := LoadMatrix("adjacency_matrix.txt")
	if err != nil {
		fmt.Println("Ошибка загрузки:", err)
		return
	}

	n := len(matrix)
	current := make([][]float64, n)
	for i := range matrix {
		current[i] = make([]float64, n)
		copy(current[i], matrix[i])
		current[i][i] = 0 // Гарантируем нули на диагонали
	}

	// Сохраняем начальную матрицу
	drawTropicalMatrix(current, "tropical_matrix_step_1.png")

	// Вычисляем степени
	for power := 2; power <= n; power++ {
		next := TropicalMultiply(current, matrix)
		drawTropicalMatrix(next, fmt.Sprintf("tropical_matrix_step_%d.png", power))

		if MatricesEqual(current, next) {
			fmt.Printf("Матрица стабилизировалась на степени %d\n", power-1)
			break
		}
		current = next
	}

	// Вывод результата
	fmt.Println("Финальная матрица:")
	PrintMatrix(current)
}
