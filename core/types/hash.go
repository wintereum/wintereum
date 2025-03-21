package types
import (
	"crypto/sha256"
	"encoding/hex"
	"strconv"
	"time"
)
type BlockHasher struct{}

func HashTime() string {
	t := time.Now()
	return t.String()
}
func (s BlockHasher) CalculateHash(b Block) string {
	record := strconv.Itoa(b.Index) + b.Timestamp + b.Data + b.PrevHash + strconv.Itoa(b.Nonce)
	hash := sha256.New()
	hash.Write([]byte(record + HashTime()))
	hashed := hash.Sum(nil)
	return hex.EncodeToString(hashed)
}
