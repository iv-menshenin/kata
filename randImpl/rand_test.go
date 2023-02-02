package randImpl

import (
	"sync"
	"sync/atomic"
	"testing"
)

type (
	randomizer interface {
		Generate() int64
	}
)

func TestMathRand_Generate(t *testing.T) {
	t.Parallel()
	checkGeneration(t, &MathRand{})
}

func TestCryptoRand_Generate(t *testing.T) {
	t.Parallel()
	checkGeneration(t, &CryptoRand{})
}

func TestBufferedRand_Generate(t *testing.T) {
	t.Parallel()
	checkGeneration(t, &BufferedRand{})
}

func checkGeneration(t *testing.T, r randomizer) {
	const iter = 1000000
	var (
		bucket [256]int64
		wg     sync.WaitGroup
	)
	wg.Add(iter)
	for n := 0; n < iter; n++ {
		go func() {
			ptr := int(uint64(r.Generate()) % uint64(len(bucket)))
			atomic.AddInt64(&bucket[ptr], 1)
			wg.Done()
		}()
	}
	wg.Wait()
	var min, max, all int64
	for _, b := range bucket {
		if min == 0 || min > b {
			min = b
		}
		if max < b {
			max = b
		}
		all += b
	}
	if diff := max - min; diff > 500 || diff < 250 {
		t.Errorf("entropy is too low (%d:%d %d)", min, max, diff)
	}
	if all != iter {
		t.Errorf("%d iterations, but expected is %d", all, iter)
	}
}

func BenchmarkMathRand_Generate(b *testing.B) {
	benchGeneration(b, &MathRand{})
}

func BenchmarkCryptoRand_Generate(b *testing.B) {
	benchGeneration(b, &CryptoRand{})
}

func BenchmarkBufferedRand_Generate(b *testing.B) {
	benchGeneration(b, &BufferedRand{})
}

func benchGeneration(b *testing.B, r randomizer) {
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		_ = r.Generate()
	}
}

func BenchmarkMathRand_Generate_Concurrent(b *testing.B) {
	benchGenerationConcurrent(b, &MathRand{})
}

func BenchmarkCryptoRand_Generate_Concurrent(b *testing.B) {
	benchGenerationConcurrent(b, &CryptoRand{})
}

func BenchmarkBufferedRand_Generate_Concurrent(b *testing.B) {
	benchGenerationConcurrent(b, &BufferedRand{})
}

func benchGenerationConcurrent(b *testing.B, r randomizer) {
	b.ReportAllocs()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_ = r.Generate()
		}
	})
}
