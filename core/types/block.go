package types

import (
	"crypto/ecdsa"
	"github.com/wintereum/wintereum/core/txs"
)

type Block struct {
	Index      	 int
	Timestamp  	 string
	Data	   	 string
	Transactions []txs.Transaction
	PrevHash   	 string
	Hash       	 string
	Nonce	   	 int
	Signature  	 string
	MerkleRoot 	 string
}

type Hasher interface {
	CalculateHash(b Block) string
}

type PoW interface {
	ProofOfWork(b *Block) string
}

type Signer interface {
	SignBlock(b *Block, privKey *ecdsa.PrivateKey) (string, error)
	VerifyBlock(b *Block, pubKey *ecdsa.PublicKey) bool
}
type MerkleTree interface {
	BuildMerkleTree(txs []txs.Transaction) string
}
