package main

import (
	"fmt"
	"image/color"
	"math/rand"
	"time"

	"github.com/fogleman/gg"
)

func generateMatrix(n, m int) [][]int {
	matrix := make([][]int, n)
	for i := range matrix {
		matrix[i] = make([]int, m)
		for j := range matrix[i] {
			matrix[i][j] = rand.Intn(10)
		}
	}
	return matrix
}

func inputMatrix(n, m int) [][]int {
	matrix := make([][]int, n)
	for i := range matrix {
		matrix[i] = make([]int, m)
		for j := range matrix[i] {
			fmt.Printf("Введите элемент [%d][%d]: ", i, j)
			fmt.Scan(&matrix[i][j])
		}
	}
	return matrix
}

func tropicalMultiplyMatrices(A, B [][]int) [][]int {
	n := len(A)
	m := len(B[0])
	p := len(B)

	if len(A[0]) != p {
		fmt.Println("Матрицы нельзя умножить: несоответствие размеров.")
		return nil
	}

	C := make([][]int, n)
	for i := range C {
		C[i] = make([]int, m)
		for j := range C[i] {
			C[i][j] = 1 << 31
		}
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			for k := 0; k < p; k++ {
				C[i][j] = min(C[i][j], A[i][k]+B[k][j])
			}
		}
	}
	return C
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func drawMatrix(matrix [][]int, filename string) {
	const cellSize = 90
	width := len(matrix[0]) * cellSize
	height := len(matrix) * cellSize

	dc := gg.NewContext(width, height)

	dc.SetColor(color.White)

	dc.Clear()
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			dc.SetRGB(0.9, 0.9, 0.9)
			dc.DrawRectangle(float64(j*cellSize), float64(i*cellSize), cellSize, cellSize)
			dc.FillPreserve()
			dc.SetRGB(0, 0, 0)
			dc.SetLineWidth(2)
			dc.Stroke()
			dc.DrawStringAnchored(fmt.Sprintf("%d", matrix[i][j]), float64(j*cellSize+cellSize/2), float64(i*cellSize+cellSize/2), 0.5, 0.5)
		}
	}
	dc.SavePNG(filename)
}

func main() {
	rand.Seed(time.Now().UnixNano())
	var choice int

	fmt.Println("Выберите способ ввода матриц:")
	fmt.Println("1 - Ввести вручную")
	fmt.Println("2 - Сгенерировать автоматически")
	fmt.Scan(&choice)

	n, m, p := 3, 4, 3
	var A, B [][]int

	if choice == 1 {
		fmt.Println("Введите матрицу A:")
		A = inputMatrix(n, p)
		fmt.Println("Введите матрицу B:")
		B = inputMatrix(p, m)
	} else {
		A = generateMatrix(n, p)
		B = generateMatrix(p, m)
	}

	fmt.Println("Матрица A:")
	for _, row := range A {
		fmt.Println(row)
	}

	fmt.Println("Матрица B:")
	for _, row := range B {
		fmt.Println(row)
	}

	C := tropicalMultiplyMatrices(A, B)

	fmt.Println("Результат тропического умножения A * B:")
	for _, row := range C {
		fmt.Println(row)
	}
	drawMatrix(A, "matrix_A.png")
	drawMatrix(B, "matrix_B.png")
	drawMatrix(C, "matrix_C.png")
}
