package BLC

import (
	"time"
	"fmt"
	"strconv"
	"bytes"
	"crypto/sha256"
)

type Block struct {
	//时间戳，创建区块的时间
	Timestamp int64
	//区块高度
	Height int64
	//上一个区块的hash
	PrevBlockHash []byte
	//交易数据包
	Data []byte
	//Hash，当前区块的hash
	Hash []byte
	//Nonce 随机数
	Nonce int64
}
func (block *Block)GetBlock(){
	fmt.Println("[")
	fmt.Println("Block Height:", block.Height)
	fmt.Println("Transaction data:", string(block.Data))
	fmt.Println("Timestamp:", block.Timestamp)
	fmt.Printf("Previous Block Hash:0x%x\n", block.PrevBlockHash)
	fmt.Printf("Hash               :0x%x\n", block.Hash)
	fmt.Printf("Nonce:%d\n", block.Nonce)
	fmt.Println("]")
}

func (block *Block)SetHash() {
	//Height转化为字节数组 []byte
	heigthBytes := IntToHex(block.Height)
	//fmt.Println(heigthBytes)
	//将时间戳转换为字节数组
	timestring := strconv.FormatInt(block.Timestamp, 2)
	timebytes := []byte(timestring)
	//fmt.Println(timebytes)
	//将数据转化为字节数组

	//将所有的数据拼接起来
	blockBytes := bytes.Join([][]byte{timebytes,heigthBytes, block.PrevBlockHash,block.Data, block.Hash}, []byte{})
	//fmt.Println(blockBytes)
	//生成Hash
	hash := sha256.Sum256(blockBytes)
	block.Hash = hash[:]
	//fmt.Printf("hash:%x", hash)
}
//crete new blcok
func NewBlock(data string, height int64, prevBlockHash []byte) *Block{
	//创建区块
	block := &Block{
		time.Now().Unix(),
		height,
		prevBlockHash,
		[]byte(data),
		nil,
		0}
	pow := NewProofOfWork(block)
	block.Nonce, block.Hash = pow.Run()
	//设置当前区块的hash
	//block.SetHash()
	//fmt.Printf("\r\nHash:0x%x", block.Hash)
	return block
}

func CreateGenenisBlock() *Block{
	block := NewBlock("Genenis block...",1, []byte{0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0} )
	//block.SetHash()
	//fmt.Printf("Hash:%x", block.Hash)
	return block
}