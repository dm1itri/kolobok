package graphs

import (
	"reflect"
	"testing"
)

func GetGraph() Graph {
	graph := New(10)
	graph.AddEdge(0, 1)
	graph.AddEdge(1, 2)
	graph.AddEdge(2, 3)
	graph.AddEdge(3, 4)
	graph.AddEdge(7, 8)

	return graph
}

func TestGraphsDFS(t *testing.T) {
	type args struct {
		vertexFrom int
		visited    []bool
	}

	tests := []struct {
		name    string
		args    args
		want    []bool
		wantErr bool
	}{
		{
			name: "in one component",
			args: args{
				vertexFrom: 1,
				visited:    make([]bool, 10),
			},
			want: []bool{false, true, true, true, true, false, false, false, false, false},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			graph := GetGraph()

			graph.DFS(tt.args.vertexFrom, tt.args.visited)

			if !reflect.DeepEqual(tt.args.visited, tt.want) {
				t.Errorf("DFS() got = %v, want %v", tt.args.visited, tt.want)
			}
		})
	}
}
