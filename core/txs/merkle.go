package txs

import (
	"crypto/sha256"
	"encoding/hex"
	"strconv"
)
func BuildMerkleTree(txs []Transaction) string {
	if len(txs) == 0 {
		return ""
	}
	hashes := make([]string, len(txs))
	for i, tx := range txs {
		data := strconv.Itoa(tx.ID) + tx.Timestamp + tx.Sender + tx.Receiver + strconv.FormatFloat(tx.Amount, 'f', -1, 64) + tx.Signature
		hash := sha256.Sum256([]byte(data))
		hashes[i] = hex.EncodeToString(hash[:])
	}
	for len(hashes) > 1 {
		tempHashes := []string{}
		for i := 0; i < len(hashes); i += 2 {
			if i + 1 < len(hashes) {
				combined := hashes[i] + hashes[i + 1]
				hash := sha256.Sum256([]byte(combined))
				tempHashes = append(tempHashes, hex.EncodeToString(hash[:]))
			} else {
				tempHashes = append(tempHashes, hashes[i])
			}
		}
		hashes = tempHashes
	}
	return hashes[0]
}
