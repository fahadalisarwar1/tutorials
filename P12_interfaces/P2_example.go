package main

import "fmt"

type human interface{
	sayHello() string
}


type man struct{
	greeting string
}

type women struct {
	greeting string
}

type child struct {
	greeting string
}

func (m man) sayHello()(string){
	return m.greeting
}

func (w women) sayHello()(string){
	return w.greeting
}

func (c child) sayHello()(string){
	return c.greeting
}


func greetUser(h human){
	fmt.Println(h.sayHello())

}


func main(){
	man1 := man{"Hello how are you"}
	greetUser(man1)

	women1 := women{"Hello!"}
	greetUser(women1)

	kid := child{"I'm hungry feed me please!!!"}
	greetUser(kid)

}