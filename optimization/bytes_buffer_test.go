package optimization

import "testing"

/*
BenchmarkLength1-8      46648654                62.92 ns/op           64 B/op          1 allocs/op
BenchmarkLength2-8      167252322                7.023 ns/op           0 B/op          0 allocs/op
*/

func BenchmarkLength1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = length1()
	}
}
func BenchmarkLength2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = length2()
	}
}
