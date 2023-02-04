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
	const (
		iter     = 1000000
		parallel = 1000
	)
	var (
		bucket [250]int64
		wg     sync.WaitGroup
		ch     = make(chan struct{}, parallel)
	)
	wg.Add(iter)
	for n := 0; n < iter; n++ {
		ch <- struct{}{}
		go func() {
			got := r.Generate()
			ptr := int(uint64(got) % uint64(len(bucket)))
			atomic.AddInt64(&bucket[ptr], 1)
			wg.Done()
			<-ch
		}()
	}
	wg.Wait()
	stat := countStat(bucket[:])
	if diff := stat.max - stat.min; diff > 500 || diff < 250 {
		t.Errorf("entropy is too bad (%d:%d %d)", stat.min, stat.max, diff)
	}
	if stat.all != iter {
		t.Errorf("%d iterations, but expected is %d", stat.all, iter)
	}
}

func countStat(buckets []int64) (result stat) {
	for _, b := range buckets {
		if result.min == 0 || result.min > b {
			result.min = b
		}
		if result.max < b {
			result.max = b
		}
		result.all += b
	}
	return
}

type stat struct {
	min, max, all int64
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

func BenchmarkUnsafeRand_Generate(b *testing.B) {
	b.ReportAllocs()
	var rnd LightRand
	for n := 0; n < b.N; n++ {
		_ = rnd.Generate()
	}
}

func benchGenerationConcurrent(b *testing.B, r randomizer) {
	b.ReportAllocs()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_ = r.Generate()
		}
	})
}
