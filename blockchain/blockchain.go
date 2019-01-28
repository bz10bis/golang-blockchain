package blockchain

import (
	"fmt"

	"github.com/dgraph-io/badger"
)

const (
	dbPath = "./db/blocks"
)

// Blockchain is a slice of Block pointers
type Blockchain struct {
	LastHash []byte
	Database *badger.DB
}

// BlockchainIterator is used to iterate through the blockchain backward
type BlockchainIterator struct {
	CurrentHash []byte
	Database    *badger.DB
}

// InitBlockChain create a Blockchain and set the genesis block
// If no blockchain is present on badgerDb => create One
// else load the blockchain to memory from badgerdb
func InitBlockChain() *Blockchain {
	var lastHash []byte

	opts := badger.DefaultOptions
	opts.Dir = dbPath
	opts.ValueDir = dbPath
	db, err := badger.Open(opts)
	Handle(err)

	// Db.update is a read and write operation
	err = db.Update(func(txn *badger.Txn) error {
		// "lh" key correspond to lastHash
		if _, err := txn.Get([]byte("lh")); err == badger.ErrKeyNotFound {
			fmt.Printf("No existing blocks found")
			genesis := Genesis()
			fmt.Printf("Genesis Verified")
			err = txn.Set(genesis.Hash, genesis.Serialize())
			Handle(err)
			err = txn.Set([]byte("lh"), genesis.Hash)
			lastHash = genesis.Hash
			return err
		}
		item, err := txn.Get([]byte("lh"))
		Handle(err)
		lastHash, err = item.Value()
		return err
	})
	Handle(err)
	blockchain := Blockchain{lastHash, db}
	return &blockchain
}

// Methods

// AddBlock append a new block to the blockchain
func (chain *Blockchain) AddBlock(data string) {
	var lastHash []byte

	// db.View() is a read only operation
	err := chain.Database.View(func(txn *badger.Txn) error {
		item, err := txn.Get([]byte("lh"))
		Handle(err)
		lastHash, err = item.Value()
		return err
	})
	Handle(err)
	newBlock := CreateBlock(data, lastHash)
	// After creating and validating the new block we serialize it in badgerdb
	err = chain.Database.Update(func(txn *badger.Txn) error {
		err := txn.Set(newBlock.Hash, newBlock.Serialize())
		Handle(err)
		// We save the new block hash into "lh" key
		err = txn.Set([]byte("lh"), newBlock.Hash)
		chain.LastHash = newBlock.Hash
		return err
	})
	Handle(err)
}

// Iterator Create an BlockchainIterator on our Blockchain
func (chain *Blockchain) Iterator() *BlockchainIterator {
	iter := &BlockchainIterator{chain.LastHash, chain.Database}
	return iter
}

// Next iterate through the next block in the blockchain
func (iter *BlockchainIterator) Next() *Block {
	var block *Block
	err := iter.Database.View(func(txn *badger.Txn) error {
		item, err := txn.Get(iter.CurrentHash)
		Handle(err)
		serializedBlock, err := item.Value()
		block = Deserialize(serializedBlock)
		return err
	})
	Handle(err)
	iter.CurrentHash = block.PrevHash
	return block
}
