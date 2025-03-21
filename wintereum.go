package github.com/wintereum/wintereum

import (
	"github.com/wintereum/wintereum/core"
	"github.com/wintereum/wintereum/core/txs"
	"github.com/wintereum/wintereum/core/types"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"fmt"
	"time"
)

func main() {
	privKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		panic(err)
	}
	pubKey := &privKey.PublicKey
	hasher := types.BlockHasher{}
	pow := types.Proof_of_Work{
		Hasher: hasher,
		Difficulty: 1,
	}
	signer := types.ECDSASigner{}
	txHasher := txs.TxsHasher{}
	bc := &core.Blockchain{
		Blocks: []*types.Block{core.GenesisBlock(&pow, &signer, privKey)},
		Hasher: hasher,
		PoW: &pow,
		Signer: &signer,
		TxHasher: txHasher,
		PrivKey: privKey,
		PubKey: pubKey,
	}

	fmt.Printf("Starting blockchain...\n")
	bc.PrintChain()
	go func() {
		counter := 1
		for {
			data := fmt.Sprintf("Block %d - Auto-generated", counter)
			bc.AddBlock(data)
			fmt.Printf("Added block %d\n", counter)
			bc.PrintNewBlock()
			time.Sleep(60 * time.Second) // Tạm dừng 2 giây giữa các block
			counter++
		}
	}()
	select {}
}
