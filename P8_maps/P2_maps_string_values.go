package main

import (
	"fmt"
)

func main(){
	fruitData := map[string]int{
		"apples": 22,
		"oranges": 34,
	}
	fmt.Println(fruitData["apples"])
	fruitData["oranges"] = 34
	fruitData["banana"] = 355

	for fruit, quantity := range fruitData{
		fmt.Println("Fruit: ", fruit, " quantity: ", quantity)
	}

	fruitData["watermelon"] = 3444

	fmt.Println("Before",len(fruitData))

	delete(fruitData, "banana")
	fmt.Println("After",len(fruitData))
	fmt.Println(fruitData)



}
