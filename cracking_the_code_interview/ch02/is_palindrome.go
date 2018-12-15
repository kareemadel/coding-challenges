package main

import (
	"fmt"
	"strings"
)

func main() {
	test("kareem")
	test("a")
	test("aba")
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
	reversed, _ := reverseList(l)
	return compareLists(reversed, l)
}

type node struct {
	Val  rune
	Next *node
}

func reverseList(l *node) (*node, *node) {
	if l == nil {
		return nil, nil
	}
	head, tail := reverseList(l.Next)
	newNode := &node{Val: l.Val}
	if tail == nil {
		head = newNode
	} else {
		tail.Next = newNode
	}
	tail = newNode
	return head, tail
}

func compareLists(l1, l2 *node) bool {
	for ; l1 != nil && l2 != nil; l1, l2 = l1.Next, l2.Next {
		if l1.Val != l2.Val {
			return false
		}
	}
	if l1 != nil || l2 != nil {
		return false
	}
	return true
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
