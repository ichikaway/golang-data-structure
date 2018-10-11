package main

import (
	"fmt"
)

type List struct {
	Id   int
	Name string
	Next *List
}

func (list *List) add(i int, name string) bool {
	for {
		if nil == list.Next {
			list.Next = &List{Id: i, Name: name}
			return true
		}
		list = list.Next
	}
	return false
}

func (list *List) get(i int) *List {
		if list.Id == i {
			return list
		}
		if list.Next == nil {
			return &List{}
		}
		return list.Next.get(i)

	/*	for {
		//fmt.Println(list.Id)
		if list.Id == i {
			return list
		}
		if list.Next == nil {
			return &List{}
		}
		list = list.Next
	}*/
}

func newList(id int, name string) *List {
	return &List{Id: id, Name: name}
}

func main() {
	/*
	i1 := &List{Id:1}
	i1.Next = &List{Id:2}
	i1.Next.Next = &List{Id:3}

	for {
		fmt.Println(i1.Id)
		if i1.Next == nil {
			break
		}
		i1 = i1.Next
	}
	*/

	i1 := newList(1, "ichi")
	i1.add(2, "Two")
	i1.add(3, "Three")

	anser := i1.get(3)
	anser2 := i1.get(1)
	var_dump(anser)
	var_dump(anser2)

}

func var_dump(v ...interface{}) {
	for _, vv := range (v) {
		fmt.Printf("%#v\n", vv)
	}
}
