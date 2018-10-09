package main

import (
	"fmt"
	"os"
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
	for {
		fmt.Println(list.Id)
		fmt.Println(list)
		fmt.Println(&list)
		if list.Id == i {
			return list
		}
		if list.Next == nil {
			fmt.Println("----------", &list)
			return &List{}
		}

		fmt.Println("aa")
		fmt.Println(list)
		list = list.Next
		//*list = *list.Next
		//*list = List{}
		fmt.Println(list)
	}
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

	//i1 := &List{Id:1}
	//var i1 *List = &List{Id:1, Name:"ichi"}
	i1 := newList(1, "ichi")
	//var i1 List = List{Id:1}
	i1.add(2, "Two")
	i1.add(3, "Three")

	anser := i1.get(3)
	anser2 := i1.get(1)
	var_dump(anser)
	var_dump(anser2)
	os.Exit(0)

	var_dump(i1)

	var val *List = i1
	for {
		fmt.Println(val.Id, val.Name)
		if val.Next == nil {
			break
		}
		val = val.Next
	}
	var_dump(val)
	//fmt.Println(i1.Next.Id)
	var_dump(i1)

}

func var_dump(v ...interface{}) {
	for _, vv := range (v) {
		fmt.Printf("%#v\n", vv)
	}
}
