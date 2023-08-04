package main

import (
	"fmt"
)

func RRA(a *Node, write bool) (aR *Node) {
	aR = a.Previous
	if write {
		fmt.Println("rra")
	}
	return
}

func RRB(b *Node, write bool) (bR *Node) {
	bR = b.Previous
	if write {
		fmt.Println("rrb")
	}
	return
}

func RRR(a, b *Node, write bool) (aR, bR *Node) {
	aR = a.Previous
	bR = b.Previous
	if write {
		fmt.Println("rrr")
	}
	return
}