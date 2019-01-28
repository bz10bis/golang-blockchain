package wallet

import (
	"bytes"
	"crypto/elliptic"
	"encoding/gob"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

const walletFile = "./db/wallet.data"

type Wallets struct {
	Wallets map[string]*Wallet
}

// Create Wallets from our file
func CreateWallets() (*Wallets, error) {
	wallets := Wallets{}
	wallets.Wallets = make(map[string]*Wallet)
	err := wallets.LoadFile()
	return &wallets, err
}

// Methods

// SaveFile save wallets data on the disk
func (ws *Wallets) SaveFile() {
	var content bytes.Buffer
	gob.Register(elliptic.P256())

	encoder := gob.NewEncoder(&content)
	err := encoder.Encode(ws)
	if err != nil {
		log.Panic(err)
	}

	// write content into walletFile and give him permission
	err = ioutil.WriteFile(walletFile, content.Bytes(), 0644)
}

// LoadFile load wallets from file to memory
func (ws *Wallets) LoadFile() error {
	if _, err := os.Stat(walletFile); os.IsNotExist(err) {
		return err
	}
	var wallets Wallets
	fileContent, err := ioutil.ReadFile(walletFile)
	if err != nil {
		return err
	}
	gob.Register(elliptic.P256())
	decoder := gob.NewDecoder(bytes.NewReader(fileContent))
	err = decoder.Decode(&wallets)
	if err != nil {
		return err
	}
	ws.Wallets = wallets.Wallets
	return nil
}

// GetWallet retrieve wallet from specific addres
func (ws Wallets) GetWallet(address string) Wallet {
	return *ws.Wallets[address]
}

func (ws Wallets) GetAllAddresses() []string {
	var allAddresses []string
	for address := range ws.Wallets {
		allAddresses = append(allAddresses, address)
	}
	return allAddresses
}

func (ws Wallets) AddWallet() string {
	wallet := MakeWallet()
	address := fmt.Sprintf("%s", wallet.Address())
	ws.Wallets[address] = wallet
	return address
}
