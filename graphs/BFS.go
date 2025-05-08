package graphs

// BFS принимает стартовую вершину и выводит вершины в порядке их посещения
func (g *graph) BFS(vertexFrom int) {
	queue := []int{vertexFrom}
	visited := make([]bool, len(g.adjacencyList))
	visited[vertexFrom] = true

	for len(queue) > 0 {
		vertex := queue[0]
		queue = queue[1:]

		println(vertex)

		for _, vertexTo := range g.adjacencyList[vertex] {
			if !visited[vertexTo] {
				queue = append(queue, vertexTo)
				visited[vertexTo] = true
			}
		}
	}
}
