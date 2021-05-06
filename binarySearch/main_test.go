package main

import (
	"fmt"
	"math/rand"
	"reflect"
	"sort"
	"strconv"
	"testing"
	"time"
)

const (
	testCount  = 1000
	itemsCount = 100000
)

func Test_entities_Mid(t *testing.T) {
	type args struct {
		e entity
	}
	tests := []struct {
		name string
		c    entities
		args args
		want int
	}{
		{
			name: "zero arr",
			c:    nil,
			args: args{e: entity{someData: "a"}},
			want: 0,
		},
		{
			name: "6 els",
			c: entities{
				{someData: "a"},
				{someData: "b"},
				{someData: "c"},
				{someData: "d"},
				{someData: "e"},
				{someData: "g"},
			},
			args: args{e: entity{someData: "f"}},
			want: 5,
		},
		{
			name: "first",
			c: entities{
				{someData: "a"},
				{someData: "b"},
				{someData: "c"},
				{someData: "d"},
				{someData: "e"},
				{someData: "f"},
				{someData: "g"},
			},
			args: args{e: entity{someData: "a"}},
			want: 0,
		},
		{
			name: "first 2",
			c: entities{
				{someData: "a"},
				{someData: "b"},
				{someData: "c"},
				{someData: "d"},
				{someData: "e"},
				{someData: "f"},
				{someData: "g"},
			},
			args: args{e: entity{someData: "0"}},
			want: 0,
		},
		{
			name: "last",
			c: entities{
				{someData: "a"},
				{someData: "b"},
				{someData: "c"},
				{someData: "d"},
				{someData: "e"},
				{someData: "f"},
				{someData: "g"},
			},
			args: args{e: entity{someData: "g"}},
			want: 6,
		},
		{
			name: "last 2",
			c: entities{
				{someData: "a"},
				{someData: "b"},
				{someData: "c"},
				{someData: "d"},
				{someData: "e"},
				{someData: "f"},
			},
			args: args{e: entity{someData: "g"}},
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.Mid(tt.args.e); got != tt.want {
				t.Errorf("Mid() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_sortedEntities_Add(t *testing.T) {
	type fields struct {
		entities entities
	}
	type args struct {
		e entity
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   sortedEntities
	}{
		{
			name: "add 1",
			fields: fields{entities: entities{
				{someData: "a"},
				{someData: "b"},
				{someData: "c"},
				{someData: "d"},
				{someData: "e"},
			}},
			args: args{e: entity{someData: "d"}},
			want: sortedEntities{entities: entities{
				{someData: "a"},
				{someData: "b"},
				{someData: "c"},
				{someData: "d"},
				{someData: "d"},
				{someData: "e"},
			}},
		},
		{
			name: "add last",
			fields: fields{entities: entities{
				{someData: "a"},
				{someData: "b"},
				{someData: "c"},
				{someData: "d"},
				{someData: "e"},
			}},
			args: args{e: entity{someData: "f"}},
			want: sortedEntities{entities: entities{
				{someData: "a"},
				{someData: "b"},
				{someData: "c"},
				{someData: "d"},
				{someData: "e"},
				{someData: "f"},
			}},
		},
		{
			name: "add first 1",
			fields: fields{entities: entities{
				{someData: "a"},
				{someData: "b"},
				{someData: "c"},
				{someData: "d"},
				{someData: "e"},
			}},
			args: args{e: entity{someData: "a"}},
			want: sortedEntities{entities: entities{
				{someData: "a"},
				{someData: "a"},
				{someData: "b"},
				{someData: "c"},
				{someData: "d"},
				{someData: "e"},
			}},
		},
		{
			name: "add first 2",
			fields: fields{entities: entities{
				{someData: "b"},
				{someData: "c"},
				{someData: "d"},
				{someData: "e"},
			}},
			args: args{e: entity{someData: "a"}},
			want: sortedEntities{entities: entities{
				{someData: "a"},
				{someData: "b"},
				{someData: "c"},
				{someData: "d"},
				{someData: "e"},
			}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := sortedEntities{
				entities: tt.fields.entities,
			}
			c.Add(tt.args.e)
			if !reflect.DeepEqual(c, tt.want) {
				t.Errorf("matching error\nwant: %+v\ngot: %+v", tt.want, c)
			}
		})
	}
}

func Test_sortedEntities_Find(t *testing.T) {
	type testR struct {
		someData string
		want     int
	}
	var (
		collection = sortedEntities{}
		tests      = make([]testR, 0, 10)
	)
	for nn := 0; nn < 1000; nn++ {
		collection.Add(entity{someData: strconv.FormatInt(int64(nn), 16) + "_" + strconv.FormatInt(rand.Int63(), 16)})
	}
	for nn := 0; nn < cap(tests); nn++ {
		r := rand.Intn(len(collection.entities))
		tests = append(tests, testR{
			someData: collection.entities[r].someData,
			want:     r,
		})
	}
	for _, tt := range tests {
		t.Run(tt.someData, func(t *testing.T) {
			if got := collection.Find(tt.someData); got != tt.want {
				t.Errorf("Find() = %v, want %v", got, tt.want)
			}
		})
	}
	// fmt.Printf("%+v", collection.entities[180:190])
}

func goTestFind(t *testing.T, findFn func([]int, int) int) {
	type args struct {
		a []int
		f int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "midden 1",
			args: args{
				a: []int{0, 1, 2, 3, 5, 7, 11, 13, 17, 21},
				f: 7,
			},
			want: 5,
		},
		{
			name: "midden 2",
			args: args{
				a: []int{0, 1, 2, 3, 5, 7, 11, 13, 17, 21},
				f: 17,
			},
			want: 8,
		},
		{
			name: "midden 3",
			args: args{
				a: []int{0, 1, 2, 3, 5, 7, 11, 13, 17, 21},
				f: 2,
			},
			want: 2,
		},
		{
			name: "first 1",
			args: args{
				a: []int{0, 1, 2, 3, 5, 7, 11, 13, 17, 21},
				f: 0,
			},
			want: 0,
		},
		{
			name: "first 2",
			args: args{
				a: []int{0, 1, 2, 3, 5, 7, 11, 13, 17},
				f: 0,
			},
			want: 0,
		},
		{
			name: "first 3",
			args: args{
				a: []int{0, 1, 2, 3, 5, 7, 11, 13},
				f: 0,
			},
			want: 0,
		},
		{
			name: "last 1",
			args: args{
				a: []int{0, 1, 2, 3, 5, 7, 11, 13, 17, 21},
				f: 21,
			},
			want: 9,
		},
		{
			name: "last 2",
			args: args{
				a: []int{0, 1, 2, 3, 5, 7, 11, 13, 17},
				f: 17,
			},
			want: 8,
		},
		{
			name: "last 3",
			args: args{
				a: []int{0, 1, 2, 3, 5, 7, 11, 13},
				f: 13,
			},
			want: 7,
		},
		{
			name: "nil array",
			args: args{
				a: nil,
				f: 13,
			},
			want: -1,
		},
		{
			name: "empty array",
			args: args{
				a: []int{},
				f: 13,
			},
			want: -1,
		},
		{
			name: "short array 1",
			args: args{
				a: []int{1},
				f: 1,
			},
			want: 0,
		},
		{
			name: "short array 2",
			args: args{
				a: []int{1, 2},
				f: 1,
			},
			want: 0,
		},
		{
			name: "short array 3",
			args: args{
				a: []int{1, 2},
				f: 2,
			},
			want: 1,
		},
		{
			name: "short array 4",
			args: args{
				a: []int{1, 2, 3},
				f: 3,
			},
			want: 2,
		},
		{
			name: "not found 1",
			args: args{
				a: []int{1, 2, 3},
				f: 6,
			},
			want: -1,
		},
		{
			name: "not found 2",
			args: args{
				a: []int{1, 2, 4, 5},
				f: 3,
			},
			want: -1,
		},
		{
			name: "not found 3",
			args: args{
				a: []int{1},
				f: 6,
			},
			want: -1,
		},
		{
			name: "not found 4",
			args: args{
				a: []int{1, 10, 100, 1000},
				f: 60,
			},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findFn(tt.args.a, tt.args.f); got != tt.want {
				t.Errorf("find() = %v, want %v", got, tt.want)
			}
		})
	}
}

func getArr(a []int, i int) interface{} {
	if i > 0 && len(a) > i {
		return a[i]
	}
	return "unknown"
}

func goTestBrutForce(t *testing.T, findFn func([]int, int, ...*uint32) int) {
	var (
		a      = make([]int, 0, itemsCount)
		tests1 = make(map[int]int)
		tests2 = make([]int, 0, testCount)
	)
	for i := 0; i < cap(a); i++ {
		a = append(a, rand.Int())
	}
	sort.Ints(a)
	for i := 0; i < testCount/2; i++ {
		idx := rand.Intn(cap(a))
		tests1[i] = idx
		nf := a[idx] + 1
		for {
			if idx > len(a)-2 || a[idx+1] > nf {
				break
			}
			nf++
		}
		tests2 = append(tests2, nf)
	}
	var iterCount uint32
	for i, v := range tests1 {
		if got := findFn(a, a[v], &iterCount); got < 0 || a[got] != a[v] {
			t.Errorf("find() = %v, want %v", got, v)
		}
		if got := findFn(a, tests2[i], &iterCount); got > -1 {
			t.Errorf("find() = %v, want -1", got)
		}
	}
	println(fmt.Sprintf("average ticks: %d\ncomplexity: %s", iterCount/testCount, testComplexity(itemsCount, int(iterCount/testCount))))
}

func goTestSimilarValues(t *testing.T, findFn func([]int, int, ...*uint32) int) {
	var (
		a     = make([]int, 0, itemsCount)
		tests = make(map[int]int)
	)
	for i := 0; i < cap(a); i++ {
		a = append(a, rand.Intn(100))
	}
	sort.Ints(a)
	for i := 0; i < testCount; i++ {
		idx := rand.Intn(cap(a))
		tests[i] = idx
	}
	var iterCount uint32
	for _, v := range tests {
		if got := findFn(a, a[v], &iterCount); got < 0 || a[got] != a[v] {
			t.Errorf("find() = %v, want %v", got, v)
		}
	}
	println(fmt.Sprintf("average ticks: %d\ncomplexity: %s", iterCount/testCount, testComplexity(itemsCount, int(iterCount/testCount))))
}

func goTestEntropic(t *testing.T, findFn func([]int, int, ...*uint32) int) {
	var (
		a      = make([]int, 0, itemsCount)
		tests  = make(map[int]int)
		offset = rand.Int()
	)
	for i := 0; i < cap(a)/3; i++ {
		a = append(a, rand.Intn(1000))
	}
	for i := 0; i < cap(a)/3; i++ {
		a = append(a, rand.Intn(1000)+offset)
	}
	offset = rand.Int()
	for i := len(a); i < cap(a); i++ {
		a = append(a, rand.Intn(1000)+offset)
	}
	for i := 0; i < testCount; i++ {
		tests[i] = rand.Int()
		a = append(a, tests[i])
	}
	sort.Ints(a)
	var iterCount uint32
	for _, v := range tests {
		if got := findFn(a, v, &iterCount); got < 0 || a[got] != v {
			t.Errorf("find() = %v, want %v", getArr(a, got), v)
		}
	}
	println(fmt.Sprintf("average ticks: %d\ncomplexity: %s", iterCount/testCount, testComplexity(itemsCount, int(iterCount/testCount))))
}

func goBenchmarkFind(b *testing.B, findFn func([]int, int) int) {
	var (
		a     = make([]int, 0, 1000)
		tests = make(map[int]int, b.N)
	)
	for i := 0; i < cap(a); i++ {
		a = append(a, rand.Int())
	}
	sort.Ints(a)
	for i := 0; i < b.N; i++ {
		tests[i] = rand.Intn(len(a))
	}
	b.ResetTimer()
	for _, v := range tests {
		if got := findFn(a, a[v]); got != v {
			b.Errorf("find() = %v, want %v", got, v)
		}
	}
}

func init() {
	rand.Seed(time.Now().UnixNano())
}
