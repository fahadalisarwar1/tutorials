package main

import (
	"fmt"
	"hash/crc32"
)

func main(){
	hasher := crc32.NewIEEE()

	data := []byte("My name is Fahad")

	hasher.Write(data)


	hashkey := hasher.Sum32()

	fmt.Println(hashkey)

}
