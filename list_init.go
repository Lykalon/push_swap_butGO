package main

import (
	"os"
	"strconv"
)

func PrimaryArgs(c chan int, args []string) {
	defer close(c)
	var index []int
	for _, str := range args {
		i, ok := strconv.Atoi(str)
		if ok != nil {
			os.Stderr.Write([]byte("Error\n"))
			os.Exit(1)
		}
		index = append(index, i)
	}
	for _, i := range index {
		c <- i
	}
}

func NewItem(num int) *Node {
	new := new(Node)
	new.num = num
	new.Next = new
	new.Previous = new
	return new
}

func PushBack(a, new *Node) *Node {
	var first *Node

	if a == nil {
		return new
	}
	first = a
	a = a.Previous
	first.Previous = new
	new.Next = first
	new.Previous = a
	a.Next = new
	return first
}

func Indexation(a *Node, sorted chan int) {
	var i int
	for num := range sorted {
		for ; num != a.num; a = a.Next {}
		a.index = i
		i++
	}
}

func ListInit(primary, sorted chan int) *Node {
	var a, new *Node

	for num := range primary {
		new = NewItem(num)
		a = PushBack(a, new)
	}
	Indexation(a, sorted)
	return a
}