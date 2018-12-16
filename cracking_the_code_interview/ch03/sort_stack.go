package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	test(100, 10)
}

func test(max, len int) {
	s := stack{}
	for i := 0; i < len; i++ {
		s.push(item{rand.Intn(max)})
	}
	fmt.Println(s.data)
	sortStack(s)
	fmt.Println(s.data)
}

func sortStack(s stack) {
	buffer := stack{}
	for !s.isEmpty() {
		temp := s.pop()
		for !buffer.isEmpty() && buffer.peek().Val > temp.Val {
			s.push(buffer.pop())
		}
		buffer.push(temp)
	}
	for !buffer.isEmpty() {
		s.push(buffer.pop())
	}
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

func (s *stack) pop() item {
	s.lock.Lock()
	defer s.lock.Unlock()

	if s.size == 0 {
		panic("stack is empty")
	}
	t := s.data[s.size-1]
	s.data = s.data[:s.size-1]
	s.size--
	return t
}

func (s stack) peek() item {
	s.lock.Lock()
	defer s.lock.Unlock()

	if s.size == 0 {
		panic("stack is empty")
	}
	return s.data[s.size-1]
}

func (s stack) isEmpty() bool {
	return s.size == 0
}

type item struct {
	Val int
}
