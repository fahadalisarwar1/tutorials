package main

import (
	"fmt"
	"os"
)

func main(){
	file, _ := os.Open("dummyfile.txt")
	defer file.Close()

	stats, _ := file.Stat()

	numberofBytes := stats.Size()
	fmt.Println(numberofBytes)

	bytesbuffer := make([]byte, numberofBytes)

	file.Read(bytesbuffer)

	data := string(bytesbuffer)
	fmt.Println(data)






}
