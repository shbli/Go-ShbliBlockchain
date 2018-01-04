package main

import (
	"github.com/ShbliBlockchain/Go-ShbliBlockchain/Blockchain"
	"github.com/ShbliBlockchain/Go-ShbliBlockchain/CLI"
)

func main() {
	bc := Blockchain.NewBlockchain()

	cli := CLI.NewCli(bc)
	cli.Run()
	//let's close the DB once done
	bc.CloseDB()
}