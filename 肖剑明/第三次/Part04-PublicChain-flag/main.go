package main

import (
	"flag"
	"fmt"
	"os"
	"log"
	"./BLC"
)
func main() {
	bc := BLC.CreateBlockChain()
	cli := BLC.CLI{bc}

	cli.Run()
}


func OsArgsTest(){
	addBlockCmd := flag.NewFlagSet("addblock", flag.ExitOnError)
	printblockCmd := flag.NewFlagSet("printblock", flag.ExitOnError)

	flagAddBlockData := addBlockCmd.String("data","http://github.com","交易数据")
	BLC.IsValidArgs()
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
		BLC.PrintUsage()
	}
	if addBlockCmd.Parsed(){
		if *flagAddBlockData == ""{
			BLC.PrintUsage()
			os.Exit(1)
		}else{
			fmt.Println(*flagAddBlockData)
		}
	}
	if printblockCmd.Parsed(){
		fmt.Println("输出所有数据......")
	}

}

func osArgsTest(){
	args := os.Args

	fmt.Printf("%v", args)
}

func flagTest(){
	flagString := flag.String("printChain", "", "打印区块链")
	flagInt := flag.Int("number", 0, "输入一个整数")
	flagBool := flag.Bool("open", false, "判断真假")
	flag.Parse()
	fmt.Printf("%s\n", *flagString)
	fmt.Printf("%d\n", *flagInt)
	fmt.Printf("%v\n", *flagBool)
}