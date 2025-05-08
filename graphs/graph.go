package graphs

type Graph[T ~int] interface {
	AddEdge(vertexFrom, vertexTo T)
	DFS(start T, visited []bool)
	BFS(start T)
	TopologicalSort() []T
}

type graph[T ~int] struct {
	adjacencyList [][]T
}

func New[T ~int](size int) Graph[T] {
	return &graph[T]{adjacencyList: make([][]T, size)}
}

func (g *graph[T]) AddEdge(vertexFrom, vertexTo T) {
	g.adjacencyList[vertexFrom] = append(g.adjacencyList[vertexFrom], vertexTo)
}
