package main

import "fmt"

func main(){
	bc:=NewBlockchain()


	bc.AddBlock("Send 1 BTC to han")
	bc.AddBlock("Send 2 more BTC to tuo")


	for _,block:=range bc.blocks{
		fmt.Printf("prev.hash:%x \n", block.PreBlockHash)
		fmt.Printf("data:%s \n", block.Data)
		fmt.Printf("hash:%x \n ",block.Hash)
		fmt.Println()

	}
}
