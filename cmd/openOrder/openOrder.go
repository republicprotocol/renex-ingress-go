package main

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/republicprotocol/renex-ingress-go/httpadapter"
	"github.com/republicprotocol/republic-go/contract"
	"github.com/republicprotocol/republic-go/crypto"
	"github.com/republicprotocol/republic-go/identity"
	"github.com/republicprotocol/republic-go/order"
	"github.com/republicprotocol/republic-go/registry"
)

type config struct {
	Ethereum                contract.Config         `json:"ethereum"`
	BootstrapMultiAddresses identity.MultiAddresses `json:"bootstrapMultiAddresses"`
}

func main() {
	ingressParam := flag.String("ingress", "", "Ingress API endpoint")
	configParam := flag.String("config", "", "Configuration file")
	keystoreParam := flag.String("keystore", "", "Keystore file")
	passphraseParam := flag.String("passphrase", "", "Optional passphrase to decrypt the keystore file")

	flag.Parse()

	if *ingressParam == "" {
		log.Fatalf("missing ingress api endpoint")
	}
	keystore, err := loadKeystore(*keystoreParam, *passphraseParam)
	if err != nil {
		log.Fatalf("cannot load keystore: %v", err)
	}
	conf, err := loadConfig(*configParam)
	if err != nil {
		log.Fatalf("cannot load config: %v", err)
	}
	contractBinder, err := loadContractBinder(conf, keystore)
	if err != nil {
		log.Fatalf("cannot load smart contract: %v", err)
	}

	onePrice := order.CoExp{
		Co:  2,
		Exp: 40,
	}
	oneVol := order.CoExp{
		Co:  5,
		Exp: 12,
	}
	buy := order.NewOrder(order.TypeLimit, order.ParityBuy, order.SettlementRenEx, time.Now().Add(1*time.Hour), order.TokensDGXREN, onePrice, oneVol, oneVol, rand.Uint64())
	sell := order.NewOrder(order.TypeLimit, order.ParitySell, order.SettlementRenEx, time.Now().Add(1*time.Hour), order.TokensDGXREN, onePrice, oneVol, oneVol, rand.Uint64())
	ords := []order.Order{buy, sell}

	for _, ord := range ords {
		message := append([]byte("Republic Protocol: open: "), ord.ID[:]...)
		signatureData := crypto.Keccak256([]byte(fmt.Sprintf("\x19Ethereum Signed Message:\n%d", len(message))), message)
		signature, err := keystore.Sign(signatureData)
		if err != nil {
			log.Fatalf("cannot sign order: %v", err)
		}

		log.Printf("order = %v, from = %v", ord.ID, contractBinder.From().String())
		request := httpadapter.OpenOrderRequest{
			Signature:             base64.StdEncoding.EncodeToString(signature),
			OrderFragmentMappings: httpadapter.OrderFragmentMappings{},
		}

		// Current epoch
		pods, err := contractBinder.Pods()
		if err != nil {
			log.Fatalf("cannot get pods from darkpool: %v", err)
		}
		request.OrderFragmentMappings = append(request.OrderFragmentMappings, generateOrderFragmentMapping(&contractBinder, ord, pods, 0))
		// Previous epoch
		pods, err = contractBinder.PreviousPods()
		if err != nil {
			log.Fatalf("cannot get previous pods from darkpool: %v", err)
		}
		request.OrderFragmentMappings = append(request.OrderFragmentMappings, generateOrderFragmentMapping(&contractBinder, ord, pods, 1))

		data, err := json.MarshalIndent(request, "", "  ")
		if err != nil {
			log.Fatalf("cannot marshal request: %v", err)
		}
		buf := bytes.NewBuffer(data)

		log.Printf("sending to %v", fmt.Sprintf("%v/orders", *ingressParam))
		res, err := http.DefaultClient.Post(fmt.Sprintf("%v/orders", *ingressParam), "application/json", buf)
		if err != nil {
			log.Fatalf("cannot send request: %v", err)
		}
		defer res.Body.Close()

		log.Printf("status: %v", res.StatusCode)
		resText, err := ioutil.ReadAll(res.Body)
		if err != nil {
			log.Fatalf("cannot read response body: %v", err)
		}
		log.Printf("body: %v", string(resText))
	}
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

func loadConfig(configFile string) (config, error) {
	file, err := os.Open(configFile)
	if err != nil {
		return config{}, err
	}
	defer file.Close()
	conf := config{}
	if err := json.NewDecoder(file).Decode(&conf); err != nil {
		return config{}, err
	}
	return conf, nil
}

func loadContractBinder(conf config, keystore crypto.Keystore) (contract.Binder, error) {
	conn, err := contract.Connect(conf.Ethereum)
	if err != nil {
		fmt.Println(fmt.Errorf("cannot connect to ethereum: %v", err))
		return contract.Binder{}, err
	}

	auth := bind.NewKeyedTransactor(keystore.EcdsaKey.PrivateKey)

	return contract.NewBinder(auth, conn)
}

func generateOrderFragmentMapping(contractBinder *contract.Binder, ord order.Order, pods []registry.Pod, depth int) httpadapter.OrderFragmentMapping {
	orderFragmentMapping := httpadapter.OrderFragmentMapping{}

	for _, pod := range pods {
		n := int64(len(pod.Darknodes))
		k := int64(2 * (len(pod.Darknodes) + 1) / 3)
		hash := base64.StdEncoding.EncodeToString(pod.Hash[:])
		ordFragments, err := ord.Split(n, k)
		if err != nil {
			log.Fatalf("cannot split order to %v: %v", hash, err)
		}
		orderFragmentMapping[hash] = []httpadapter.OrderFragment{}
		for i, ordFragment := range ordFragments {
			marshaledOrdFragment := httpadapter.OrderFragment{
				Index: int64(i + 1),
			}

			pubKey, err := contractBinder.PublicKey(pod.Darknodes[i])
			if err != nil {
				log.Fatalf("cannot get public key of %v: %v", pod.Darknodes[i], err)
			}

			encryptedFragment, err := ordFragment.Encrypt(pubKey)
			marshaledOrdFragment.ID = base64.StdEncoding.EncodeToString(encryptedFragment.ID[:])
			marshaledOrdFragment.EpochDepth = int32(depth)
			marshaledOrdFragment.OrderID = base64.StdEncoding.EncodeToString(encryptedFragment.OrderID[:])
			marshaledOrdFragment.OrderParity = encryptedFragment.OrderParity
			marshaledOrdFragment.OrderSettlement = encryptedFragment.OrderSettlement
			marshaledOrdFragment.OrderType = encryptedFragment.OrderType
			marshaledOrdFragment.OrderExpiry = encryptedFragment.OrderExpiry.Unix()
			marshaledOrdFragment.Tokens = base64.StdEncoding.EncodeToString(encryptedFragment.Tokens)
			marshaledOrdFragment.Price = []string{
				base64.StdEncoding.EncodeToString(encryptedFragment.Price.Co),
				base64.StdEncoding.EncodeToString(encryptedFragment.Price.Exp),
			}
			marshaledOrdFragment.Volume = []string{
				base64.StdEncoding.EncodeToString(encryptedFragment.Volume.Co),
				base64.StdEncoding.EncodeToString(encryptedFragment.Volume.Exp),
			}
			marshaledOrdFragment.MinimumVolume = []string{
				base64.StdEncoding.EncodeToString(encryptedFragment.MinimumVolume.Co),
				base64.StdEncoding.EncodeToString(encryptedFragment.MinimumVolume.Exp),
			}
			marshaledOrdFragment.Nonce = base64.StdEncoding.EncodeToString(encryptedFragment.Nonce)
			orderFragmentMapping[hash] = append(orderFragmentMapping[hash], marshaledOrdFragment)
		}
	}

	return orderFragmentMapping
}
