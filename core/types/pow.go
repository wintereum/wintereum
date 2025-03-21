package types

import (
	"strings"
	"time"
)

type Proof_of_Work struct{
	Hasher Hasher
	Difficulty int
	LastBlockTime time.Time
}

func (p *Proof_of_Work) ProofOfWork(b *Block) string {
	target := strings.Repeat("0", p.Difficulty)
	for {
		hash := p.Hasher.CalculateHash(*b)
		if hash[:p.Difficulty] == target {
			b.Hash = hash
			now := time.Now()
			if !p.LastBlockTime.IsZero() {
				elapsed := now.Sub(p.LastBlockTime).Seconds()
				if elapsed < 15 {
					p.Difficulty++
				} else if elapsed > 20 {
					if p.Difficulty > 1 {
						p.Difficulty--
					}
				}
			}
			p.LastBlockTime = now
			return hash
		}
		b.Nonce++
	}
}
