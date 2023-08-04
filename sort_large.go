package main

func SortLarge(a *Node) {
	min := 0
	max := 30

	b := FirstRun(a, min, max)
	a = SecondRun(b, ListLenght(b) - 1)
}