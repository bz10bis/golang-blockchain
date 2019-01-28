package blockchain

// TxOutput contain the value of the transaction and the key to access it
type TxOutput struct {
	Value  int
	PubKey string
}

// TxInput contain the ID of the tx the index of the output and the signature
type TxInput struct {
	ID  []byte
	Out int
	Sig string
}

// Methods

// CanUnlock check if the account in "data" can access the tx
func (in *TxInput) CanUnlock(data string) bool {
	return in.Sig == data
}

// CanBeUnlocked check if the account in "data" can access the tx
func (out *TxOutput) CanBeUnlocked(data string) bool {
	return out.PubKey == data
}
