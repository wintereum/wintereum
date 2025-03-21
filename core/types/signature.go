package types

import (
	"crypto/ecdsa"
	"crypto/rand"
	"encoding/hex"
	"math/big"
)

type ECDSASigner struct{}

func (s *ECDSASigner) SignBlock(b *Block, privKey *ecdsa.PrivateKey) (string, error) {
	hashBytes, _ := hex.DecodeString(b.Hash)
	r, sSig, err := ecdsa.Sign(rand.Reader, privKey, hashBytes)
	if err != nil {
		return "", err
	}
	signature := append(r.Bytes(), sSig.Bytes()...)
	return hex.EncodeToString(signature), nil
}
func (s *ECDSASigner) VerifyBlock(b *Block, pubKey *ecdsa.PublicKey) bool {
	hashBytes, _ := hex.DecodeString(b.Hash)
	sigBytes, _ := hex.DecodeString(b.Signature)
	r := big.NewInt(0)
	sSig := big.NewInt(0)
	sigLen := len(sigBytes) / 2
	r.SetBytes(sigBytes[:sigLen])
	sSig.SetBytes(sigBytes[sigLen:])
	return ecdsa.Verify(pubKey, hashBytes, r, sSig)
}
