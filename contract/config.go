package contract

// Network is used to represent a Republic Protocol network.
type Network string

const (
	// NetworkMainnet represents the mainnet
	NetworkMainnet Network = "mainnet"
	// NetworkTestnet represents the internal F∅ testnet
	NetworkTestnet Network = "testnet"
	// NetworkFalcon represents the internal Falcon testnet
	NetworkFalcon Network = "falcon"
	// NetworkNightly represents the internal Nightly testnet
	NetworkNightly Network = "nightly"
	// NetworkLocal represents a local network
	NetworkLocal Network = "local"
)

// Config defines the different settings for connecting to Ethereum on
// different Republic Protocol networks.
type Config struct {
	Network                    Network `json:"network"`
	URI                        string  `json:"uri"`
	RenExBrokerVerifierAddress string  `json:"renExBrokerVerifier"`
	OrderbookAddress           string  `json:"orderbook"`
	WyreAddress                string  `json:"wyre"`
}
