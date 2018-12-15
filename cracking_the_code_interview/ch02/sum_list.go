package main

func main() {

}

func addTwoNumbers(l1 *node, l2 *node) *node {
	return addTwoNumHelper(l1, l2, 0)
}

func addTwoNumHelper(l1, l2 *node, carry int) *node {
	if carry == 0 {
		if l1 == nil && l2 == nil {
			return nil
		} else if l1 == nil {
			return l2
		} else if l2 == nil {
			return l1
		}
	}
	val := carry
	l1Next := l1
	l2Next := l2
	if l1 != nil {
		val += l1.Val
		l1Next = l1.Next
	}
	if l2 != nil {
		val += l2.Val
		l2Next = l2.Next
	}
	if val > 9 {
		carry = 1
		val %= 10
	} else {
		carry = 0
	}
	result := &node{
		Val:  val,
		Next: nil,
	}
	if l1 != nil || l2 != nil {
		more := addTwoNumHelper(l1Next, l2Next, carry)
		result.Next = more
	}
	return result
}

type node struct {
	Val  int
	Next *node
}
