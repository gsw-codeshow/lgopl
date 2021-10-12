package benchmark_test

import (
	"lgopl/ch9/benchmark"
	"sync"
	"testing"
	"time"
)

func Work() {
	time.Sleep(10 * time.Millisecond)
	return
}

func BenchmarkChan(b *testing.B) {
	b.ResetTimer()
	waitGroup := benchmark.NewWaitGroup()
	for i := 0; i < b.N; i++ {
		waitGroup.Add(1)
		go func() {
			defer waitGroup.Done()
			Work()
		}()
	}
	waitGroup.Wait()
	return
}

func BenchmarkWaitGroup(b *testing.B) {
	b.ResetTimer()
	var waitGroup sync.WaitGroup
	for i := 0; i < b.N; i++ {
		waitGroup.Add(1)
		go func() {
			defer waitGroup.Done()
			Work()
		}()
	}
	waitGroup.Wait()
	return
}
