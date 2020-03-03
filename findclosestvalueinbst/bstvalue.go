package bstvalue

import "math"

type BST struct {
	value int
	left  *BST
	right *BST
}

// Avg: O(Log(N)) time || O(Log(N)) space
// Wst: O(N) time || O(N) space
func (tree *BST) FindClosestValueRecursively(target int) int {
	return tree.helperRecursive(target, math.MaxInt32)
}

func (tree *BST) helperRecursive(target, closest int) int {
	if absdiff(target, closest) > absdiff(target, tree.value) {
		closest = tree.value
	}
	if target > tree.value && tree.right != nil {
		return tree.right.helperRecursive(target, closest)
	} else if target < tree.value && tree.left != nil {
		return tree.left.helperRecursive(target, closest)
	}
	return closest
}

// Avg: O(Log(N)) time || O(1) space
// Wst: O(N) time \\ O(1) space
func (tree *BST) FindClosestValueIteratively(target int) int {
	return tree.helperIterative(target, math.MaxInt32)
}

func (tree *BST) helperIterative(target, closest int) int {
	currentNode := tree
	for currentNode != nil {
		if absdiff(target, closest) > absdiff(target, currentNode.value) {
			closest = currentNode.value
		}
		if target > currentNode.value {
			currentNode = currentNode.right
		} else if target < currentNode.value {
			currentNode = currentNode.left
		} else {
			break
		}
	}
	return closest
}

func absdiff(a, b int) int {
	result := math.Abs(float64(a) - float64(b))
	return int(result)
}
