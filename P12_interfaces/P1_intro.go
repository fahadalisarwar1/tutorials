package main

import (
	"fmt"
	"math"
)

type geometry interface{
	area() float64
	perimeter() float64
}

type rectangle struct{
	widht, height float64
}

type circle struct{
	radius float64
}


func (r rectangle) area()(float64){
	return r.widht * r.height
}

func (r rectangle) perimeter()(float64){
	return 2 * (r.widht + r.height)
}

func (c circle) area()(float64){
	return math.Pi *c.radius *c.radius
}

func (c circle) perimeter()(float64){
	return 2 * math.Pi * c.radius
}

func measure(g geometry){
	fmt.Println(g)
	fmt.Println(g.perimeter())

	fmt.Println(g.area())

}


func main(){
	c := circle{radius:10}
	r := rectangle{widht:30, height:40}


	measure(c)
	measure(r)
	// interface defines the behaviour of an object

}