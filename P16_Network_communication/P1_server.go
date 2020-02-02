package main

import (
	"encoding/gob"
	"fmt"
	"net"
	"strings"
)

// TCP is the primary protocol used for communication over the Internet. Any time you
// interact with a web page, play a multiplayer computer game, stream a movie, or video
// chat, there’s a good chance your computer is communicating with a remote server


func handleConnection(conn net.Conn){
	var msg string
	err := gob.NewDecoder(conn).Decode(&msg)

	if err != nil{
		fmt.Println(err)
	}else{
		msg_SERVER := strings.ToUpper(msg)
		fmt.Println(msg_SERVER)

	}





}


func main(){

	ip := "localhost"
	port := "9090"  // 0 - 65355 2^16
	address := ip + ":" + port
	listener, err := net.Listen("tcp", address)

	if err != nil{
		fmt.Println(err)
		return
	}

	for {
		connection, err := listener.Accept()

		if err != nil{
			fmt.Println(err)
			continue
		}

		fmt.Println("go a connection ", connection.RemoteAddr().String())

		defer connection.Close()

		go handleConnection(connection)




	}




}
