package indicativeGraph

import "math"

type Edge struct {
	destination int
	weight      int
}

type Graph struct {
	adList map[int][]Edge
}

func New() *Graph {
	return &Graph{adList: make(map[int][]Edge)}
}

func (g *Graph) AddEdge(source, destination, weight int) {
	edge := Edge{
		destination: destination,
		weight:      weight,
	}

	g.adList[source] = append(g.adList[source], edge)
}

func minDistance(d map[int]int, v map[int]bool) int {
	min := math.MaxInt64
	minVertex := -1

	for vertex, distance := range d {
		if !v[vertex] && distance < min {
			min = distance
			minVertex = vertex
		}
	}
	return minVertex
}

func (g *Graph) FindMinRoute(start int) map[int]int {
	distances := make(map[int]int)
	for vertex := range g.adList {
		distances[vertex] = math.MaxInt64
	}

	distances[start] = 0
	visited := make(map[int]bool)

	for len(visited) < len(g.adList) {
		minVertex := minDistance(distances, visited)
		visited[minVertex] = true
		edges := g.adList[minVertex]
		for _, edge := range edges {
			if !visited[edge.destination] {
				newDistance := distances[minVertex] + edge.weight
				if newDistance < distances[edge.destination] {
					distances[edge.destination] = newDistance
				}
			}
		}
	}

	return distances
}
