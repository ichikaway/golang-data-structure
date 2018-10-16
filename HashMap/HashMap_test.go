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


}
