package graphs

import "kolobok/utils"

func (g *graph[T]) IsEulerian() bool {
	g.fleury(0)

	for vertexFrom := range g.adjacencyList {
		if len(g.adjacencyList[vertexFrom]) != 0 {
			return false
		}
	}

	return true
}

// fleury алгоритм поиска эйлерова цикла в неориентированном графк
// ассимптотика по времени O(E^2)
func (g *graph[T]) fleury(vertexFrom T) {
	for len(g.adjacencyList[vertexFrom]) > 0 {
		for _, vertexTo := range g.adjacencyList[vertexFrom] {
			if len(g.adjacencyList[vertexFrom]) == 1 || !g.isBridge(vertexFrom, vertexTo) {
				g.RemoveEdge(vertexFrom, vertexTo)
				g.fleury(vertexTo)

				break
			}
		}
	}
}

func (g *graph[T]) isBridge(vertexFrom, vertexTo T) bool {
	visited := make([]bool, len(g.adjacencyList))
	g.DFS(vertexFrom, visited)
	visitedCount := utils.CountInSlice(visited, true)
	vertexFromEdges := g.adjacencyList[vertexFrom]
	vertexToEdges := g.adjacencyList[vertexTo]
	g.RemoveEdge(vertexFrom, vertexTo)
	visited = make([]bool, len(g.adjacencyList))
	g.DFS(vertexFrom, visited)
	g.adjacencyList[vertexFrom] = vertexFromEdges
	g.adjacencyList[vertexTo] = vertexToEdges

	return utils.CountInSlice(visited, true) != visitedCount
}
