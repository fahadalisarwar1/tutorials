package main

import (
	"fmt"
	"time"
)

func main(){

	presentTime := time.Now()
	fmt.Println(presentTime)

	birthday := time.Date(1994,10, 1, 00, 00, 00,00, time.Local)
	fmt.Println(birthday)

	fmt.Printf("%T\n", birthday)

	fmt.Println(birthday.Year())

	fmt.Println(birthday.Before(presentTime))
	fmt.Println(birthday.After(presentTime))

	fmt.Println(birthday.Equal(presentTime))

	age := presentTime.Sub(birthday)
	fmt.Println(age.Hours()/(365*24))




}
