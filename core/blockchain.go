package core

import (
	"crypto/ecdsa"
	"fmt"
	"time"

	"github.com/wintereum/wintereum/core/txs"
	"github.com/wintereum/wintereum/core/types"
)

type Blockchain struct {
	Blocks 		[]*types.Block
	Hasher 		types.Hasher
	PoW 		types.PoW
	Signer 		types.Signer
	TxHasher 	txs.TxsHasher
	PrivKey 	*ecdsa.PrivateKey
	PubKey 		*ecdsa.PublicKey
}

func GenesisBlock(pow types.PoW, signer types.Signer, privKey *ecdsa.PrivateKey) *types.Block {
	GenesisBlock := &types.Block{
		Index: 0,
		Timestamp: time.Now().String(),
		Data: "Genesis Block",
		PrevHash: "",
		Nonce: 0,
	}
	GenesisBlock.Hash = pow.ProofOfWork(GenesisBlock)
	sig, err := signer.SignBlock(GenesisBlock, privKey)
	if err != nil {
		panic(err)
	}
	GenesisBlock.Signature = sig
	return GenesisBlock
}
func(bc *Blockchain) AddBlock(data string) {
	prevBlock := bc.Blocks[len(bc.Blocks) - 1]
	newBlock := NewBlock(data, []txs.Transaction{}, prevBlock, bc.PoW, bc.Signer, bc.TxHasher, bc.PrivKey)
	if !bc.Signer.VerifyBlock(newBlock, bc.PubKey) {
		panic("Block verification failed")
	}
	bc.Blocks = append(bc.Blocks, newBlock)
}
func NewBlock(data string, tx []txs.Transaction, prevBlock *types.Block, pow types.PoW, signer types.Signer, txHasher txs.TxsHasher, privKey *ecdsa.PrivateKey) *types.Block {
	block := &types.Block {
		Index: prevBlock.Index + 1,
		Timestamp: time.Now().String(),
		Data: data,
		Transactions: tx,
		PrevHash: prevBlock.Hash,
		Nonce: 0,
	}
	block.MerkleRoot = txs.BuildMerkleTree(tx)
	block.Hash = pow.ProofOfWork(block)
	sig, err := signer.SignBlock(block, privKey)
	if err != nil {
		panic(err)
	}
	block.Signature = sig
	return block
}
func (bc *Blockchain) PrintNewBlock() {
	fmt.Printf("Index:      %d\n", bc.Blocks[len(bc.Blocks) - 1].Index)
	fmt.Printf("Timestamp:  %s\n", bc.Blocks[len(bc.Blocks) - 1].Timestamp)
	fmt.Printf("Data:       %s\n", bc.Blocks[len(bc.Blocks) - 1].Data)
	if len(bc.Blocks[len(bc.Blocks) - 1].Transactions) > 0 {
		fmt.Printf("Transaction:")
		for i, tx := range bc.Blocks[len(bc.Blocks) - 1].Transactions {
			fmt.Printf("   %d. %s -> %s: %.2f (Sig: %s)\n",
				i, tx.Sender, tx.Receiver, tx.Amount, tx.Signature[:8]+"...\n")
		}
	} else {
		fmt.Printf("Transaction: None\n")
	}
	fmt.Printf("PrevHash:   %s\n", bc.Blocks[len(bc.Blocks) - 1].PrevHash)
	fmt.Printf("Hash:       %s\n", bc.Blocks[len(bc.Blocks) - 1].Hash)
	fmt.Printf("Nonce: 	    %d\n", bc.Blocks[len(bc.Blocks) - 1].Nonce)
	fmt.Printf("Signature:  %s\n", bc.Blocks[len(bc.Blocks) - 1].Signature)
	fmt.Printf("MerkleRoot: %s\n", bc.Blocks[len(bc.Blocks) - 1].MerkleRoot)
	fmt.Printf("Verified: %v\n", bc.Signer.VerifyBlock(bc.Blocks[len(bc.Blocks) - 1], bc.PubKey))
	fmt.Println("----------------------------------------------------------------------------------")
}
func (bc *Blockchain) PrintChain() {
	for _, block := range bc.Blocks {
		fmt.Printf("Index: %d\n", block.Index)
		fmt.Printf("Timestamp: %s\n", block.Timestamp)
		fmt.Printf("Data: %s\n", block.Data)
		if len(block.Transactions) > 0 {
			fmt.Printf("Transaction:")
			for i, tx := range block.Transactions {
				fmt.Printf("  %d. %s -> %s: %.2f (Sig: %s)\n",
					i, tx.Sender, tx.Receiver, tx.Amount, tx.Signature[:8]+"...\n")
			}
		} else {
			fmt.Printf("Transaction: None\n")
		}
		fmt.Printf("PrevHash: %s\n", block.PrevHash)
		fmt.Printf("Hash: %s\n", block.Hash)
		fmt.Printf("Nonce: %d\n", block.Nonce)
		fmt.Printf("Signature: %s\n", block.Signature)
		fmt.Printf("MerkleRoot: %s\n", block.MerkleRoot)
		fmt.Printf("Verified: %v\n\n", bc.Signer.VerifyBlock(block, bc.PubKey))
	}
	fmt.Println("-------------------------------------------------------------------------------------")
}
