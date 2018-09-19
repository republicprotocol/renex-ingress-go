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
		case NetworkTestnet:
			config.URI = "https://kovan.infura.io"
		case NetworkFalcon:
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
		case NetworkTestnet:
			config.RenExBrokerVerifierAddress = "0x5BF19a6Ea8631BB722ade58e0D2C5813740c88fD"
			config.WyreAddress = "0xB14fA2276D8bD26713A6D98871b2d63Da9eefE6f"
		case NetworkFalcon:
			config.RenExBrokerVerifierAddress = "0xb6A95aED1588bE477981dcdEacd13776570ecB3D"
			config.WyreAddress = "0xB14fA2276D8bD26713A6D98871b2d63Da9eefE6f"
		case NetworkNightly:
			config.RenExBrokerVerifierAddress = "0xcf2F6b4b698Cd6a6B3eb1d874a939742d15f8e7E"
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
