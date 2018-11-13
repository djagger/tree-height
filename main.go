package main

import "fmt"

// Please write the java code for the following task - we need to calculate height of a tree when
// given its root node. Please design your own classes. We don’t need tree rebalancing etc.
// Just the basic class(es) sufficient to represent the tree and be able to calculate the height.
// Feel free to make any assumptions. If you have time please provide unit test.
// Max time limit is 30 minutes to complete.

type node struct {
	value int
	left  *node
	right *node
}

func (n *node) height() int {
	if n == nil {
		return 0
	}

	return getHeight(n)
}

func getHeight(n *node) int {
	leftHeight := 0
	rightHeight := 0

	if n.left != nil {
		leftHeight = getHeight(n.left)
		leftHeight++
	} else {
		leftHeight = 0
	}

	if n.right != nil {
		rightHeight = getHeight(n.right)
		rightHeight++
	} else {
		rightHeight = 0
	}

	height := 0

	if rightHeight > leftHeight {
		height += rightHeight
	} else {
		height += leftHeight
	}

	return height
}

func main() {
	// For example:
	//
	//        1
	//      /  \
	//     /    \
	//    2     12
	//   / \     \
	//  3  10    5
	//    /
	//   7
	//  / \
	// 9   4

	treeExample := node{
		value: 1,
		left: &node{
			value: 2,
			left: &node{
				value: 3,
			},
			right: &node{
				value: 10,
				left: &node{
					value: 7,
					left: &node{
						value: 9,
					},
					right: &node{
						value: 4,
					},
				},
			},
		},
		right: &node{
			value: 12,
			right: &node{
				value: 5,
			},
		},
	}

	fmt.Println("Height of this tree is:", treeExample.height())
}
