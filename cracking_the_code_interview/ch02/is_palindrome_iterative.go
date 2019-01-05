package main

import (
	"fmt"
	"strings"
	"sync"
)

func main() {
	test("kareem")
	test("a")
	test("aba")
	test("abca")
	test("abba")
	test("abbaaa")
	test("")
}

func test(s string) {
	l := buildList(s)
	printList(l)
	fmt.Println(isPalindrome(l))
}

func isPalindrome(l *node) bool {
	if l == nil {
		return false
	}
	slow, fast := l, l
	var runeStack stack
	runeStack.push(slow.Val)
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		runeStack.push(slow.Val)
	}
	if fast == nil {
		runeStack.pop()
	}
	for slow != nil {
		val, ok := runeStack.pop()
		if !ok || val != slow.Val {
			return false
		}
		slow = slow.Next
	}
	return runeStack.isEmpty()
}

type node struct {
	Val  rune
	Next *node
}

type stack struct {
	lock sync.Mutex
	s    []rune
}

func (s *stack) push(v rune) {
	s.lock.Lock()
	defer s.lock.Unlock()

	s.s = append(s.s, v)
}

func (s *stack) pop() (rune, bool) {
	s.lock.Lock()
	defer s.lock.Unlock()

	l := len(s.s)
	if l == 0 {
		return 0, false
	}

	res := s.s[l-1]
	s.s = s.s[:l-1]
	return res, true
}

func (s stack) isEmpty() bool {
	return len(s.s) == 0
}

func buildList(s string) *node {
	r := []rune(s)
	length := len(r)
	if length == 0 {
		return nil
	}
	var head *node
	for i := length - 1; i >= 0; i-- {
		head = &node{r[i], head}
	}
	return head
}

func printList(l *node) {
	if l == nil {
		fmt.Println()
		return
	}
	var output strings.Builder
	for l != nil {
		output.WriteString(fmt.Sprintf("%c->", l.Val))
		l = l.Next
	}
	fmt.Println(output.String()[:output.Len()-2])
}
