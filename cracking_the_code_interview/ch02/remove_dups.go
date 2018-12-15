package main

func main() {

}

func removeDuplicates(l *node) {
	listMap := make(map[int]bool)
	var prev *node
	for l != nil {
		status := listMap[l.Val]
		if status {
			prev.Next = l.Next
		} else {
			listMap[l.Val] = true
			prev = l
		}
		l = l.Next
	}
}

func removeDuplicatesNoBuff(l *node) {
	current := l
	for current != nil {
		runner := current
		for runner.Next != nil {
			if current.Val == runner.Next.Val {
				runner.Next = runner.Next.Next
			} else {
				runner = runner.Next
			}
		}
		current = current.Next
	}
}

type node struct {
	Val  int
	Next *node
}
