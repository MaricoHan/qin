package main

// Blockchain 区块链：指针的切片
type Blockchain struct {
	blocks []*Block
}

// NewBlockchain 初始化链，只有创世区块
func NewBlockchain() *Blockchain{
	return &Blockchain{[]*Block{NewGenesisBlock()}}
}

// AddBlock 后续添加区块
func (bc *Blockchain) AddBlock(data string){
	prevBlock := bc.blocks[len(bc.blocks)-1]
	newBlock := NewBlock(data, prevBlock.Hash)
	bc.blocks = append(bc.blocks, newBlock)
}
