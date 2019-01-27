package blockchain

// Block represent a single block on the blockchain
// it contain the previous block hash, the data inside the block
// and the current Hash wich is a combination of current data and previous hash
type Block struct {
	Hash     []byte
	Data     []byte
	PrevHash []byte
	Nonce    int
}

// Blockchain is a slice of Block pointers
type Blockchain struct {
	Blocks []*Block
}

// Functions

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

// InitBlockChain create a Blockchain and set the genesis block
func InitBlockChain() *Blockchain {
	return &Blockchain{[]*Block{Genesis()}}
}

// Methods

// DeriveHash take to slices of bytes and use them to create a new hash
// func (b *Block) DeriveHash() {
// 	info := bytes.Join([][]byte{b.Data, b.PrevHash}, []byte{})
// 	hash := sha256.Sum256(info)
// 	b.Hash = hash[:]
// }

// AddBlock append a new block to the blockchain
func (chain *Blockchain) AddBlock(data string) {
	prevBlock := chain.Blocks[len(chain.Blocks)-1]
	newBlock := CreateBlock(data, prevBlock.Hash)
	chain.Blocks = append(chain.Blocks, newBlock)
}
