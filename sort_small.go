package main

func SortThree(a *Node) *Node {
	first, second, third := a.num, a.Next.num, a.Next.Next.num
	if first < second && second < third {return a}
	if third > second && third > first {
		a = SA(a, true)
	} else if first > second && first > third {
		a = RA(a, true)
		if second > third {a = SA(a, true)}
	} else if second > first && second > third {
		a = RRA(a, true)
		if third > first {a = SA(a, true)}
	}
	return a
}

func SortFour(a *Node, index int) *Node {
	var b *Node
	if DirectionToRotate(a, index) {
		for a.index != index {a = RA(a, true)}
	} else {
		for a.index != index {a = RRA(a, true)}
	}
	a, b = PB(a, b, true)
	a = SortThree(a)
	a, b = PA(a, b, true)
	return a
}

func SortFive(a *Node) *Node {
	var b *Node
	if DirectionToRotate(a, 0) {
		for a.index != 0 {a = RA(a, true)}
	} else {
		for a.index != 0 {a = RRA(a, true)}
	}
	a, b = PB(a, b, true)
	a = SortFour(a, 1)
	a, b = PA(a, b, true)
	return a
}

func SortSmall(a *Node, lenght int) {
	switch lenght {
	case 2 : a = SA(a, true)
	case 3 : a = SortThree(a)
	case 4 : a = SortFour(a, 0)
	case 5 : a = SortFive(a)
	}
}