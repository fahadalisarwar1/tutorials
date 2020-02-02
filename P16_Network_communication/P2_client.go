package main

import (
	"encoding/gob"
	"fmt"
	"net"
)

func main(){
	clientConn, err := net.Dial("tcp", "localhost:9090")
	if err != nil{
		fmt.Println(err)
		return
	}

	mess := "Fahad Ali "

	err = gob.NewEncoder(clientConn).Encode(mess)

	if err != nil{
		fmt.Println(err)
		return
	}

	clientConn.Close()

}
