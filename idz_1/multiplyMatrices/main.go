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

func multiplyMatrices(A, B [][]int) [][]int {
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
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			for k := 0; k < p; k++ {
				C[i][j] += A[i][k] * B[k][j]
			}
		}
	}
	return C
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
			dc.SetRGB(0.9, 0.9, 0.9) // светло-серый цвет
			dc.DrawRectangle(float64(j*cellSize), float64(i*cellSize), cellSize, cellSize)
			dc.FillPreserve()

			dc.SetRGB(0, 0, 0)
			dc.SetLineWidth(2)
			dc.Stroke()

			dc.SetRGB(0, 0, 0)
			dc.DrawStringAnchored(fmt.Sprintf("%d", matrix[i][j]), float64(j*cellSize+cellSize/2), float64(i*cellSize+cellSize/2), 0.5, 0.5)
		}
	}

	dc.SetRGB(0, 0, 0)
	dc.SetLineWidth(4)
	dc.Stroke()

	err := dc.SavePNG(filename)
	if err != nil {
		fmt.Println("Ошибка сохранения изображения:", err)
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())

	// Размерности матриц
	n := 3 // кол-во строк в А
	m := 4 // кол-во столбцов в А и строк в В
	p := 3 // кол-во столбцов в В

	A := generateMatrix(n, p)
	B := generateMatrix(p, m)

	fmt.Println("Матрица A:")
	for _, row := range A {
		fmt.Println(row)
	}

	fmt.Println("\nМатрица B:")
	for _, row := range B {
		fmt.Println(row)
	}

	C := multiplyMatrices(A, B)

	fmt.Println("\nРезультат умножения A * B:")
	for _, row := range C {
		fmt.Println(row)
	}

	drawMatrix(A, "matrix_A.png")
	drawMatrix(B, "matrix_B.png")
	drawMatrix(C, "matrix_C.png")
}
