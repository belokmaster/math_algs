package main

import (
	"fmt"
	"image/color"
	"math/rand"
	"os"
	"time"

	"github.com/fogleman/gg"
	"gonum.org/v1/gonum/graph/encoding/dot"
	"gonum.org/v1/gonum/graph/simple"
)

func shuffle(slice []int64) {
	for i := range slice {
		j := rand.Intn(i + 1)
		slice[i], slice[j] = slice[j], slice[i]
	}
}

func printAdjacencyMatrix(g *simple.DirectedGraph, numNodes int) {
	fmt.Println("\nМатрица смежности:")

	// Печатаем заголовок с номерами вершин
	fmt.Print("    ")
	for i := 0; i < numNodes; i++ {
		fmt.Printf("%d ", i)
	}
	fmt.Println()

	// Печатаем разделительную линию
	fmt.Print("   ")
	for i := 0; i < numNodes; i++ {
		fmt.Print("--")
	}
	fmt.Println()

	// Печатаем саму матрицу
	for i := 0; i < numNodes; i++ {
		fmt.Printf("%d | ", i)
		for j := 0; j < numNodes; j++ {
			if g.HasEdgeFromTo(int64(i), int64(j)) {
				fmt.Print("1 ")
			} else {
				fmt.Print("0 ")
			}
		}
		fmt.Println()
	}
}

func saveAdjacencyMatrixToFile(g *simple.DirectedGraph, numNodes int, filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	// Записываем заголовок
	fmt.Fprint(file, "    ")
	for i := 0; i < numNodes; i++ {
		fmt.Fprintf(file, "%d ", i)
	}
	fmt.Fprintln(file)

	// Записываем разделительную линию
	fmt.Fprint(file, "   ")
	for i := 0; i < numNodes; i++ {
		fmt.Fprint(file, "--")
	}
	fmt.Fprintln(file)

	// Записываем саму матрицу
	for i := 0; i < numNodes; i++ {
		fmt.Fprintf(file, "%d | ", i)
		for j := 0; j < numNodes; j++ {
			if g.HasEdgeFromTo(int64(i), int64(j)) {
				fmt.Fprint(file, "1 ")
			} else {
				fmt.Fprint(file, "0 ")
			}
		}
		fmt.Fprintln(file)
	}

	return nil
}

func drawTropicalAdjacencyMatrixImage(g *simple.DirectedGraph, numNodes int, filename string) {
	const (
		cellSize   = 60
		labelSpace = 40
	)

	width := numNodes*cellSize + 2*labelSpace
	height := numNodes*cellSize + 2*labelSpace

	dc := gg.NewContext(width, height)

	// Белый фон
	dc.SetColor(color.White)
	dc.Clear()

	// Устанавливаем шрифт
	if err := dc.LoadFontFace("Arial.ttf", 14); err != nil {
		dc.LoadFontFace("/usr/share/fonts/truetype/dejavu/DejaVuSans.ttf", 14)
	}

	// Рисуем подписи вершин по горизонтали
	for i := 0; i < numNodes; i++ {
		x := float64(labelSpace + i*cellSize + cellSize/2)
		y := float64(labelSpace / 2)
		dc.DrawStringAnchored(fmt.Sprintf("V%d", i), x, y, 0.5, 0.5)
	}

	// Рисуем подписи вершин по вертикали
	for i := 0; i < numNodes; i++ {
		x := float64(labelSpace / 2)
		y := float64(labelSpace + i*cellSize + cellSize/2)
		dc.DrawStringAnchored(fmt.Sprintf("V%d", i), x, y, 0.5, 0.5)
	}

	// Рисуем саму тропическую матрицу
	for i := 0; i < numNodes; i++ {
		for j := 0; j < numNodes; j++ {
			// Цвет фона ячейки
			if g.HasEdgeFromTo(int64(i), int64(j)) {
				dc.SetRGB(0.8, 0.9, 1.0) // светло-голубой для ненулевых значений
			} else {
				dc.SetRGB(1, 1, 1) // белый для нулевых значений
			}

			x := float64(labelSpace + j*cellSize)
			y := float64(labelSpace + i*cellSize)
			dc.DrawRectangle(x, y, cellSize, cellSize)
			dc.FillPreserve()

			// Границы ячеек
			dc.SetRGB(0, 0, 0)
			dc.SetLineWidth(1)
			dc.Stroke()

			// Текст в ячейках (∞ для отсутствующих рёбер)
			dc.SetRGB(0, 0, 0)
			value := "∞"
			if i == j { // Если на главной диагонали, ставим 0
				value = "0"
			} else if g.HasEdgeFromTo(int64(i), int64(j)) {
				value = "1"
			}
			dc.DrawStringAnchored(value,
				x+cellSize/2,
				y+cellSize/2,
				0.5, 0.5)
		}
	}

	// Заголовок
	dc.SetRGB(0, 0, 0)
	dc.DrawStringAnchored("Тропическая матрица смежности",
		float64(width)/2,
		20,
		0.5, 0.5)

	// Сохраняем изображение
	err := dc.SavePNG(filename)
	if err != nil {
		fmt.Println("Ошибка сохранения изображения:", err)
	} else {
		fmt.Println("Тропическая матрица смежности сохранена как изображение в", filename)
	}
}

func drawAdjacencyMatrixImage(g *simple.DirectedGraph, numNodes int, filename string) {
	const (
		cellSize   = 60
		labelSpace = 40
	)

	width := numNodes*cellSize + 2*labelSpace
	height := numNodes*cellSize + 2*labelSpace

	dc := gg.NewContext(width, height)

	// Белый фон
	dc.SetColor(color.White)
	dc.Clear()

	// Устанавливаем шрифт
	if err := dc.LoadFontFace("Arial.ttf", 14); err != nil {
		dc.LoadFontFace("/usr/share/fonts/truetype/dejavu/DejaVuSans.ttf", 14)
	}

	// Рисуем подписи вершин по горизонтали
	for i := 0; i < numNodes; i++ {
		x := float64(labelSpace + i*cellSize + cellSize/2)
		y := float64(labelSpace / 2)
		dc.DrawStringAnchored(fmt.Sprintf("V%d", i), x, y, 0.5, 0.5)
	}

	// Рисуем подписи вершин по вертикали
	for i := 0; i < numNodes; i++ {
		x := float64(labelSpace / 2)
		y := float64(labelSpace + i*cellSize + cellSize/2)
		dc.DrawStringAnchored(fmt.Sprintf("V%d", i), x, y, 0.5, 0.5)
	}

	// Рисуем саму матрицу
	for i := 0; i < numNodes; i++ {
		for j := 0; j < numNodes; j++ {
			// Цвет фона ячейки
			if g.HasEdgeFromTo(int64(i), int64(j)) {
				dc.SetRGB(0.8, 0.9, 1.0) // светло-голубой для ненулевых значений
			} else {
				dc.SetRGB(1, 1, 1) // белый для нулевых значений
			}

			x := float64(labelSpace + j*cellSize)
			y := float64(labelSpace + i*cellSize)
			dc.DrawRectangle(x, y, cellSize, cellSize)
			dc.FillPreserve()

			// Границы ячеек
			dc.SetRGB(0, 0, 0)
			dc.SetLineWidth(1)
			dc.Stroke()

			// Текст в ячейках
			dc.SetRGB(0, 0, 0)
			value := "0"
			if g.HasEdgeFromTo(int64(i), int64(j)) {
				value = "1"
			}
			dc.DrawStringAnchored(value,
				x+cellSize/2,
				y+cellSize/2,
				0.5, 0.5)
		}
	}

	// Заголовок
	dc.SetRGB(0, 0, 0)
	dc.DrawStringAnchored("Матрица смежности",
		float64(width)/2,
		20,
		0.5, 0.5)

	// Сохраняем изображение
	err := dc.SavePNG(filename)
	if err != nil {
		fmt.Println("Ошибка сохранения изображения:", err)
	} else {
		fmt.Println("Матрица смежности сохранена как изображение в", filename)
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())

	numNodes := 10
	g := simple.NewDirectedGraph()

	// Добавляем вершины
	for i := int64(0); i < int64(numNodes); i++ {
		g.AddNode(simple.Node(i))
	}

	// Создаём случайный Гамильтонов путь
	nodes := make([]int64, numNodes)
	for i := int64(0); i < int64(numNodes); i++ {
		nodes[i] = i
	}
	shuffle(nodes)

	// Добавляем рёбра Гамильтонова пути
	for i := 0; i < numNodes-1; i++ {
		g.SetEdge(simple.Edge{F: simple.Node(nodes[i]), T: simple.Node(nodes[i+1])})
	}

	// Добавляем случайные дополнительные рёбра (не образующие Гамильтонов цикл)
	existingEdges := make(map[[2]int64]bool)
	for i := 0; i < numNodes-1; i++ {
		existingEdges[[2]int64{nodes[i], nodes[i+1]}] = true
	}

	// Не добавляем ребро из последней вершины в первую — чтобы не образовать цикл
	maxExtraEdges := 10
	for i := 0; i < maxExtraEdges; i++ {
		from := rand.Int63n(int64(numNodes))
		to := rand.Int63n(int64(numNodes))

		if from == to || existingEdges[[2]int64{from, to}] || (from == nodes[numNodes-1] && to == nodes[0]) {
			continue
		}

		g.SetEdge(simple.Edge{F: simple.Node(from), T: simple.Node(to)})
		existingEdges[[2]int64{from, to}] = true
	}

	// Выводим матрицу смежности в консоль
	printAdjacencyMatrix(g, numNodes)

	// Сохраняем матрицу смежности в текстовый файл
	err := saveAdjacencyMatrixToFile(g, numNodes, "adjacency_matrix.txt")
	if err != nil {
		fmt.Printf("\nОшибка при сохранении матрицы смежности: %v\n", err)
	} else {
		fmt.Printf("\nМатрица смежности сохранена в файл: adjacency_matrix.txt\n")
	}

	// Сохраняем матрицы смежностей как изображения
	drawAdjacencyMatrixImage(g, numNodes, "adjacency_matrix.png")
	drawTropicalAdjacencyMatrixImage(g, numNodes, "tropical_adjacency_matrix.png")

	// Экспорт в DOT-формат
	dotData, err := dot.Marshal(g, "HamiltonianPathGraph", "", "  ")
	if err != nil {
		panic(err)
	}

	// Сохраняем в файл
	err = os.WriteFile("peterson.dot", dotData, 0644)
	if err != nil {
		panic(err)
	}

	fmt.Println("\nDOT файл успешно создан: peterson.dot")
}
