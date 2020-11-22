package cryptob

import (
	"strings"

	"github.com/howi-lib/cryptob/crypto/btc"
)

// Load wallet for symbol and address
func Load(symbol, address string) (Wallet, error) {
	symbol = strings.ToUpper(symbol)
	switch symbol {
	case "BTC":
		return btc.Load(address), nil
	}
	return nil, nil
}
