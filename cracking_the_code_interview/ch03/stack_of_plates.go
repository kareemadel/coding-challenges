package main

import (
	"sync"
)

func main() {

}

type plateStack struct {
	lock     sync.Mutex
	stacks   []stack
	size     int
	capacity int
}

func (p *plateStack) push(t item) {
	p.lock.Lock()
	defer p.lock.Unlock()

	p.appendStack()
	last, _ := p.getLastStack()
	last.push(t)

}

func (p *plateStack) pop() (item, bool) {
	p.lock.Lock()
	defer p.lock.Unlock()

	p.removeLastStack()
	last, ok := p.getLastStack()
	if !ok {
		return item{}, false
	}
	return last.pop()
}

func (p *plateStack) appendStack() {
	p.lock.Lock()
	defer p.lock.Unlock()

	lastStack, ok := p.getLastStack()
	if !ok || lastStack.isFull() {
		newStack := stack{}
		newStack.init(p.capacity)
		p.stacks = append(p.stacks, newStack)
		p.size++
	}
}

func (p *plateStack) removeLastStack() {
	last, ok := p.getLastStack()
	if ok && last.isEmpty() {
		p.stacks = p.stacks[:p.size-1]
		p.size--
	}
}

func (p plateStack) getLastStack() (stack, bool) {
	if p.size > 0 {
		return p.stacks[p.size-1], true
	}
	return stack{}, false
}

type stack struct {
	lock     sync.Mutex
	data     []item
	size     int
	capacity int
}

func (s *stack) push(t item) bool {
	s.lock.Lock()
	defer s.lock.Unlock()

	if !s.isFull() {
		s.data[s.size] = t
		s.size++
		return true
	}
	return false
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

func (s stack) isFull() bool {
	return s.size == s.capacity
}

func (s *stack) init(capacity int) {
	s.lock.Lock()
	defer s.lock.Unlock()

	s.size = 0
	s.capacity = capacity
	s.data = make([]item, capacity)

}

type item struct {
	Val int
}
