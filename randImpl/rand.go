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
	"unsafe"
)

type (
	MathRand   struct{}
	CryptoRand struct{}
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

type BufferedRand struct {
	init   sync.Once
	reader io.Reader

	mux    sync.RWMutex
	head   int64
	buffer [4096]byte

	refMb  unsafe.Pointer
	refSeq int64
}

func (r *BufferedRand) Generate() int64 {
	r.init.Do(func() {
		// 0.3 ns slower
		r.reader = bufio.NewReaderSize(crand.Reader, len(r.buffer)*2)
		r.fillBuffer()
	})
	for {
		waitSign := atomic.LoadInt64(&r.refSeq)
		headPtr := atomic.AddInt64(&r.head, -8)
		if headPtr < 0 {
			if headPtr == -8 {
				r.mux.Lock()
				r.fillBuffer()
				r.mux.Unlock()
				continue
			}
			waitChan := (*chan struct{})(atomic.LoadPointer(&r.refMb))
			if atomic.LoadInt64(&r.refSeq) != waitSign {
				continue
			}
			<-*waitChan
			continue
		}
		r.mux.RLock()
		ptri := int(headPtr) % (len(r.buffer) - 8)
		result := binary.LittleEndian.Uint64(r.buffer[ptri : ptri+8])
		r.mux.RUnlock()
		return int64(result)
	}
}

func (r *BufferedRand) fillBuffer() {
	if _, err := r.reader.Read(r.buffer[:]); err != nil {
		panic(err)
	}
	atomic.StoreInt64(&r.head, int64(len(r.buffer)))
	var newChan = make(chan struct{})
	atomic.AddInt64(&r.refSeq, 1)
	waitChan := (*chan struct{})(atomic.LoadPointer(&r.refMb))
	atomic.StorePointer(&r.refMb, unsafe.Pointer(&newChan))
	if waitChan != nil {
		close(*waitChan)
	}
}

type LightRand struct {
	b [8]byte
	o sync.Once
	r io.Reader
}

func (r *LightRand) Generate() int64 {
	r.o.Do(func() {
		r.r = bufio.NewReader(crand.Reader)
	})
	_, err := r.r.Read(r.b[:])
	if err != nil {
		panic(err)
	}
	return int64(binary.LittleEndian.Uint64(r.b[:]))
}
