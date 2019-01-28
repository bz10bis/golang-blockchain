package blockchain

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"fmt"
)

type TxOutput struct {
	Value  int
	PubKey string
}

type TxInput struct {
	ID  []byte
	Out int
	Sig string
}

type Transaction struct {
	ID     []byte
	Inputs []TxInput
	Ouputs []TxOutput
}

// CoinbaseTx is the only reference in the genesis block it set the coinbase of
// the blockchain and the reward for each mined block
func CoinbaseTx(to, data string) *Transaction {
	if data == "" {
		data = fmt.Sprintf("Coin to %s", to)
	}
	// The in tx has np previous tx, no output, and signature is just the data
	txin := TxInput{[]byte{}, -1, data}
	// The out tx has the reward and to is the address
	txout := TxOutput{100, to}

	tx := Transaction{nil, []TxInput{txin}, []TxOutput{txout}}
	tx.SetID()

	return &tx
}

// Methods

// SetID take the transaction and create a hash with it
func (tx *Transaction) SetID() {
	var encoded bytes.Buffer
	var hash [32]byte
	encoder := gob.NewEncoder(&encoded)
	err := encoder.Encode(tx)
	Handle(err)

	hash = sha256.Sum256(encoded.Bytes())
	tx.ID = hash[:]
}
