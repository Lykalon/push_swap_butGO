package main

func ListLenght(a *Node) int {
	count := 1
	if a == nil {
		return 0
	}
	for temp := a; temp.Next != a; temp = temp.Next {
		count++
	}
	return count
}