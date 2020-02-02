package main

import "fmt"

func main(){
	var firstName string
	var ptrFirstName *string

	firstName = "Fahad"
	ptrFirstName = &firstName

	fmt.Println(firstName)
	fmt.Println(ptrFirstName)

	fmt.Println(&ptrFirstName)
	fmt.Println(*ptrFirstName)



}
