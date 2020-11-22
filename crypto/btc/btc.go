package btc

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type response struct {
	Balance float64 `json:"balance"`
}

const (
	// SYMBOL for BTC
	SYMBOL = "BTC"
	// NAME for BTC
	NAME = "Bitcoin"
)

// Load BTC Address
func Load(address string) BTC {
	return BTC{
		address: address,
	}
}

// BTC struct
type BTC struct {
	address string
}

// Symbol returns symbol of related crypto
func (btc BTC) Symbol() string {
	return SYMBOL
}

// Name returns full name of that crypto asset
func (btc BTC) Name() string {
	return NAME
}

// Balance returns amount of crypto in this wallet
func (btc BTC) Balance() float64 {
	// DEMO
	resp, err := http.Get("https://api-r.bitcoinchain.com/v1/address/" + btc.Address())
	if err != nil {
		return 0.00
	}
	defer resp.Body.Close()
	var response []response
	decoder := json.NewDecoder(resp.Body)
	if err := decoder.Decode(&response); err != nil {
		fmt.Println(err)
		return 0.00
	}
	return response[0].Balance
}

// Address returns address in that blockchain
func (btc BTC) Address() string {
	return btc.address
}
