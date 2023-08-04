package main

func SortMedium(a *Node) {
	min := 0
	max := 15

	b := FirstRun(a, min, max)
	a = SecondRun(b, ListLenght(b) - 1)
}