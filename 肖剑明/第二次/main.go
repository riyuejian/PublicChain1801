package main

import (
	"./BLC"
	"fmt"
)
func main(){
	bc := BLC.CreateBlockChain()
	bc.AddNewBlock("A send 10 btc to B")
	bc.AddNewBlock("A send 5 eth to C")
	bc.AddNewBlock("A send 2 eth to d")
	bc.AddNewBlock("A send 1 eth to E")
	//bc.AddNewBlock("A send 3 eth to F")
	bc.BlockInfo(1)
	bc.BlockInfo(2)
	bc.BlockInfo(3)
	bc.BlockInfo(4)
	bc.BlockInfo(5)

	fmt.Println("Current Block Height:", len(bc.Blocks))

}