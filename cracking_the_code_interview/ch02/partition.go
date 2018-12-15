package main

func main() {

}

func partition(l *node, p int) *node {
	if l == nil || l.Next == nil {
		return l
	}
	var leftPartitionEnd *node
	var prev *node
	runner := l
	for runner != nil {
		if runner.Val < p {
			if prev == nil {
				leftPartitionEnd = runner
			} else if leftPartitionEnd == nil {
				leftPartitionEnd = runner
				prev.Next = runner.Next
				runner.Next = l
				l = runner

			} else {
				temp := runner
				prev.Next = runner.Next
				temp.Next = leftPartitionEnd.Next
				leftPartitionEnd.Next = temp
				leftPartitionEnd = temp
			}
		}
		prev = runner
		runner = runner.Next
	}
	return l
}

type node struct {
	Val  int
	Next *node
}
