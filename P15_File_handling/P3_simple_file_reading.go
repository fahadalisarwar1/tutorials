package main

import (
	"fmt"
	"io/ioutil"
)

func main(){
	data, err := ioutil.ReadFile("dummyfile.txt")

	if err != nil{
		return
	}

	fmt.Println(data)
	fmt.Println(string(data))

}
