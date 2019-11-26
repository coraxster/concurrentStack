package main

import (
	"runtime"
	"sync/atomic"
	"unsafe"
)

type Stack struct {
	head unsafe.Pointer
}

type Element struct {
	val  int
	next unsafe.Pointer
}

func NewStack() *Stack {
	return &Stack{}
}

func (s *Stack) Enqueue(val int) {
	newEl := &Element{
		val: val,
	}
	newElP := unsafe.Pointer(newEl)
	for {
		h := s.head
		newEl.next = h
		if atomic.CompareAndSwapPointer(&s.head, h, newElP) {
			return
		}
		runtime.Gosched()
	}
}

func (s *Stack) Dequeue() interface{} {
	for {
		hP := s.head
		if hP == nil {
			return nil
		}
		h := (*Element)(hP)
		if atomic.CompareAndSwapPointer(&s.head, hP, h.next) {
			return h.val
		}
		runtime.Gosched()
	}
}

func (s *Stack) Len() int {
	c := 0
	el := (*Element)(s.head)
	for {
		if el == nil {
			return c
		}
		c++
		el = (*Element)(el.next)
	}
}
