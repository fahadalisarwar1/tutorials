package main

import "fmt"

func main(){
	bytes := []byte{0x48, 0x49}

	mystr := string(bytes)
	fmt.Println(mystr)

	decBytes := []byte{67, 97, 102, 195, 169}

	str := string(decBytes)
	fmt.Println(str)
	fmt.Println(len(str))
	fmt.Println(str[0])
}
