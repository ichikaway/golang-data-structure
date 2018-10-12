package main

import "fmt"

type Node struct {
	Key int
	Val string
	left *Node
	right *Node
}

type BinTree struct {
	root *Node
}

func (node *Node) add (key int, val string) {
	if node.Key == key {
		node.Val = val
	} else if node.Key > key {
		//fmt.Println(key, node.left)
		if node.left == nil {
			node.left = createNode(key, val)
		} else {
			node.left.add(key,val)
		}
		//fmt.Println(key, node.left)
	} else if node.Key < key {
		//fmt.Println(key, node.right)
		if node.right == nil {
			node.right = createNode(key, val)
		} else {
			node.right.add(key,val)
		}
		//fmt.Println(key, node.right)
	}
}

func (node *Node) get (key int) string {
	if node.Key == key {
		return node.Val
	} else if node.Key > key {
		if node.left == nil {
			return ""
		} else {
			return node.left.get(key)
		}
	} else if node.Key < key {
		if node.right == nil {
			return ""
		} else {
			return node.right.get(key)
		}
	}
	return ""
}

func (tree *BinTree) add (key int, val string) {
	tree.root.add(key, val)
}

func (tree *BinTree) get (key int) string {
	return tree.root.get(key)
}

func createNode(key int, val string) *Node {
	return &Node{Key:key, Val:val}
}


func createBinTree(key int, val string) BinTree {
	node := createNode(key, val)
	return BinTree{root:node}
}

func main() {
	bintree := createBinTree(4, "bbb")
	bintree.add(3, "ccc")
	bintree.add(2, "bbb2")
	bintree.add(1, "aaa")
	bintree.add(8, "888")
	bintree.add(5, "555")
	bintree.add(9, "999")
	//fmt.Println("eabcd" < "def")

	fmt.Println("val:", bintree.get(9))
	fmt.Println("val:", bintree.get(100))
/*
	var_dump(bintree.root)
	var_dump(bintree.root.left)
	var_dump(bintree.root.left.left)
	var_dump(bintree.root.left.left.left)
	var_dump(bintree.root.right)
	var_dump(bintree.root.right.left)
	var_dump(bintree.root.right.right)*/
}


func var_dump(v ...interface{}) {
	for _, vv := range (v) {
		fmt.Printf("%#v\n", vv)
	}
}
