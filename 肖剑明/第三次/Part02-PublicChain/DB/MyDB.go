package DB

import (
	"github.com/boltdb/bolt"
	"log"
	"fmt"
	"../BLC"
)

type DataBaseObject struct{
	Name string
	Db *bolt.DB
}
type DbData struct{
	Key []byte
	Value []byte
}
//创建数据库
func CreateDB(dbName string) *DataBaseObject{

	db, err := bolt.Open(dbName, 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	return &DataBaseObject{dbName, db}
}

//打开数据库
func (db *DataBaseObject)Open()error{
	mydb, err := bolt.Open(db.Name, 0600, nil)
	if err != nil {
		log.Fatal(err)
		fmt.Println("Open database", db.Name, "failed...")
	}
	db.Db = mydb
	return err
}

//创建表
func (db *DataBaseObject)CreateTable(name string)error{
	err := db.Open()
	defer db.Db.Close()
	if err != nil{
		log.Panic("CreateTable: Open database failed...")
		return err
	}
	err = db.Db.Update(func(tx *bolt.Tx) error {

		_, err := tx.CreateBucket([]byte(name))
		if err != nil {
			return fmt.Errorf("create bucket: %s", err)
		}
		return nil
	})
	return nil
}

//插入数据
func (db *DataBaseObject)Insert(tableName string, data DbData)error{
	err := db.Open()
	defer db.Db.Close()
	if err != nil{
		log.Panic("Insert: Open database failed...")
		return err
	}
	err = db.Db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(tableName))
		if bucket !=nil{
			err := bucket.Put(data.Key, data.Value)
			if err != nil{
				log.Panic("Insert data failed....")
			}
		}
		return nil
	})
	return nil
}

//查询数据
func (db *DataBaseObject) Select(tableName string, key string)(error, *DbData) {
	db.Open()
	var dbData DbData
	defer db.Db.Close()
	db.Db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(tableName))
		dbData.Key = []byte(key)
		dbData.Value = b.Get(dbData.Key)
		fmt.Printf("\nThe answer is: %v\n", dbData.Value)
		BLC.DeSerialize(dbData.Value).GetBlockInfo()
		return nil
	})
	return nil, &dbData
}

