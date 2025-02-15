package main

/*
Exercise: Equivalent Binary Trees
There can be many different binary trees with the same sequence of values
stored in it. For example, here are two binary trees storing the sequence 1, 1,
2, 3, 5, 8, 13.

          3                                      8
         / \                                    / \
        /   \                                  3   13
       1     8                                / \
      / \   / \                              1   5
     1   2 5  13                            / \
                                           1   2

A function to check whether two binary trees store the same sequence is quite
complex in most languages. We'll use Go's concurrency and channels to write a
simple solution.

This example uses the tree package, which defines the type:

type Tree struct {
    Left  *Tree
    Value int
    Right *Tree
}

1. Implement the Walk function.

2. Test the Walk function.

The function tree.New(k) constructs a randomly-structured (but always sorted)
binary tree holding the values k, 2k, 3k, ..., 10k.

Create a new channel ch and kick off the walker:

```
go Walk(tree.New(1), ch)
```

Then read and print 10 values from the channel. It should be the numbers 1, 2,
3, ..., 10.

3. Implement the Same function using Walk to determine whether t1 and t2 store
the same values.

4. Test the Same function.

Same(tree.New(1), tree.New(1)) should return true, and Same(tree.New(1),
tree.New(2)) should return false.

The documentation for Tree can be found here:
https://pkg.go.dev/golang.org/x/tour/tree#Tree
*/

import (
	"fmt"

	"golang.org/x/tour/tree"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	var walker func(t *tree.Tree)
	walker = func(t *tree.Tree) {
		if t == nil {
			return // base case to stop the tree walk
		}
		// traverse the tree in-order
		walker(t.Left)
		ch <- t.Value
		walker(t.Right)
	}
	walker(t)
	close(ch) // close the channel
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	ch1, ch2 := make(chan int, 10), make(chan int, 10)
	go Walk(t1, ch1)
	go Walk(t2, ch2)
	for v1 := range ch1 {
		v2, ok := <-ch2
		// if no item in ch2 or v1 != v2
		if !ok || v1 != v2 {
			return false
		}
	}
	// check if there are still items in ch2
	_, ok := <-ch2
	return !ok
}

func main() {
	fmt.Println(Same(tree.New(1), tree.New(1))) // true
	fmt.Println(Same(tree.New(2), tree.New(1))) // false
}
