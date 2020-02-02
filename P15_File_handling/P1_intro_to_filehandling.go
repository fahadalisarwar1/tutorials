package main

import (
	"log"
	"os"
)

// One of the fundamental aspects of UNIX is that everything is a file.
// The operating system provides us an interface to the device in the form of a file.
// The reader and writer interfaces in Go are similar abstractions.
// We simply read and write bytes, without the need to understand where or how the reader gets its data or where the writer
// is sending the data.

func main(){

	file, err := os.Create("dummyfile.txt")
	if err != nil{
		log.Fatal(err)
		return
	}

	statement := "this is a dummy file "
	byteArray := []byte(statement)

	file.Write(byteArray)
	file.Close()





}