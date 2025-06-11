package binarysearchtree

type BinarySearchTree struct {
	left  *BinarySearchTree
	data  int
	right *BinarySearchTree
}

// NewBst creates and returns a new BinarySearchTree.
func NewBst(i int) *BinarySearchTree {
	return &BinarySearchTree{data: i}
}

// Insert inserts an int into the BinarySearchTree.
// Inserts happen based on the rules of a binary search tree
func (bst *BinarySearchTree) Insert(i int) {
	newNode := &BinarySearchTree{data: i}
	curr := bst.data

	if i > curr {
		if bst.right == nil {
			bst.right = newNode
			return
		}

		bst.right.Insert(i)
		return
	}

	if bst.left == nil {
		bst.left = newNode
		return
	}

	bst.left.Insert(i)
}

// SortedData returns the ordered contents of BinarySearchTree as an []int.
// The values are in increasing order starting with the lowest int value.
// A BinarySearchTree that has the numbers [1,3,7,5] added will return the
// []int [1,3,5,7].
func (bst *BinarySearchTree) SortedData() []int {
	result := []int{}
	if bst.left != nil {
		result = bst.left.SortedData()
	}
	result = append(result, bst.data)
	if bst.right != nil {
		rightSorted := bst.right.SortedData()
		result = append(result, rightSorted...)
	}
	return result
}
