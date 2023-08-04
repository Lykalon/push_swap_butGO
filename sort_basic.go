package main

func DirectionToRotate(stack *Node, index int) bool {
	var count_r int
	for stack.index != index {
		stack = stack.Next
		count_r++
	}
	if count_r <= ListLenght(stack) / 2 {
		return true
	}
	return false
}

func SecondRun(b *Node, maxIndex int) *Node {
	var a *Node

	for b != nil {
		if DirectionToRotate(b, maxIndex) {
			for b.index != maxIndex {
				b = RB(b, true)
			}
		} else {
			for b.index != maxIndex {
				b = RRB(b, true)
			}
		}
		a, b = PA(a, b, true)
		maxIndex--
	}
	return a
}

func FirstRun(a *Node, min, max int) *Node {
	var b *Node

	for a != nil {
		if min <= a.index && a.index <= max {
			a, b = PB(a, b, true)
			min++
			max++
		} else if a.index > max {
			a = RA(a, true)
		} else if a.index < min {
			a, b = PB(a, b, true)
			b = RB(b, true)
			min++
			max++
		}
	}
	return b
}