package main

func extract(stack *Node) (stackR, item *Node) {
	item = stack
	last := stack.Previous
	stackR = stack.Next
	item.Next = item
	item.Previous = item
	if item == stackR {
		stackR = nil
		return
	}
	stackR.Previous = last
	last.Next = stackR
	return
}

func Swap(stack *Node) *Node {
	stack, first := extract(stack)
	stack, second := extract(stack)
	stack = PushBack(stack, first)
	stack = stack.Previous
	stack = PushBack(stack, second)
	stack = stack.Previous
	return stack
}

func Push(src, dst *Node) (srcR, dstR *Node) {
	srcR, temp := extract(src)
	dst = PushBack(dst, temp)
	dstR = dst.Previous
	return
}