package graphs

import (
	"reflect"
	"testing"
)

func TestGraphsTopologicalSort(t *testing.T) {
	tests := []struct {
		name string
		args Graph[int]
		want []int
	}{
		{
			name: "one component",
			args: getGraphOneComponent(),
			want: []int{1, 3, 2, 4, 0},
		},
		{
			name: "two component",
			args: getGraphTwoComponent(),
			want: []int{6, 5, 1, 3, 2, 4, 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.args.TopologicalSort()

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DFS() got = %v, want %v", got, tt.want)
			}
		})
	}
}
