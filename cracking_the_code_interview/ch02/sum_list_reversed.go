package main

import (
	"fmt"
	"strings"
)

func main() {
	l1 := buildList(9, 9, 9)
	l2 := buildList(9, 9, 9)
	printList(addTwoNumbers(l1, l2))
	l1 = buildList(9, 9, 9)
	l2 = buildList(9, 9)
	printList(addTwoNumbers(l1, l2))
	l1 = buildList()
	l2 = buildList()
	printList(addTwoNumbers(l1, l2))
	l1 = buildList(0)
	l2 = buildList(0)
	printList(addTwoNumbers(l1, l2))
	l1 = buildList(0)
	l2 = buildList(1)
	printList(addTwoNumbers(l1, l2))
	l1 = buildList(9)
	l2 = buildList(9)
	printList(addTwoNumbers(l1, l2))
}

func addTwoNumbers(l1, l2 *node) *node {
	len1, len2 := l1.len(), l2.len()
	if len1 < len2 {
		l1 = padWithZeros(l1, len1, len2)
	} else if len2 < len1 {
		l2 = padWithZeros(l2, len2, len1)
	}
	if l1 == nil && l2 == nil {
		return nil
	}
	sum, carry := addTwoNumHelper(l1, l2)
	if carry != 0 {
		sum = &node{carry, sum}
	}
	return sum
}

type node struct {
	Val  int
	Next *node
}

func addTwoNumHelper(l1, l2 *node) (*node, int) {
	if l1 == nil && l2 == nil {
		return nil, 0
	}

	sum, carry := addTwoNumHelper(l1.Next, l2.Next)
	result := &node{
		l1.Val + l2.Val + carry,
		sum,
	}
	if result.Val > 9 {
		carry = result.Val / 10
		result.Val %= 10
	} else {
		carry = 0
	}
	return result, carry
}

func padWithZeros(l *node, len, target int) *node {
	for len < target {
		newNode := &node{Val: 0}
		newNode.Next = l
		l = newNode
		len++
	}
	return l
}

func (l *node) len() int {
	counter := 0
	for l != nil {
		counter++
		l = l.Next
	}
	return counter
}

func buildList(nums ...int) *node {
	length := len(nums)
	if length == 0 {
		return nil
	}
	var l *node
	for i := length - 1; i >= 0; i-- {
		newNode := &node{nums[i], l}
		l = newNode
	}
	return l
}

func printList(l *node) {
	if l == nil {
		fmt.Println()
		return
	}
	var output strings.Builder
	for l != nil {
		output.WriteString(fmt.Sprintf("%v->", l.Val))
		l = l.Next
	}
	fmt.Println(output.String()[:output.Len()-2])
}
