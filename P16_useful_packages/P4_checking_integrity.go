package main

import (
	"fmt"
	"hash/crc32"
	"io"
	"os"
)

func getHashValue(filename string)(uint32, error){
	file, err := os.Open(filename)
	if err != nil{
		fmt.Println("Couldnt open file")
		return 0, err
	}

	defer file.Close()

	hasher := crc32.NewIEEE()

	_, err = io.Copy(hasher, file)
	if err != nil{
		fmt.Println("ERror couldnt hash")
		return 0, err

	}


	return hasher.Sum32(), nil

}

func main(){
	h1, err := getHashValue("file1.txt")

	if err != nil{
		fmt.Println(err)
	}

	h2, err := getHashValue("file2.txt")
	if err !=nil{
		fmt.Println(err)
	}

	fmt.Println(h1, h2, h1 == h2)

}
