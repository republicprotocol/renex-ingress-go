package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/republicprotocol/renex-ingress-go/httpadapter"
	"github.com/republicprotocol/renex-ingress-go/ingress"
	"github.com/republicprotocol/republic-go/contract"
	"github.com/republicprotocol/republic-go/crypto"
	"github.com/republicprotocol/republic-go/dht"
	"github.com/republicprotocol/republic-go/grpc"
	"github.com/republicprotocol/republic-go/identity"
	"github.com/republicprotocol/republic-go/logger"
	"github.com/republicprotocol/republic-go/swarm"
)

type config struct {
	Ethereum                contract.Config         `json:"ethereum"`
	BootstrapMultiAddresses identity.MultiAddresses `json:"bootstrapMultiAddresses"`
}

func main() {
	logger.SetFilterLevel(logger.LevelDebugLow)

	done := make(chan struct{})
	defer close(done)
	defer logger.Info("shutting down...")

	networkParam := os.Getenv("NETWORK")
	if networkParam == "" {
		log.Fatalf("cannot read network environment")
	}
	configParam := fmt.Sprintf("env/%v/config.json", networkParam)
	keystoreParam := fmt.Sprintf("env/%v/%v.keystore.json", networkParam, os.Getenv("DYNO"))
	keystorePassphraseParam := os.Getenv("KEYSTORE_PASSPHRASE")

	config, err := loadConfig(configParam)
	if err != nil {
		log.Fatalf("cannot load config: %v", err)
	}

	keystore, err := loadKeystore(keystoreParam, keystorePassphraseParam)
	if err != nil {
		log.Fatalf("cannot load keystore: %v", err)
	}

	multiAddr, err := getMultiaddress(keystore, os.Getenv("PORT"))
	if err != nil {
		log.Fatalf("cannot get multi-address: %v", err)
	}

	conn, err := contract.Connect(config.Ethereum)
	if err != nil {
		log.Fatalf("cannot connect to ethereum: %v", err)
	}
	auth := bind.NewKeyedTransactor(keystore.EcdsaKey.PrivateKey)
	binder, err := contract.NewBinder(auth, conn)
	if err != nil {
		log.Fatalf("cannot create contract binder: %v", err)
	}

	dht := dht.NewDHT(multiAddr.Address(), 20)
	swarmClient := grpc.NewSwarmClient(multiAddr)
	swarmer := swarm.NewSwarmer(swarmClient, &dht)
	orderbookClient := grpc.NewOrderbookClient()
	ingresser := ingress.NewIngress(&binder, swarmer, orderbookClient, time.Second)
	ingressAdapter := httpadapter.NewIngressAdapter(ingresser)

	go func() {
		ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
		defer cancel()
		if err := swarmer.Bootstrap(ctx, config.BootstrapMultiAddresses); err != nil {
			log.Printf("error bootstrapping: %v", err)
		}

		syncErrs := ingresser.Sync(done)
		go func() {
			for err := range syncErrs {
				logger.Error(fmt.Sprintf("error syncing: %v", err))
			}
		}()

		processErrs := ingresser.ProcessRequests(done)
		go func() {
			for err := range processErrs {
				logger.Error(fmt.Sprintf("error processing: %v", err))
			}
		}()
	}()

	log.Printf("address %v", multiAddr)
	log.Printf("ethereum %v", auth.From.Hex())
	for _, multiAddr := range dht.MultiAddresses() {
		log.Printf("  %v", multiAddr)
	}
	log.Printf("listening at 0.0.0.0:%v...", os.Getenv("PORT"))
	if err := http.ListenAndServe(fmt.Sprintf("0.0.0.0:%v", os.Getenv("PORT")), httpadapter.NewIngressServer(ingressAdapter)); err != nil {
		log.Fatalf("error listening and serving: %v", err)
	}
}

func getMultiaddress(keystore crypto.Keystore, port string) (identity.MultiAddress, error) {
	// Get our IP address
	ipInfoOut, err := exec.Command("curl", "https://ipinfo.io/ip").Output()
	if err != nil {
		return identity.MultiAddress{}, err
	}
	ipAddress := strings.Trim(string(ipInfoOut), "\n ")
	ingressMultiaddress, err := identity.NewMultiAddressFromString(fmt.Sprintf("/ip4/%s/tcp/%s/republic/%s", ipAddress, port, keystore.Address()))
	if err != nil {
		return identity.MultiAddress{}, fmt.Errorf("cannot obtain trader multi address %v", err)
	}
	return ingressMultiaddress, nil
}

func loadConfig(configFile string) (config, error) {
	file, err := os.Open(configFile)
	if err != nil {
		return config{}, err
	}
	defer file.Close()
	c := config{}
	if err := json.NewDecoder(file).Decode(&c); err != nil {
		return config{}, err
	}
	return c, nil
}

func loadKeystore(keystoreFile, passphrase string) (crypto.Keystore, error) {
	file, err := os.Open(keystoreFile)
	if err != nil {
		return crypto.Keystore{}, err
	}
	defer file.Close()

	if passphrase == "" {
		keystore := crypto.Keystore{}
		if err := json.NewDecoder(file).Decode(&keystore); err != nil {
			return keystore, err
		}
		return keystore, nil
	}

	keystore := crypto.Keystore{}
	keystoreData, err := ioutil.ReadAll(file)
	if err != nil {
		return keystore, err
	}
	if err := keystore.DecryptFromJSON(keystoreData, passphrase); err != nil {
		return keystore, err
	}
	return keystore, nil
}
