package main

import (
	"fmt"

	"github.com/howi-lib/cryptob"
)

func main() {
	address := "1A1zP1eP5QGefi2DMPTfTL5SLmv7DivfNa"
	symbol := "BTC"

	wallet, err := cryptob.Load(symbol, address)

	if err != nil {
		panic(err)
	}

	fmt.Println("Address: ", address)
	fmt.Println("Balance is: ", wallet.Balance())
}
