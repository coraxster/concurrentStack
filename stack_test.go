package main

import (
	"sync"
	"testing"
)

// BenchmarkEnqueue-4   	   10000	    174305 ns/op
func BenchmarkEnqueue(b *testing.B) {
	for n := 0; n < b.N; n++ {
		l := NewStack()
		for i := 0; i < n; i++ {
			l.Enqueue(i)
		}
	}
}

// BenchmarkDequeue-4   	  102180	    117269 ns/op
func BenchmarkDequeue(b *testing.B) {
	l := NewStack()
	for i := 0; i < b.N; i++ {
		l.Enqueue(i)
	}
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		for i := 0; i < n; i++ {
			l.Dequeue()
		}
	}
}

// BenchmarkEnqueueDequeue-4   	   10000	    296212 ns/op
func BenchmarkEnqueueDequeue(b *testing.B) {
	for n := 0; n < b.N; n++ {
		l := NewStack()
		for i := 0; i < n; i++ {
			l.Enqueue(i)
			l.Dequeue()
		}
	}
}

// BenchmarkConcurrentEnqueueDequeue-4   	   10000	    484614 ns/op
func BenchmarkConcurrentEnqueueDequeue(b *testing.B) {
	for n := 0; n < b.N; n++ {
		l := NewStack()
		wg := sync.WaitGroup{}
		worker := func(k int) {
			for i := 0; i < k; i++ {
				l.Enqueue(i)
				l.Dequeue()
			}
			wg.Done()
		}
		cf := 32
		wg.Add(cf)
		for i := 0; i < cf; i++ {
			go worker(n / cf)
		}
		wg.Wait()
	}
}

// BenchmarkSliceEnqueue-4   	   48784	    121959 ns/op
func BenchmarkSliceEnqueue(b *testing.B) {
	for n := 0; n < b.N; n++ {
		sl := make([]int, 0)
		m := sync.Mutex{}
		for i := 0; i < n; i++ {
			m.Lock()
			sl = append(sl, n)
			m.Unlock()
		}
	}
}

// BenchmarkSliceEnqueueDequeue-4   	   10000	    219312 ns/op
func BenchmarkSliceEnqueueDequeue(b *testing.B) {
	for n := 0; n < b.N; n++ {
		sl := make([]int, 0)
		m := sync.Mutex{}
		for i := 0; i < n; i++ {
			m.Lock()
			sl = append(sl, n)
			m.Unlock()
			m.Lock()
			if len(sl) > 0 {
				sl = sl[1:]
			}
			m.Unlock()
		}
	}
}

// BenchmarkSliceConcurrentEnqueueDequeue-4   	   10000	    587031 ns/op
func BenchmarkSliceConcurrentEnqueueDequeue(b *testing.B) {
	for n := 0; n < b.N; n++ {
		sl := make([]int, 0)
		m := sync.Mutex{} //RWMutex works slowly
		wg := sync.WaitGroup{}
		worker := func(k int) {
			for i := 0; i < k; i++ {
				m.Lock()
				sl = append(sl, n)
				m.Unlock()
				m.Lock()
				if len(sl) > 0 {
					sl = sl[1:]
				}
				m.Unlock()
			}
			wg.Done()
		}
		cf := 32
		wg.Add(cf)
		for i := 0; i < cf; i++ {
			go worker(n / cf)
		}
		wg.Wait()
	}
}
