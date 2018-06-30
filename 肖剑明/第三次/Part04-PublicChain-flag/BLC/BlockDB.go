package BLC

import (
	"github.com/boltdb/bolt"
	"log"
	"fmt"
)

type BlockDatabase struct{
	Name string
	BlockDb *bolt.DB
}

//创建数据库
func CreateBlockDatabase(dbName string) *BlockDatabase{

	db, err := bolt.Open(dbName, 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	return &BlockDatabase{dbName, db}
}

//打开数据库
func (db *BlockDatabase)Open()error{
	tmpDb, err := bolt.Open(db.Name, 0600, nil)
	if err != nil {
		log.Fatal(err)
		fmt.Println("Open database", db.Name, "failed...")
	}
	db.BlockDb = tmpDb
	return err
}

//创建表
func (db *BlockDatabase)CreateTable(name string)error{
	err := db.Open()
	defer db.BlockDb.Close()
	if err != nil{
		log.Panic("CreateTable: Open database failed...")
		return err
	}
	err = db.BlockDb.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucket([]byte(name))
		if err != nil {
			return fmt.Errorf("create bucket: %s", err)
		}
		return nil
	})
	return nil
}

//插入数据
func (db *BlockDatabase)Insert(tableName string, key []byte, data []byte)error{
	err := db.Open()
	defer db.BlockDb.Close()
	if err != nil{
		log.Panic("Insert: Open database failed...")
		return err
	}
	err = db.BlockDb.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(tableName))
		if bucket !=nil{
			err := bucket.Put(key, data)
			if err != nil{
				log.Panic("Insert data failed....")
			}
			err = bucket.Put([]byte("AA"), key)
			if err != nil{
				log.Panic(err)
			}
		}
		return nil
	})
	return nil
}

//查询数据
func (db *BlockDatabase) Select(tableName string, key []byte)(error,*Block) {
	var block *Block
	db.Open()
	defer db.BlockDb.Close()
	err := db.BlockDb.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(tableName))
		data := b.Get(key)
		//fmt.Printf("\nThe answer is: %v\n", data)
		block = DeSerialize(data)
		return nil
	})
	return err,block
}
