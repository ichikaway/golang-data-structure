package main

import (
	"fmt"
	"testing"
)

func TestBasicfunction(t *testing.T) {

	hashMap := newHashMap(BIN_SIZE)
	hashMap.add("ichi", "kawa")
	hashMap.add("ichi2", "kawa2")
	hashMap.add("333", "shi3")

	result := hashMap.get("333")
	if result != "shi3" {
		t.Fatal("failed test: " + result)
	}

	result = hashMap.get("nokey")
	if result != "" {
		t.Fatal("failed test: " + result)
	}
}


/**
	要素を追加していくと、rehashが走り、配列のサイズが倍になることを確認
 */
func TestRehash(t *testing.T) {

	hashMap := newHashMap(5)
	hashMap.add("ichi", "kawa")
	hashMap.add("ichi2", "kawa2")
	size1 := hashMap.binsize
	count1 := hashMap.count
	hashMap.add("yasu", "shi")
	hashMap.add("111", "shi1")
	hashMap.add("222", "shi2")
	hashMap.add("333", "shi3")
	hashMap.add("444", "shi4")
	size2 := hashMap.binsize
	hashMap.add("555", "shi5")
	hashMap.add("666", "shi6")
	size3 := hashMap.binsize
	count2 := hashMap.count


	result := hashMap.get("333")
	if result != "shi3" {
		t.Fatal("failed test: " + result)
	}

	if size1 != 5 {
		t.Fatal("failed test size1")
	}
	if size2 != 10 || size3 != 10 {
		fmt.Println(size2)
		t.Fatal("failed test size2,3")
	}
	if count1 != 2 || count2 != 9 {
		t.Fatal("failed test count")
	}
}
