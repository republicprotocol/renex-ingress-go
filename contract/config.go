package contract

import (
	republicContract "github.com/republicprotocol/republic-go/contract"
	"github.com/republicprotocol/republic-go/identity"
)

// Network is used to represent a Republic Protocol network.
type Network string

const (
	// NetworkMainnet represents the mainnet
	NetworkMainnet Network = "mainnet"
	// NetworkTestnet represents the internal F∅ testnet
	NetworkTestnet Network = "testnet"
	// NetworkNightly represents the internal Nightly testnet
	NetworkNightly Network = "nightly"
	// NetworkLocal represents a local network
	NetworkLocal Network = "local"
)

type Config struct {
	Ethereum                republicContract.Config `json:"republic"`
	RenExEthereum           RenExConfig             `json:"renex"`
	BootstrapMultiAddresses identity.MultiAddresses `json:"bootstrapMultiAddresses"`
}

// RenExConfig defines the different settings for connecting to Ethereum on
// different Republic Protocol networks.
type RenExConfig struct {
	Network                    Network `json:"network"`
	URI                        string  `json:"uri"`
	RenExBrokerVerifierAddress string  `json:"renExBrokerVerifier"`
	OrderbookAddress           string  `json:"orderbook"`
	WyreAddress                string  `json:"wyre"`
}
