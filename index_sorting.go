package main

import (
	"os"
	"strconv"
)

func CheckRepeats(index []int) {
	for i := 0; i < len(index) - 1; i++ {
		if index[i] == index[i + 1] {
			os.Stderr.Write([]byte("Error\n"))
			os.Exit(1)
		}
	}
}

func BubbleSort(index []int) []int {
	l := len(index) - 1
	for j := 0; j < l; l-- {
		for i := 0; i < l; i++ {
			if index[i] > index[i + 1] {
				index[i], index[i + 1] = index[i + 1], index[i]
			}
		}
	}
	return index
}

func SortedArgs(c chan int, args []string) {
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
	index = BubbleSort(index)
	CheckRepeats(index)
	for _, i := range index {
		c <- i
	}
}