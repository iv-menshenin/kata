package dijkstra

import (
	"math/rand"
	"reflect"
	"testing"
)

const zero = 0

var testData = [][]int{
	{zero, 8, zero, zero, zero, zero, 5, zero, zero, 132, zero, zero, zero, zero},
	{zero, zero, 13, zero, zero, zero, zero, zero, 9, zero, zero, zero, zero, zero},
	{zero, zero, zero, 3, zero, 21, zero, zero, zero, zero, zero, zero, zero, zero},
	{zero, zero, zero, zero, 12, zero, zero, zero, zero, zero, zero, zero, zero, zero},
	{zero, zero, zero, 12, zero, 21, zero, zero, zero, zero, zero, zero, zero, zero},
	{zero, zero, zero, zero, zero, zero, zero, zero, zero, zero, zero, zero, zero, 9},
	{zero, zero, zero, zero, zero, zero, zero, 35, zero, zero, zero, zero, zero, zero},
	{zero, zero, zero, zero, zero, zero, zero, zero, zero, zero, zero, 100, zero, zero},
	{zero, zero, zero, zero, zero, 54, zero, zero, zero, zero, zero, 65, zero, zero},
	{zero, zero, zero, zero, zero, zero, zero, zero, zero, zero, zero, 18, zero, zero},
	{zero, zero, zero, zero, zero, zero, zero, zero, zero, zero, zero, zero, zero, zero},
	{zero, zero, zero, zero, zero, zero, zero, zero, zero, zero, zero, zero, zero, zero},
	{zero, zero, zero, zero, zero, zero, zero, zero, zero, zero, 12, 3, zero, zero},
	{zero, zero, zero, zero, 45, zero, zero, zero, zero, zero, zero, zero, 7, zero},
}

func TestPathFinder_GetPath(t *testing.T) {
	var pFinder = New(testData)
	t.Run("0_10", func(t *testing.T) {
		if cost, path := pFinder.GetPath(0, 10); cost != 70 {
			t.Errorf("expected 70, got: %d", cost)
		} else if !reflect.DeepEqual(path, []int{0, 1, 2, 5, 13, 12, 10}) {
			t.Errorf("wrong path: %v", path)
		}
	})
	t.Run("0_11", func(t *testing.T) {
		if cost, path := pFinder.GetPath(0, 11); cost != 61 {
			t.Errorf("expected 62, got: %d", cost)
		} else if !reflect.DeepEqual(path, []int{0, 1, 2, 5, 13, 12, 11}) {
			t.Errorf("wrong path: %v", path)
		}
	})
	t.Run("0_4", func(t *testing.T) {
		if cost, path := pFinder.GetPath(0, 4); cost != 36 {
			t.Errorf("expected 36, got: %d", cost)
		} else if !reflect.DeepEqual(path, []int{0, 1, 2, 3, 4}) {
			t.Errorf("wrong path: %v", path)
		}
	})
	t.Run("0_0", func(t *testing.T) {
		if cost, _ := pFinder.GetPath(0, 0); cost != infinity {
			t.Errorf("expected infinity, got: %d", cost)
		}
	})
	t.Run("11_8", func(t *testing.T) {
		if cost, _ := pFinder.GetPath(11, 8); cost != infinity {
			t.Errorf("expected infinity, got: %d", cost)
		}
	})
	t.Run("0_18", func(t *testing.T) {
		if cost, _ := pFinder.GetPath(0, 18); cost != infinity {
			t.Errorf("expected infinity, got: %d", cost)
		}
	})
}

func Benchmark_PathFinder_GetPath(b *testing.B) {
	var pFinder = New(testData)
	b.ResetTimer()
	for i := 0; i < b.N; i += 2 {
		if cost, _ := pFinder.GetPath(0, 11); cost != 61 {
			b.Errorf("expected 62, got: %d", cost)
		}
		if cost, _ := pFinder.GetPath(0, 10); cost != 70 {
			b.Errorf("expected 70, got: %d", cost)
		}
	}
}

func Benchmark_PathFinder_GetPath_2(b *testing.B) {
	const dim = 1024
	var geo = make([][]int, dim)
	for i := 0; i < len(geo); i++ {
		geo[i] = make([]int, dim)
	}
	for i := 0; i < 1000000; i++ {
		geo[rand.Intn(dim)][rand.Intn(dim)] = rand.Intn(100)
	}
	b.ResetTimer()
	var pFinder = New(geo)
	for i := 0; i < b.N; i++ {
		_, _ = pFinder.GetPath(rand.Intn(dim), rand.Intn(dim))
	}
}

func TestPathFinder_nextNode(t *testing.T) {
	type fields struct {
		explored []bool
	}
	type args struct {
		exploration []int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want1  int
		want2  bool
	}{
		{
			name:   "first",
			fields: fields{explored: []bool{false, false, false, false}},
			args:   args{exploration: []int{1, 2, 3, 4}},
			want1:  0,
			want2:  true,
		},
		{
			name:   "last",
			fields: fields{explored: []bool{false, false, false, false}},
			args:   args{exploration: []int{4, 3, 2, 1}},
			want1:  3,
			want2:  true,
		},
		{
			name:   "explored_1",
			fields: fields{explored: []bool{false, false, false, true}},
			args:   args{exploration: []int{4, 3, 2, 1}},
			want1:  2,
			want2:  true,
		},
		{
			name:   "explored_all",
			fields: fields{explored: []bool{true, true, true, true}},
			args:   args{exploration: []int{4, 3, 2, 1}},
			want1:  -1,
			want2:  false,
		},
		{
			name:   "empty_slice",
			fields: fields{explored: []bool{}},
			args:   args{exploration: []int{}},
			want1:  -1,
			want2:  false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &matrixExplorer{
				explored:        tt.fields.explored,
				exploredWeights: tt.args.exploration,
			}
			got1, got2 := p.nextNode()
			if got1 != tt.want1 {
				t.Errorf("nextNode() got = %v, want %v", got1, tt.want1)
			}
			if got2 != tt.want2 {
				t.Errorf("nextNode() got2 = %v, want %v", got2, tt.want2)
			}
		})
	}
}
