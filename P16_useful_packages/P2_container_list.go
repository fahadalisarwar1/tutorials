package main

import (
	"container/list"
	"fmt"
)

func main(){

	ll := list.New()

	ll.PushFront(200)

	ll.PushBack(300)

	ll.PushBack(500)


	for element := ll.Front(); element != nil; element = element.Next(){
		fmt.Println(element.Value)

	}




}
