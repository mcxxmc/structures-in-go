package main

import (
	"fmt"
	"some-data-structures/structures"
)

func main() {
	nums := []int{1, 2, 3, 4}
	stk1 := structures.NewStack()
	for _, num := range nums {
		stk1.Push(num)
	}
	stk2 := stk1.Copy()
	fmt.Println(stk1.Pop())
	for stk2.HasNext() {
		fmt.Println(stk2.Pop())
	}
}
