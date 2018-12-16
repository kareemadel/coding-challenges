package main

import (
	"sync"
)

func main() {

}

type queue struct {
	lock     sync.Mutex
	oldStack stack
	newStack stack
}

func (q *queue) enqueue(t item) {
	q.lock.Lock()
	defer q.lock.Unlock()

	q.newStack.push(t)
}

func (q *queue) dequeue() (item, bool) {
	q.lock.Lock()
	defer q.lock.Unlock()

	if q.oldStack.isEmpty() {
		q.shift()
	}
	return q.oldStack.pop()
}

func (q queue) peek() (item, bool) {
	if !q.oldStack.isEmpty() {
		return q.oldStack.peek()
	} else if !q.newStack.isEmpty() {
		q.shift()
		return q.oldStack.peek()
	} else {
		return item{}, false
	}
}

func (q *queue) shift() {
	q.lock.Lock()
	defer q.lock.Unlock()

	if q.oldStack.isEmpty() {
		for !q.newStack.isEmpty() {
			popped, _ := q.newStack.pop()
			q.oldStack.push(popped)
		}
	}
}

func (q queue) isEmpty() bool {
	return q.oldStack.isEmpty() && q.newStack.isEmpty()
}

type stack struct {
	lock sync.Mutex
	data []item
	size int
}

func (s *stack) push(t item) {
	s.lock.Lock()
	defer s.lock.Unlock()

	s.data = append(s.data, t)
	s.size++
}

func (s *stack) pop() (item, bool) {
	s.lock.Lock()
	defer s.lock.Unlock()

	if s.size == 0 {
		return item{}, false
	}
	t := s.data[s.size-1]
	s.data = s.data[:s.size-1]
	s.size--
	return t, true
}

func (s stack) peek() (item, bool) {
	s.lock.Lock()
	defer s.lock.Unlock()

	if s.size == 0 {
		return item{}, false
	}
	return s.data[s.size-1], true
}

func (s stack) isEmpty() bool {
	return s.size == 0
}

type item struct {
	Val int
}
