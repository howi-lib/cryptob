package main

import (
	"flag"
	"fmt"

	"github.com/howi-lib/cryptob"
)

var address string

func init() {
	flag.StringVar(&address, "address", "1A1zP1eP5QGefi2DMPTfTL5SLmv7DivfNa", "Your BTC address")

}
func main() {
	flag.Parse()
	symbol := "BTC"

	wallet, err := cryptob.Load(symbol, address)

	if err != nil {
		panic(err)
	}

	fmt.Println("Address: ", address)
	fmt.Println("Balance is: ", wallet.Balance())
}
