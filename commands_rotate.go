package main

import (
	"fmt"
)

func RA(a *Node, write bool) (aR *Node) {
	aR = a.Next
	if write {
		fmt.Println("ra")
	}
	return
}

func RB(b *Node, write bool) (bR *Node) {
	bR = b.Next
	if write {
		fmt.Println("rb")
	}
	return
}

func RR(a, b *Node, write bool) (aR, bR *Node) {
	aR = a.Next
	bR = b.Next
	if write {
		fmt.Println("rr")
	}
	return
}
