package cryptob

// Wallet interface provides API abstraction for Wallet
// on selected blockchain
type Wallet interface {
	// Symbol returns symbol of related crypto
	Symbol() string
	// Name returns full name of that crypto asset
	Name() string
	// Balance returns amount of crypto in this wallet
	Balance() float64
	// Address returns address in that blockchain
	Address() string
}
