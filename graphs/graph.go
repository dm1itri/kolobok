package graphs

import "slices"

type Graph[T ~int] interface {
	AddEdge(vertexFrom, vertexTo T)
	AddOrientedEdge(vertexFrom, vertexTo T)
	RemoveEdge(vertexFrom, vertexTo T)
	RemoveOrientedEdge(vertexFrom, vertexTo T)
	DFS(start T, visited []bool)
	BFS(start T)
	TopologicalSort() []T
	IsEulerian() bool
}

type graph[T ~int] struct {
	adjacencyList [][]T
}

func New[T ~int](size int) Graph[T] {
	return &graph[T]{adjacencyList: make([][]T, size)}
}

func (g *graph[T]) AddEdge(vertexFrom, vertexTo T) {
	g.AddOrientedEdge(vertexFrom, vertexTo)
	g.AddOrientedEdge(vertexTo, vertexFrom)
}

func (g *graph[T]) AddOrientedEdge(vertexFrom, vertexTo T) {
	g.adjacencyList[vertexFrom] = append(g.adjacencyList[vertexFrom], vertexTo)
}
func (g *graph[T]) RemoveEdge(vertexFrom, vertexTo T) {
	g.RemoveOrientedEdge(vertexFrom, vertexTo)
	g.RemoveOrientedEdge(vertexTo, vertexFrom)
}

func (g *graph[T]) RemoveOrientedEdge(vertexFrom, vertexTo T) {
	for i := range g.adjacencyList[vertexFrom] {
		if g.adjacencyList[vertexFrom][i] == vertexTo {
			g.adjacencyList[vertexFrom] = slices.Delete(g.adjacencyList[vertexFrom], i, i+1)
			break
		}
	}
}
