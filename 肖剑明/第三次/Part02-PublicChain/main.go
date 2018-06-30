package main

import (
	"github.com/boltdb/bolt"
	"log"
	"fmt"
	"./DB"
	"./BLC"
)
func main(){
	bc := BLC.CreateBlockChain()
	bc.AddNewBlock("A send 10 BTM to B")
	bc.AddNewBlock("A send 1 BTC to C")

	db := DB.CreateBlockDatabase("BlockDB.db")
	db.CreateTable("Blocks")

	db.Insert("Blocks", bc.Blocks[0].Hash, bc.Blocks[0].Serialize())
	db.Insert("Blocks", bc.Blocks[1].Hash, bc.Blocks[1].Serialize())
	db.Insert("Blocks", bc.Blocks[2].Hash, bc.Blocks[2].Serialize())
	_, block:= db.Select("Blocks", bc.Blocks[0].Hash)
	block.GetBlockInfo()
	_, block = db.Select("Blocks", bc.Blocks[1].Hash)
	block.GetBlockInfo()
	_, block = db.Select("Blocks", bc.Blocks[2].Hash)
	block.GetBlockInfo()
	//CreateBoltdbTest()


}

func dbTest(){
	db := DB.CreateDB("block.db")
	//fmt.Println(db.Name)
	//db.CreateTable("hashTable")
	//fmt.Println("Create table success")
	//db.Insert("hashTable", DB.DbData{[]byte("A"), []byte("A send 10TC to B....")})
	//fmt.Println("insert table success")
	//db.Insert("hashTable", DB.DbData{[]byte("B"), []byte("A send 10BTC to C....")})
	//db.Select("hashTable","A")
	//db.Select("hashTable","B")
	//db.Insert("hashTable", DB.DbData{[]byte("C"), []byte("B send 0BTC to C....")})
	//db.Insert("hashTable", DB.DbData{[]byte("D"), []byte("C send 0BTC to D....")})
	db.Select("hashTable","A")
	db.Select("hashTable","B")
	db.Select("hashTable","C")
	db.Select("hashTable","D")
	db.Select("hashTable","E")
}

func CreateBoltdbTest(){
	// Open the my.db data file in your current directory.
	// It will be created if it doesn't exist.
	db, err := bolt.Open("my.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.Update(func(tx *bolt.Tx) error {

		bBucket, err := tx.CreateBucket([]byte("blockBucket"))
		if err != nil {
			return fmt.Errorf("create bucket: %s", err)
		}
		if bBucket != nil{
			err := bBucket.Put([]byte("A"), []byte("A send 10BTC to B......."))
			if err != nil{
				log.Panic("插入数据失败")
			}
		}
		return nil
	})

	db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("blockBucket"))
		v := b.Get([]byte("A"))
		fmt.Printf("The answer is: %s\n", v)
		return nil
	})
}