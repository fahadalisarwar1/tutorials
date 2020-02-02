package main

import (
	"log"
	"os"
)

func main(){
	fileinfo, err := os.Stat("dummyfile.txt")
	if err != nil{
		if os.IsNotExist(err){
			log.Fatal("File doesnt exist")
		}
	}

	log.Println(fileinfo)
}
