package graphs

import "slices"

// TopologicalSort корректно работает только на направленном ациклическом графе
// визуализация https://www.cs.usfca.edu/~galles/visualization/TopoSortDFS.html
func (g *graph[T]) TopologicalSort() []T {
	visited := make([]bool, len(g.adjacencyList))
	order := make([]T, 0, len(g.adjacencyList))

	for vertex := range g.adjacencyList {
		if !visited[vertex] {
			g.dfsForTopSort(T(vertex), visited, &order)
		}
	}

	slices.Reverse(order)

	return order
}

func (g *graph[T]) dfsForTopSort(vertexFrom T, visited []bool, order *[]T) {
	if visited[vertexFrom] {
		return
	}

	visited[vertexFrom] = true
	for _, vertexTo := range g.adjacencyList[vertexFrom] {
		g.dfsForTopSort(vertexTo, visited, order)
	}

	*order = append(*order, vertexFrom)
}
