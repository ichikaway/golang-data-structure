package main

import (
	"fmt"
	"hash/fnv"
)

const BIN_SIZE int = 5

type List struct {
	Key  string
	Val  string
	Next *List
}

type HashTable struct {
	count int
	binsize  int
	array []List
}

/**
文字列から、整数値のハッシュ値を返す
 */
func hashInt(s string) int {
	h := fnv.New32a()
	h.Write([]byte(s))
	return int(h.Sum32())
}

/**
HashMapの配列のIndexを求める
 */
func makeIndex(s string, binsize int) int {
	i := hashInt(s)
	return i % binsize
}

func newHashMap(binsize int) *HashTable {
	list := make([]List, binsize)
	return &HashTable{count: 0, binsize: binsize, array: list }
}

/**
HashMapの配列の中にある連結リストに、key valの値を追加
 */
func (list *List) add(key string, val string) bool {
	for {
		if list.Key == key {
			list.Val = val
			return true
		}
		if nil == list.Next {
			list.Next = &List{Key: key, Val: val}
			return true
		}
		list = list.Next
	}
	return false
}

/**
HashMapの配列の中にある連結リストから、keyに該当する値を返す
 */
func (list *List) get(key string) string {
	for {
		fmt.Println("key:", list.Key)
		if list.Key == key {
			return list.Val
		}
		if nil == list.Next {
			return ""
		}
		list = list.Next
	}
}

/**
HashTableからkeyを基準に値を取得する
Hashの配列からindexを求めて、その中の連結リストから該当keyを探索
 */
func (table *HashTable) get(key string) string {
	i := makeIndex(key, table.binsize)
	if table.array[i].Key == "" {
		return ""
	} else {
		return table.array[i].get(key)
	}
}

/**
HashTableにkey valを保存
Hashの配列からindexを求めて、その中の連結リストに追加
同一keyがあった場合は上書きする
 */
func (table *HashTable) add(key string, val string) {
	if key == "" {
		panic("Key needs not empty")
	}
	i := makeIndex(key, table.binsize)

	if table.count > table.binsize {
		*table = table.rehash()
	}

	//fmt.Println(i)
	//fmt.Println(table.array[i])
	table.count++
	if table.array[i].Key == "" {
		table.array[i] = List{Key: key, Val: val}
	} else {
		table.array[i].add(key, val)
	}
}

/**
新しいHashMapを作って、既存のデータをそちらに移動させる
そに際に、配列のサイズを2倍にする
 */
func (table *HashTable) rehash() HashTable{
	newBinsize := table.binsize * 2
	newTable := newHashMap(newBinsize)

	var_dump(table)

	size := len(table.array)
	for i := 0; i < size; i++ {
		if table.array[i].Key == "" {
			continue
		}
		newTable.add(table.array[i].Key, table.array[i].Val)

		var_dump(table.array[i].Next)
		var list *List = table.array[i].Next
		for {
			if list == nil {
				break
			}
			newTable.add(list.Key, list.Val)
			list = list.Next
		}
	}
	var_dump(newTable)
	return *newTable
}




func main() {

	hashMap := newHashMap(BIN_SIZE)
	hashMap.add("ichi", "kawa")
	hashMap.add("ichi2", "kawa2")
	hashMap.add("yasu", "shi")
	hashMap.add("111", "shi1")
	hashMap.add("222", "shi2")
	hashMap.add("333", "shi3")
	hashMap.add("444", "shi4")

	fmt.Println(hashMap.get("333"))

	//fmt.Println(hashMap)
	var_dump(hashMap.array[3].Next)
	var_dump(hashMap)
	//var_dump(hashMap)

}

func var_dump(v ...interface{}) {
	for _, vv := range (v) {
		fmt.Printf("%#v\n", vv)
	}
}
