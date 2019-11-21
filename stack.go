package main

import (
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
		newEl.next = nil
		hP := s.head
		if hP == nil && atomic.CompareAndSwapPointer(&s.head, nil, newElP) {
			return
		}
		newEl.next = hP
		if atomic.CompareAndSwapPointer(&s.head, hP, newElP) {
			return
		}
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
