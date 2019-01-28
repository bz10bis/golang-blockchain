package blockchain

import (
	"bytes"
	"encoding/gob"
	"log"
)

// Block represent a single block on the blockchain
// it contain the previous block hash, the data inside the block
// and the current Hash wich is a combination of current data and previous hash
type Block struct {
	Hash     []byte
	Data     []byte
	PrevHash []byte
	Nonce    int
}

// Functions

// Handle is a small function that print out the error
func Handle(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// CreateBlock create a new block and return a pointer to that block
func CreateBlock(data string, prevHash []byte) *Block {
	block := &Block{[]byte{}, []byte(data), prevHash, 0}
	// block.DeriveHash()
	pow := NewProof(block)
	nonce, hash := pow.Run()
	block.Hash = hash[:]
	block.Nonce = nonce
	return block
}

// Genesis create the first block of the blockchain
func Genesis() *Block {
	return CreateBlock("Genesis", []byte{})
}

// Deserialize take data from badgerdb and convert it to a block
func Deserialize(data []byte) *Block {
	var block Block
	decoder := gob.NewDecoder(bytes.NewReader(data))
	err := decoder.Decode(&block)
	Handle(err)
	return &block
}

// Methods

// Serialize convert a block into bytes which is the only format badgerdb accept
func (b *Block) Serialize() []byte {
	var res bytes.Buffer
	encoder := gob.NewEncoder(&res)
	err := encoder.Encode(b)
	Handle(err)
	return res.Bytes()
}
