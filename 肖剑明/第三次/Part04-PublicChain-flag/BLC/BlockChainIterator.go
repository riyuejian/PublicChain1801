package BLC

import (
	"math/big"
	"log"
)


type BlockChainIterator struct {
	CurrentHash []byte
	Db *BlockDatabase
}

//区块链迭代器，创建一个迭代器
func (bc BlockChain)CreateBlockChainIterator()*BlockChainIterator{

	return &BlockChainIterator{bc.Tip, bc.Db}
}

//迭代器next函数
func (bcI *BlockChainIterator)Next()*Block{
	var hashInt big.Int
	hashInt.SetBytes(bcI.CurrentHash)
	if hashInt.Cmp(big.NewInt(0)) == 0{
		return nil
	}
	err, block := bcI.Db.Select(BlockChainTableName, bcI.CurrentHash)
	bcI.CurrentHash = block.PrevHash
	if err != nil {
		log.Panic(err, "BlockIterator Next func err")
	}
	return block
}

//使用迭代器打印区块链
func (bc BlockChain)PrintBloackChain(){
	bcI := bc.CreateBlockChainIterator()
	for{
		block := bcI.Next()
		if block == nil{
			return
		}
		block.GetBlockInfo()
	}
}