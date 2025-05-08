package graphs

type Graph interface {
	AddEdge(vertexFrom, vertexTo int)
	DFS(start int, visited []bool)
	BFS(start int)
}

type graph struct {
	adjacencyList [][]int
}

func New(size int) Graph {
	return &graph{adjacencyList: make([][]int, size)}
}

func (g *graph) AddEdge(vertexFrom, vertexTo int) {
	g.adjacencyList[vertexFrom] = append(g.adjacencyList[vertexFrom], vertexTo)
}
