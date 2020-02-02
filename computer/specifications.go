package computer

import "fmt"

type ProcessingUnit struct{
	Maker string
	MHz int
	Price float64
	Architecture string
}


func (pUnit ProcessingUnit) PrintSummary(){
	fmt.Println(pUnit.Maker)
	fmt.Println(pUnit.MHz)
	fmt.Println(pUnit.Price)
	fmt.Println(pUnit.Architecture)
}


func (pUnit *ProcessingUnit) GiveDiscount(newPrice float64){
	pUnit.Price = newPrice
}





