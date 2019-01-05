package main

import (
	"fmt"
	"strings"
)

func main() {
	test("kareem")
	test("a")
	test("aba")
	test("abbaaa")
	test("abba")
	test("")
}

func test(s string) {
	l := buildList(s)
	printList(l)
	fmt.Println(isPalindrome(l))
}

func isPalindrome(n *node) bool {
	if n == nil {
		return false
	}
	l := listLen(n)
	_, status := isPalindromeHelper(n, l)
	return status
}

func isPalindromeHelper(n *node, l int) (*node, bool) {
	if l == 0 {
		return n, true
	} else if l == 1 {
		return n.Next, true
	}
	tail, status := isPalindromeHelper(n.Next, l-2)
	if !status || tail.Val != n.Val {
		return nil, false
	}
	return tail.Next, true
}

func listLen(n *node) int {
	counter := 0
	for n != nil {
		counter++
		n = n.Next
	}
	return counter
}

type node struct {
	Val  rune
	Next *node
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
