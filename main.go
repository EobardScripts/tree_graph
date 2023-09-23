package main

import (
	"fmt"
	"math/rand"
	"time"
	"tree_graf/graph/graph"
	"tree_graf/graph/indicativeGraph"
	"tree_graf/tree"
)

// Случайное число
func randInt(min int, max int) int {
	return min + rand.Intn(max-min)
}

func treeFabric(n int) []tree.Tree {
	var trees []tree.Tree
	for i := 0; i < n; i++ {
		t := tree.New()
		countData := randInt(1, 100)
		for j := 0; j < countData; j++ {
			t.Insert(randInt(0, 10))
		}
		trees = append(trees, *t)
	}
	return trees
}

func graphFabric(n int) []graph.Graph {
	var graphs []graph.Graph
	for i := 0; i < n; i++ {
		g := graph.New()
		countData := randInt(1, 100)
		for j := 0; j < countData; j++ {
			g.AddEdge(randInt(1, 10), randInt(1, 10))
		}
		graphs = append(graphs, *g)
	}
	return graphs
}

func indGraphsFabric(n int) []indicativeGraph.Graph {
	var graphs []indicativeGraph.Graph
	for i := 0; i < n; i++ {
		g := indicativeGraph.New()
		countData := randInt(1, 100)
		for j := 0; j < countData; j++ {
			g.AddEdge(randInt(1, 10), randInt(1, 10), randInt(1, 10))
		}
		graphs = append(graphs, *g)
	}
	return graphs
}

func main() {
	//Инициализация для случайных чисел
	rand.Seed(time.Now().UnixNano())

	trees := treeFabric(2)
	graphs := graphFabric(2)
	indGraphs := indGraphsFabric(2)

	fmt.Println("Бинарное дерево: ")
	for _, el := range trees {
		el.Print()
		f := randInt(1, 20)
		fmt.Printf("Ищем %d -> %v\n", f, el.Find(f))
		fmt.Printf("Удаляем %d -> %v\n", f, el.Delete(f))
		el.Print()
	}

	fmt.Println("Граф, поиск в ширину: ")
	for _, el := range graphs {
		r := randInt(1, 10)
		fmt.Printf("Стартовый узел %d\n", r)
		el.BFS(r)
		fmt.Println()
	}

	fmt.Println("\nГраф, определение кратчайшего пути: ")
	for _, el := range indGraphs {
		start := randInt(1, 10)
		distances := el.FindMinRoute(start)
		fmt.Printf("Минимальное расстояние от узла %d:\n", start)
		for vertex, distance := range distances {
			fmt.Printf("Узел %d: %d\n", vertex, distance)
		}
	}
}
