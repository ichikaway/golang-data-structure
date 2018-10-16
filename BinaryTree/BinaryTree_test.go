package main

import (
	"testing"
)

func TestBasicfunction(t *testing.T) {
	bintree := createBinTree(4, "bbb")
	bintree.add(3, "ccc")
	bintree.add(2, "bbb2")
	bintree.add(1, "aaa")
	bintree.add(8, "888")
	bintree.add(5, "555")
	bintree.add(9, "999")
	//fmt.Println("eabcd" < "def")

	result := bintree.get(9)
	if result != "999a" {
		t.Fatal("failed test: " + result)
	}

	result = bintree.get(100)
	if result != "a" {
		t.Fatal("failed test: " + result)
	}
}
