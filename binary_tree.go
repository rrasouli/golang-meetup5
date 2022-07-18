package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"sync"
)

type TreeNode struct {
	key       int
	value     int
	leftNode  *TreeNode
	rightNode *TreeNode
}

type BinarySearchTree struct {
	rootNode *TreeNode
	lock     sync.RWMutex
}

func (tree *BinarySearchTree) readFile(filePath string) string {

	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatal(err)
	}

	return string(content)

}

func (tree *BinarySearchTree) InsertElement(key int, value int) {
	tree.lock.Lock()
	defer tree.lock.Unlock()

	var newNode *TreeNode = &TreeNode{key: key, value: value}
	if tree.rootNode == nil {
		tree.rootNode = newNode
	} else {
		InsertNode(tree.rootNode, newNode)
	}
}

func InsertNode(parentNode *TreeNode, newNode *TreeNode) {

	if newNode.key < parentNode.key {
		if parentNode.leftNode == nil {
			parentNode.leftNode = newNode
		} else {
			InsertNode(parentNode.leftNode, newNode)
		}
	}

	if newNode.key > parentNode.key {
		if parentNode.rightNode == nil {
			parentNode.rightNode = newNode
		} else {
			InsertNode(parentNode.rightNode, newNode)
		}
	}
}

func (tree *BinarySearchTree) PreOrderTraverseTree(function func(int)) {
	tree.lock.RLock()
	defer tree.lock.RUnlock()
	preOrderTraverseTree(tree.rootNode, function)
}

//  preOrderTraverseTree method
func preOrderTraverseTree(treeNode *TreeNode, function func(int)) {
	if treeNode != nil {
		function(treeNode.value)
		preOrderTraverseTree(treeNode.leftNode, function)
		preOrderTraverseTree(treeNode.rightNode, function)
	}
}

func (tree *BinarySearchTree) SearchNode(key int) bool {
	tree.lock.RLock()
	defer tree.lock.RUnlock()
	return searchNode(tree.rootNode, key)
}

func searchNode(treeNode *TreeNode, key int) bool {
	if treeNode == nil {
		return false
	}
	if treeNode.key == key {
		return true
	}
	if key < treeNode.key {
		return searchNode(treeNode.leftNode, key)
	}
	if key > treeNode.key {
		return searchNode(treeNode.rightNode, key)
	}
	return true
}

//printTreeNode method
func printTreeNode(treeNode *TreeNode) {
	if treeNode != nil {
		fmt.Println(" Value", treeNode.value)
		fmt.Printf("TreeNode Left")
		printTreeNode(treeNode.leftNode)
		fmt.Printf("TreeNode Right")
		printTreeNode(treeNode.rightNode)
	} else {
		fmt.Printf("Nil\n")
	}

}

func main() {

	var tree *BinarySearchTree = &BinarySearchTree{}
	charsArray := tree.readFile("charsFile.txt")
	asciiArray := []rune(charsArray)
	for _, runeValue := range asciiArray {
		tree.InsertElement(int(runeValue), int(runeValue))

	}
	tree.PreOrderTraverseTree(func(i int) {})
	searchVal := 'ד'
	search := tree.SearchNode(int(rune(searchVal)))
	if search {
		fmt.Printf("Success found %v\n", string(searchVal))
	} else {
		fmt.Printf("Failed %v Not Found\n", string(searchVal))
	}

}
