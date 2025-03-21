package txs

import (
	"crypto/ecdsa"
)

type Transaction struct {
	ID		  int
	Timestamp string
	Sender	  string
	Receiver  string
	Amount	  float64
	Hash	  string
	Signature string
}
type TxHasher interface {
	HashTransaction(tx Transaction) string
	HashSignature(tx Transaction) string
}
type Signer interface {
	SignTransaction(t *Transaction, privKey *ecdsa.PrivateKey) (string, error)
	VerifyTransaction(t *Transaction, pubKey *ecdsa.PublicKey) bool
}
type MerkleTree interface {
	BuildMerkleTree(txs []Transaction) string
}
