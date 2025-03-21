package txs

import (
	"crypto/sha256"
	"encoding/hex"
	"strconv"
)
type TxsHasher struct{}
func (t *TxsHasher) HashTransaction(tx Transaction) string {
	record := strconv.Itoa(tx.ID) + tx.Timestamp + tx.Sender + tx.Receiver + strconv.FormatFloat(tx.Amount, 'f', -1, 64)
	hash := sha256.New()
	hash.Write([]byte(record))
	hashed := hash.Sum(nil)
	return hex.EncodeToString(hashed)
}
func (t *TxsHasher) HashSignature(tx Transaction) string {
	record := strconv.Itoa(tx.ID) + tx.Timestamp + tx.Sender + tx.Receiver + strconv.FormatFloat(tx.Amount, 'f', -1, 64) + tx.Signature
	hash := sha256.New()
	hash.Write([]byte(record))
	hashed := hash.Sum(nil)
	return hex.EncodeToString(hashed)
}
