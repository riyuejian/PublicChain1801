package BLC

import "fmt"

type BlockChain struct{
	Blocks []*Block
}

// 1、创建区块链
func CreateBlockChain() *BlockChain{
	//创建创世区块
	gensisBlock := CreateGenenisBlock()
	//创建区块链对象
	//bc := BlockChain{[]*Block{gensisBlock}}
	return &BlockChain{[]*Block{gensisBlock}}
	//添加创世区块到区块链
	//bc.Blocks = append(bc.Blocks, gensisBlock)
	//return &bc
}

func (bc *BlockChain)AddNewBlock(data string){
	var currentHeight int64 = int64(len(bc.Blocks))
	block := NewBlock(data, currentHeight + 1, bc.Blocks[currentHeight -1].Hash)
	bc.Blocks = append(bc.Blocks, block)
}

func (bc *BlockChain)BlockInfo(height int64){
	if height <=0 || height > int64(len(bc.Blocks)){
		fmt.Printf("[]")
	}else{
		bc.Blocks[height-1].GetBlock()
	}

}