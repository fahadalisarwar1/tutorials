package main

import "fmt"

func main(){
	greet := func(userName string)(greetings string){
		return "Hello user "+userName+" Welcome to this tutorial"
	}
	fmt.Println(greet("Fahad ali sarwar"))
	fmt.Println(greet("john"))

}