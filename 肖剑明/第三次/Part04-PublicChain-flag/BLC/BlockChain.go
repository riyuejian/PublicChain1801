package BLC

import (
	"math/big"
	"os"
	"fmt"
	"github.com/boltdb/bolt"
	"log"
)

type BlockChain struct {
	//Blocks []*Block
	Tip []byte //最新的区块hash
	CurrentHeight int64 //当前区块的高度
	Db *BlockDatabase
}


const BlockChainDBName string = "BlockChainDB.db"
const BlockChainTableName string = "Blocks"


//判断区块链数据库是否存在
func dbExist()bool{
	if _, err := os.Stat(BlockChainDBName); os.IsNotExist(err){
		return false
	}
	return true
}
//创建区块链
func CreateBlockChain()*BlockChain{
	var blockChain BlockChain
	if dbExist(){
		fmt.Println("创始区块已经存在......")
		db,err := bolt.Open(BlockChainDBName,0600, nil)
		defer db.Close()
		if err != nil{
			log.Panic(err)
		}
		var blockChain BlockChain
		err = db.View(func(tx *bolt.Tx) error {
			b := tx.Bucket([]byte(BlockChainTableName))
			if b!=nil{
				hash := b.Get([]byte("AA"))
				blockChain.Tip = hash
				//
				blockChain.Db = &BlockDatabase{BlockChainDBName,db}
			}
			return nil
		})
		return &blockChain
	}
	blockChain.Db = CreateBlockDatabase(BlockChainDBName)
	blockChain.Db.CreateTable(BlockChainTableName)
	gensisBlock := CreateGenesisBlock()

	blockChain.Tip = gensisBlock.Hash
	blockChain.CurrentHeight = 1
	blockChain.Db.Insert(BlockChainTableName, gensisBlock.Hash, gensisBlock.Serialize())
	return &blockChain
}

//添加新的区块
func (bc *BlockChain)AddNewBlock(data string){
	block := CreateNewBlock(data, bc.CurrentHeight + 1, bc.Tip)
	bc.Tip = block.Hash
	bc.CurrentHeight++
	bc.Db.Insert(BlockChainTableName, block.Hash, block.Serialize())
}

//遍历区块链,输出所有区块信息
func (bc *BlockChain)BlockIterator(){
	target := big.NewInt(0)
	var hashInt big.Int
	hash := bc.Tip
	for {
		hashInt.SetBytes(hash)
		//创始区块，就直接退出
		if hashInt.Cmp(target) == 0{
			return
		}
		_, block := bc.Db.Select(BlockChainTableName,hash)
		block.GetBlockInfo()
		hash = block.PrevHash
	}
}
