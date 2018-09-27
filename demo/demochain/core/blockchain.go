package core

import (
	"fmt"
	"log"
)

type BlockChain struct {
	Blocks []*Block
}

func NewBlockChain() *BlockChain {
	genesisBlock := GenerateGenesisBlock()
	blockChain := BlockChain{}
	blockChain.AppendBlock(&genesisBlock)
	return &blockChain
}

func (bc *BlockChain) SendData(data string) {
	preBlock := bc.Blocks[len(bc.Blocks)-1]
	newBlock := GenerateNewBlock(*preBlock, data)
	bc.AppendBlock(&newBlock)
}

func (bc *BlockChain) AppendBlock(newBlock *Block) {
	if len(bc.Blocks) == 0 {
		bc.Blocks = append(bc.Blocks, newBlock)
		return
	}
	if isValid(*newBlock, *bc.Blocks[len(bc.Blocks)-1]) {
		bc.Blocks = append(bc.Blocks, newBlock)
	} else {
		log.Fatal("invalid block")
	}
}

func isValid(newBlock Block, oldBlock Block) bool {
	if newBlock.Index != oldBlock.Index+1 {
		return false
	}
	if newBlock.PreBlockHash != oldBlock.Hash {
		return false
	}
	if calculateHash(newBlock) != newBlock.Hash {
		return false
	}
	return true
}

func (bc *BlockChain) Print() {
	for _, block := range bc.Blocks {
		fmt.Printf("Index: %d\nPrev.Hash: %s\nCurr.Hash: %s\nData: %s\nTimeStamp: %d\n\n", block.Index, block.PreBlockHash, block.Hash, block.Data, block.Timestamp)
	}
}
