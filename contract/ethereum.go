package contract

import (
	"fmt"

	"github.com/ethereum/go-ethereum/ethclient"
	ethrpc "github.com/ethereum/go-ethereum/rpc"
)

// Conn contains the client and the contracts deployed to it
type Conn struct {
	RawClient *ethrpc.Client
	Client    *ethclient.Client
	Config    Config
}

// Connect to a URI.
func Connect(config Config) (Conn, error) {
	if config.URI == "" {
		switch config.Network {
		case NetworkMainnet:
			config.URI = "https://mainnet.infura.io"
		case NetworkTestnet:
			config.URI = "https://kovan.infura.io"
		case NetworkNightly:
			config.URI = "https://kovan.infura.io"
		case NetworkLocal:
			config.URI = "http://localhost:8545"
		default:
			return Conn{}, fmt.Errorf("cannot connect to %s: unsupported", config.Network)
		}
	}

	if config.RenExBrokerVerifierAddress == "" {
		switch config.Network {
		case NetworkMainnet:
			config.RenExBrokerVerifierAddress = "0x31a0d1a199631d244761eeba67e8501296d2e383"
			config.OrderbookAddress = "0x6b8bb175c092de7d81860b18db360b734a2598e0"
			config.WyreAddress = "0x9f2a7b5e6280727cd6c8486f5f96e5f76164f2df"
		case NetworkTestnet:
			config.RenExBrokerVerifierAddress = "0x5BF19a6Ea8631BB722ade58e0D2C5813740c88fD"
			config.OrderbookAddress = "0xA53Da4093c682a4259DE38302341BFEf7e9f7a4f"
			config.WyreAddress = "0xB14fA2276D8bD26713A6D98871b2d63Da9eefE6f"
		case NetworkNightly:
			config.RenExBrokerVerifierAddress = "0xcf2F6b4b698Cd6a6B3eb1d874a939742d15f8e7E"
			config.OrderbookAddress = "0x376127aDc18260fc238eBFB6626b2F4B59eC9b66"
			config.WyreAddress = "0xB14fA2276D8bD26713A6D98871b2d63Da9eefE6f"
		case NetworkLocal:
		default:
			return Conn{}, fmt.Errorf("no default contract address on %s", config.Network)
		}
	}

	ethclient, err := ethclient.Dial(config.URI)
	if err != nil {
		return Conn{}, err
	}

	return Conn{
		Client: ethclient,
		Config: config,
	}, nil
}
