package bracketSeria

import (
	"testing"
)

type funcTest func(int) []string

func testMakeBracketSeria(t *testing.T, fn funcTest) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "0",
			args: args{n: 0},
			want: nil,
		},
		{
			name: "1",
			args: args{n: 1},
			want: []string{"()"},
		},
		{
			name: "2",
			args: args{n: 2},
			want: []string{"(())", "()()"},
		},
		{
			name: "3",
			args: args{n: 3},
			want: []string{"()()()", "()(())", "(())()", "((()))", "(()())"},
		},
		{
			name: "4",
			args: args{n: 4},
			want: []string{"(())(())", "()((()))", "()()(())", "()()()()", "()(())()", "()(()())", "(()(()))", "(()()())", "(()())()", "(())()()", "((())())", "((()))()", "((()()))", "(((())))"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := fn(tt.args.n)
			var m = map[string]int{}
			for _, g := range got {
				m[g]++
			}
			for _, w := range tt.want {
				if m[w]--; m[w] == 0 {
					delete(m, w)
				}
			}
			if len(m) != 0 {
				t.Errorf("matching error, got: %v, want: %v", got, tt.want)
			}
		})
	}
}

func Test_makeBracketRecursive(t *testing.T) {
	testMakeBracketSeria(t, makeBracketRecursive)
}

func Test_makeBracketCycle(t *testing.T) {
	testMakeBracketSeria(t, makeBracketCycle)
}

func benchMakeBracketSeria(b *testing.B, fn funcTest) {
	for i := 0; i < b.N; i++ {
		for n := 4; n < 6; n++ {
			_ = fn(n)
		}
	}
}

func Benchmark_makeBracketRecursive(t *testing.B) {
	benchMakeBracketSeria(t, makeBracketRecursive)
}

func Benchmark_makeBracketCycle(t *testing.B) {
	benchMakeBracketSeria(t, makeBracketCycle)
}
