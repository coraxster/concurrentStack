package main

import (
	"fmt"
	"log"
	"sync"
)

func main() {
	l := NewStack()
	wg := sync.WaitGroup{}
	worker := func() {
		for i:=0; i<1000; i++ {
			l.Enqueue(i)
			j := l.Dequeue().(int)
			if j < 0 || j >= 1000 {
				log.Fatal("error")
			}
		}
		wg.Done()
	}
	wcorkerCount := 1000
	wg.Add(wcorkerCount)
	for c:=0;c<wcorkerCount;c++ {
		go worker()
	}
	wg.Wait()
	fmt.Println(l.Len())
}
