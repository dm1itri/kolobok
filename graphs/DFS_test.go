package graphs

import (
	"reflect"
	"testing"
)

func getGraphTwoComponent() Graph[int] {
	graph := New[int](7)
	graph.AddEdge(1, 0)
	graph.AddEdge(1, 2)
	graph.AddEdge(1, 3)
	graph.AddEdge(3, 2)
	graph.AddEdge(2, 4)
	graph.AddEdge(6, 5)

	return graph
}

func getGraphOneComponent() Graph[int] {
	graph := New[int](5)
	graph.AddEdge(1, 0)
	graph.AddEdge(1, 2)
	graph.AddEdge(1, 3)
	graph.AddEdge(3, 2)
	graph.AddEdge(2, 4)

	return graph
}

func TestGraphsDFS(t *testing.T) {
	type args struct {
		vertexFrom int
		visited    []bool
	}

	tests := []struct {
		name string
		args args
		want []bool
	}{
		{
			name: "in one component",
			args: args{
				vertexFrom: 3,
				visited:    make([]bool, 5),
			},
			want: []bool{false, false, true, true, true},
		},
		{
			name: "in two component",
			args: args{
				vertexFrom: 1,
				visited:    make([]bool, 7),
			},
			want: []bool{true, true, true, true, true, false, false},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			graph := getGraphTwoComponent()

			graph.DFS(tt.args.vertexFrom, tt.args.visited)

			if !reflect.DeepEqual(tt.args.visited, tt.want) {
				t.Errorf("DFS() got = %v, want %v", tt.args.visited, tt.want)
			}
		})
	}
}
