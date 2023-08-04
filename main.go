package main

import (
	"os"
)

type Node struct {
	num int
	index int
	Next *Node
	Previous *Node
}

func CheckForSort(a *Node) bool {
	temp := a.Next
	for ; temp != a; temp = temp.Next {
		if temp.Previous.num > temp.num {
			return false
		}
	}
	return true
}

func main() {
	if len(os.Args) > 1 {
		var a *Node
		primary := make(chan int, len(os.Args[1:]))
		sorted := make(chan int, len(os.Args[1:]))
		go PrimaryArgs(primary, os.Args[1:])
		go SortedArgs(sorted, os.Args[1:])
		a = ListInit(primary, sorted)
		if CheckForSort(a) {
			return
		}
		l := ListLenght(a)
		switch {
		case l <= 5 : SortSmall(a, l)
		case l <= 100 : SortMedium(a)
		default : SortLarge(a)
		}
	}
}