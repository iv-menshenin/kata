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
		from int
	}
	tests := []struct {
		name   string
		fields fields
		want   Exploration
	}{
		{
			name: "from_zero",
			fields: fields{
				from: 0,
			},
			want: Exploration{
				{path: nil, cost: 0},
				{path: []int{1}, cost: 8},
				{path: nil, cost: 0},
				{path: nil, cost: 0},
				{path: nil, cost: 0},
				{path: nil, cost: 0},
				{path: []int{6}, cost: 5},
				{path: nil, cost: 0},
				{path: nil, cost: 0},
				{path: []int{9}, cost: 132},
				{path: nil, cost: 0},
				{path: nil, cost: 0},
				{path: nil, cost: 0},
				{path: nil, cost: 0},
			},
		},
		{
			name: "dead_end",
			fields: fields{
				from: 11,
			},
			want: make(Exploration, 14),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var e = make(Exploration, 14)
			e.explorePathsFrom(testRoadMap, tt.fields.from)
			if !reflect.DeepEqual(e, tt.want) {
				t.Errorf("want: %v, got: %v", tt.want, e)
			}
		})
	}
}

func Test_explorePathFromTo(t *testing.T) {
	path := testRoadMap.explorePathFromTo(0, 10)
	if !reflect.DeepEqual(path, Path{path: []int{0, 1, 2, 5, 13, 12, 10}, cost: 70, explored: true}) {
		t.Errorf("wrong path: %v", path)
	}
	path = testRoadMap.explorePathFromTo(0, 11)
	if !reflect.DeepEqual(path, Path{path: []int{0, 1, 2, 5, 13, 12, 11}, cost: 61, explored: true}) {
		t.Errorf("wrong path: %v", path)
	}
	path = testRoadMap.explorePathFromTo(0, 4)
	if !reflect.DeepEqual(path, Path{path: []int{0, 1, 2, 3, 4}, cost: 36, explored: true}) {
		t.Errorf("wrong path: %v", path)
	}
	path = testRoadMap.explorePathFromTo(0, 0)
	if !reflect.DeepEqual(path, Path{path: []int{0}, cost: 0, explored: true}) {
		t.Errorf("wrong path: %v", path)
	}
	path = testRoadMap.explorePathFromTo(10, 0)
	if !reflect.DeepEqual(path, Path{path: []int{}, cost: 0, explored: false}) {
		t.Errorf("wrong path: %v", path)
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

func Benchmark_relaSlice_GetPath_2(b *testing.B) {
	var rm = make(RoadMap, maxNodes*10)
	for i := 0; i < len(rm); i++ {
		rm[i] = PathWeight{
			from:   rand.Intn(maxNodes),
			to:     rand.Intn(maxNodes),
			weight: rand.Float64() * 12,
		}
	}
	rm[maxNodes*10-1] = PathWeight{
		from:   maxNodes - 1,
		to:     0,
		weight: 99,
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = rm.explorePathFromTo(rand.Intn(maxNodes), rand.Intn(maxNodes))
	}
}
