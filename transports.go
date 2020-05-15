package model_go

// FirstBlockTransport is a container of the first two blocks of the chain
// with their following Firsts Transactions
type GenesisBlocksTransport struct {
	Block1        *Block
	Block2        *Block
	Transaction11 *Transaction
	Transaction21 *Transaction
}
