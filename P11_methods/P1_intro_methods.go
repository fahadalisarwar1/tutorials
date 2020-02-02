package main

import "github.com/fahadalisarwar1/tutorials/computer"

func main(){

	p1 := computer.ProcessingUnit{
		Maker: "intel",
		MHz: 1200,
		Price: 500,
		Architecture: "x86",
	}


	p1.PrintSummary()


	p1.GiveDiscount(450)


	p1.PrintSummary()
}


