package main

import (
	"fmt"
)

func SA(a *Node, write bool) (aR *Node) {
	if ListLenght(a) > 1 {
		aR = Swap(a)
	}
	if write {
		fmt.Println("sa")
	}
	return
}

func SB(b *Node, write bool) (bR *Node) {
	if ListLenght(b) > 1 {
		bR = Swap(b)
	}
	if write {
		fmt.Println("sb")
	}
	return
}

func SS(a, b *Node, write bool) (aR, bR *Node) {
	if ListLenght(a) > 1 {
		aR = Swap(a)
	}
	if ListLenght(b) > 1 {
		bR = Swap(b)
	}
	if write {
		fmt.Println("ss")
	}
	return
}

func PA(a, b *Node, write bool) (aR, bR *Node) {
	if ListLenght(b) > 0 {
		bR, aR = Push(b, a)
	}
	if write {
		fmt.Println("pa")
	}
	return
}

func PB(a, b *Node, write bool) (aR, bR *Node) {
	if ListLenght(a) > 0 {
		aR, bR = Push(a, b)
	}
	if write {
		fmt.Println("pb")
	}
	return
}