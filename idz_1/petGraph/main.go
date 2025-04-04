package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"

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

	// Выводим матрицу смежности
	printAdjacencyMatrix(g, numNodes)

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
