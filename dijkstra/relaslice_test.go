package dijkstra

import (
	"math/rand"
	"reflect"
	"testing"
)

var testRoadMap = RoadMap{
	{from: 0, to: 1, weight: 8},
	{from: 0, to: 6, weight: 5},
	{from: 0, to: 9, weight: 132},
	{from: 1, to: 2, weight: 13},
	{from: 1, to: 8, weight: 9},
	{from: 2, to: 3, weight: 3},
	{from: 2, to: 5, weight: 21},
	{from: 3, to: 4, weight: 12},
	{from: 4, to: 3, weight: 12},
	{from: 4, to: 5, weight: 21},
	{from: 5, to: 13, weight: 9},
	{from: 6, to: 7, weight: 35},
	{from: 7, to: 11, weight: 100},
	{from: 8, to: 5, weight: 54},
	{from: 8, to: 11, weight: 65},
	{from: 9, to: 11, weight: 18},
	{from: 12, to: 11, weight: 3},
	{from: 12, to: 10, weight: 12},
	{from: 13, to: 4, weight: 45},
	{from: 13, to: 12, weight: 7},
}

func Test_explorePathsFrom(t *testing.T) {
	type fields struct {
		path Path
	}
	tests := []struct {
		name   string
		fields fields
		want   []Path
	}{
		{
			name: "from_zero",
			fields: fields{
				path: Path{
					path: []int{0},
					cost: 0,
				},
			},
			want: []Path{
				{path: []int{0, 1}, cost: 8},
				{path: []int{0, 6}, cost: 5},
				{path: []int{0, 9}, cost: 132},
			},
		},
		{
			name: "from_point",
			fields: fields{
				path: Path{
					path: []int{0, 1, 8},
					cost: 17,
				},
			},
			want: []Path{
				{path: []int{0, 1, 8, 5}, cost: 71},
				{path: []int{0, 1, 8, 11}, cost: 82},
			},
		},
		{
			name: "dead_end",
			fields: fields{
				path: Path{
					path: []int{0, 9, 11},
					cost: 150,
				},
			},
			want: []Path{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := testRoadMap.explorePathsFrom(tt.fields.path)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("want: %v, got: %v", tt.want, got)
			}
		})
	}
}

func Test_explorePathFromTo(t *testing.T) {
	path := testRoadMap.explorePathFromTo(0, 10)
	if !reflect.DeepEqual(path, Path{path: []int{0, 1, 2, 5, 13, 12, 10}, cost: 70}) {
		t.Errorf("wrong path: %v", path)
	}
	path = testRoadMap.explorePathFromTo(0, 11)
	if !reflect.DeepEqual(path, Path{path: []int{0, 1, 2, 5, 13, 12, 11}, cost: 61}) {
		t.Errorf("wrong path: %v", path)
	}
	path = testRoadMap.explorePathFromTo(0, 4)
	if !reflect.DeepEqual(path, Path{path: []int{0, 1, 2, 3, 4}, cost: 36}) {
		t.Errorf("wrong path: %v", path)
	}
	path = testRoadMap.explorePathFromTo(0, 0)
	if !reflect.DeepEqual(path, Path{path: []int{0}, cost: 0}) {
		t.Errorf("wrong path: %v", path)
	}
	path = testRoadMap.explorePathFromTo(10, 0)
	if !reflect.DeepEqual(path, Path{path: nil, cost: 0}) {
		t.Errorf("wrong path: %v", path)
	}
}

func Benchmark_explorePathsFrom(b *testing.B) {
	var paths = make([]Path, 0, 16)
	for i := 0; i < b.N; i++ {
		if len(paths) == 0 {
			paths = testRoadMap.explorePathsFrom(Path{
				path: []int{0},
				cost: 0,
			})
			continue
		}
		var newPaths = make([]Path, 0, 16)
		for nn, path := range paths {
			if nn > 0 {
				i++
			}
			newPaths = append(newPaths, testRoadMap.explorePathsFrom(path)...)
			if i == b.N {
				break
			}
		}
		paths = newPaths
	}
}

func Benchmark_relaSlice_GetPath(b *testing.B) {
	for i := 0; i < b.N; i += 2 {
		if p := testRoadMap.explorePathFromTo(0, 11); p.cost != 61 {
			b.Errorf("expected 61, got: %0.2f", p.cost)
		}
		_ = testRoadMap.explorePathFromTo(rand.Intn(14), rand.Intn(14))
	}
}
