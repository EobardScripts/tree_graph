package graph

import "fmt"

type Graph struct {
	adList map[int][]int
}

func New() *Graph {
	return &Graph{adList: make(map[int][]int)}
}

func (g *Graph) AddEdge(u, v int) {
	g.adList[u] = append(g.adList[u], v)
}

func (g *Graph) BFS(start int) {
	visited := make(map[int]bool)
	queue := []int{start}
	visited[start] = true

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		fmt.Printf("%d ", node)
		neighbors := g.adList[node]
		for _, neighbor := range neighbors {
			if !visited[neighbor] {
				queue = append(queue, neighbor)
				visited[neighbor] = true
			}
		}
	}
}
