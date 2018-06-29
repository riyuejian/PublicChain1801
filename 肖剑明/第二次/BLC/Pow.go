package BLC

import (
	"math/big"
	"fmt"
	"strconv"
	"bytes"
	"crypto/sha256"
)
//有效hash值的前导0个数
const TargetBit  = 16

type ProofOfWork struct{
	block *Block
	target *big.Int
}

func NewProofOfWork(block *Block) *ProofOfWork{
	//创建一个大数1
	bigNum := big.NewInt(1)
	//左移构造前导targetbit/4个0的目标难度值
	bigNum.Lsh(bigNum, uint(256 - TargetBit))
	//fmt.Printf("Bignum:0x%64X\n",bigNum)
	return  &ProofOfWork{block, bigNum}
}

func (pow *ProofOfWork)PrepareData(nonce int64)[]byte  {
	heigthBytes := IntToHex(pow.block.Height)
	timestring := strconv.FormatInt(pow.block.Timestamp, 2)
	timebytes := []byte(timestring)
	data := bytes.Join([][]byte{
		timebytes,
		heigthBytes,
		pow.block.PrevBlockHash,
		pow.block.Data,
		IntToHex(nonce)},
		[]byte{})
	return data
}

func (pow *ProofOfWork)Run()(int64,[]byte){
	var hashInt big.Int
	var nonce int64 = 0
	var hash [32]byte
	for{

		hash = sha256.Sum256(pow.PrepareData(nonce)) //计算区块的hash值
		hashInt.SetBytes(hash[:])
		// Cmp compares x and y and returns:
		//   -1 if x <  y
		//    0 if x == y
		//   +1 if x >  y
		if hashInt.Cmp(pow.target) == -1{
			fmt.Printf("\r0x%x", hash)
			break
		}else{
			fmt.Printf("\r0x%x", hash)
		}
		nonce++
	}
	fmt.Printf("\n")
	return nonce, hash[:]
}
