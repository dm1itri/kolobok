package graphs

func (g *graph) DFS(vertexFrom int, visited []bool) {
	if visited[vertexFrom] {
		return
	}

	visited[vertexFrom] = true
	for _, vertexTo := range g.adjacencyList[vertexFrom] {
		g.DFS(vertexTo, visited)
	}
}
