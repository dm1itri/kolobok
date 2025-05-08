package graphs

import (
	"reflect"
	"testing"
)

func TestGraphsIsEulerian(t *testing.T) {
	tests := []struct {
		name string
		args Graph[int]
		want bool
	}{
		{
			name: "one component",
			args: getGraphOneComponent(),
			want: false,
		},
		{
			name: "two component",
			args: getGraphTwoComponent(),
			want: false,
		},
		{
			name: "is eurelian",
			args: getGraphIsEurelian(),
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.args.IsEulerian()

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("IsEurelean() got = %v, want %v", got, tt.want)
			}
		})
	}
}
