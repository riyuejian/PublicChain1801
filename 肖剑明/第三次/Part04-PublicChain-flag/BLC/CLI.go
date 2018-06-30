package BLC

import (
	"os"
	"fmt"
	"log"
	"flag"
)

type CLI struct{
	BlockChain *BlockChain
}

//创建命令行对象
func CreateCli()*CLI{
	return &CLI{CreateBlockChain()}
}

func (cli *CLI)Run(){
	IsValidArgs()

	addBlockCmd := flag.NewFlagSet("addblock", flag.ExitOnError)
	printblockCmd := flag.NewFlagSet("printblock", flag.ExitOnError)

	flagAddBlockData := addBlockCmd.String("data","http://github.com","交易数据")

	switch os.Args[1]{
	case "addblock":
		err := addBlockCmd.Parse(os.Args[2:])
		if err != nil{
			log.Panic(err)
		}
	case "printblock":
		err := printblockCmd.Parse(os.Args[2:])
		if err != nil{
			log.Panic(err)
		}
	default:
		PrintUsage()
	}
	if addBlockCmd.Parsed(){
		if *flagAddBlockData == ""{
			PrintUsage()
			os.Exit(1)
		}else{
			//fmt.Println(*flagAddBlockData)
			cli.BlockChain.AddNewBlock(*flagAddBlockData)
		}
	}
	if printblockCmd.Parsed(){
		//fmt.Println("输出所有数据......")
		cli.BlockChain.PrintBloackChain()
	}
}
func IsValidArgs(){
	if len(os.Args) <2{
		PrintUsage()
		os.Exit(1)
	}
}

func PrintUsage(){
	fmt.Println("Usage:")
	fmt.Println("\taddblock -data DATA --交易数据")
	fmt.Println("\tprintblock --输出区块链")
}