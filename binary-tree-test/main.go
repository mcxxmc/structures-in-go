package main

import (
	"fmt"
	structures "some-data-structures/structures"
	"strconv"
)

type Record struct {
	Val int
}

func (r *Record) String() string {
	return strconv.Itoa(r.Val)
}

func (r *Record) Copy() interface{} {
	return &Record{Val: r.Val}
}

func main() {
	fmt.Println("The first tree")
	bTree := structures.NewIntBinaryTree()
	insertions := []int{10, 5, 15, 3, 8, 20, 0, 24}
	for _, num := range insertions {
		bTree.Insert(num)
	}
	bTree.PrintTree()
	depth := bTree.Depth()
	fmt.Println("the depth of the tree is " + strconv.Itoa(depth))

	fmt.Println("make a copy")
	copied := bTree.Copy()
	copied.PrintTree()
	depth = copied.Depth()
	fmt.Println("the depth of the tree is " + strconv.Itoa(depth))

	fmt.Println("rebuild a tree")
	newTree := bTree.Rebuild()
	newTree.PrintTree()
	depth = newTree.Depth()
	fmt.Println("the depth of the tree is " + strconv.Itoa(depth))

	fmt.Println("Delete 8")
	newTree.Delete(8)
	newTree.PrintTree()
	depth = newTree.Depth()
	fmt.Println("the depth of the tree is " + strconv.Itoa(depth))

	fmt.Println("Delete 20")
	newTree.Delete(20)
	newTree.PrintTree()
	depth = newTree.Depth()
	fmt.Println("the depth of the tree is " + strconv.Itoa(depth))

	fmt.Println("Delete 24")
	newTree.Delete(24)
	newTree.PrintTree()
	depth = newTree.Depth()
	fmt.Println("the depth of the tree is " + strconv.Itoa(depth))

	fmt.Println("Delete 15")
	newTree.Delete(15)
	newTree.PrintTree()
	depth = newTree.Depth()
	fmt.Println("the depth of the tree is " + strconv.Itoa(depth))

	fmt.Println("Customized interface")
	// test with customized struct
	compare := func(a, b interface{}) int {
		ra := a.(*Record)
		rb := b.(*Record)
		if ra.Val > rb.Val {
			return 1
		} else if ra.Val == rb.Val {
			return 0
		}
		return -1
	}

	customizedTree := structures.NewBinaryTree(compare)
	for _, num := range insertions {
		customizedTree.Insert(&Record{Val: num})
	}
	customizedTree.PrintTree()
	depth = customizedTree.Depth()
	fmt.Println("the depth of the tree is " + strconv.Itoa(depth))
}
