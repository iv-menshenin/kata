package randImpl

import (
	"bufio"
	crand "crypto/rand"
	"encoding/binary"
	"io"
	"math"
	"math/big"
	mrand "math/rand"
	"sync"
	"sync/atomic"
)

type (
	MathRand     struct{}
	CryptoRand   struct{}
	BufferedRand struct {
		o sync.Once
		b io.Reader
		c [4096]byte
		a int64
		x sync.Mutex
	}
)

func (r *MathRand) Generate() int64 {
	return mrand.Int63()
}

func (r *CryptoRand) Generate() int64 {
	var bint big.Int
	bint.SetInt64(math.MaxInt64)
	result, err := crand.Int(crand.Reader, &bint)
	if err != nil {
		panic(err)
	}
	return result.Int64()
}

func (r *BufferedRand) Generate() int64 {
	init := func() {
		if _, err := r.b.Read(r.c[:]); err != nil {
			panic(err)
		}
		atomic.StoreInt64(&r.a, int64(len(r.c)))
	}
	r.o.Do(func() {
		// 0.3 ns slower
		r.b = bufio.NewReader(crand.Reader)
		init()
	})
	for {
		ptr := atomic.AddInt64(&r.a, -8)
		if ptr < 0 {
			if ptr == -8 {
				init()
			}
			continue
		}
		ptri := int(ptr) % (len(r.c) - 8)
		return int64(binary.LittleEndian.Uint64(r.c[ptri : ptri+8]))
	}
}
