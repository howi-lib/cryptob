package main

import (
	"flag"
	"fmt"

	"github.com/howi-lib/cryptob"
)

var (
	address string
	symbol  string
)

func init() {
	flag.StringVar(&address, "address", "1A1zP1eP5QGefi2DMPTfTL5SLmv7DivfNa", "Your BTC address")
	flag.StringVar(&symbol, "symbol", "BTC", "Currency symbol")

}
func main() {
	flag.Parse()

	wallet, err := cryptob.Load(symbol, address)

	if err != nil {
		panic(err)
	}

	fmt.Println("Symbol: ", symbol)
	fmt.Println("Address: ", address)
	fmt.Println("Balance is: ", wallet.Balance())
}
