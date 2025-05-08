package graphs

func getGraphTwoComponent() Graph[int] {
	graph := New[int](7)
	graph.AddOrientedEdge(1, 0)
	graph.AddOrientedEdge(1, 2)
	graph.AddOrientedEdge(1, 3)
	graph.AddOrientedEdge(3, 2)
	graph.AddOrientedEdge(2, 4)
	graph.AddOrientedEdge(6, 5)

	return graph
}

func getGraphOneComponent() Graph[int] {
	graph := New[int](5)
	graph.AddOrientedEdge(1, 0)
	graph.AddOrientedEdge(1, 2)
	graph.AddOrientedEdge(1, 3)
	graph.AddOrientedEdge(3, 2)
	graph.AddOrientedEdge(2, 4)

	return graph
}

func getGraphIsEurelian() Graph[int] {
	graph := New[int](6)
	graph.AddEdge(3, 2)
	graph.AddEdge(2, 0)
	graph.AddEdge(0, 1)
	graph.AddEdge(1, 2)
	graph.AddEdge(2, 4)
	graph.AddEdge(4, 5)
	graph.AddEdge(5, 3)
	graph.AddEdge(3, 4)

	return graph
}
