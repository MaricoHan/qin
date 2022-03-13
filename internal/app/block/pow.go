package block

import (
	"bytes"
	"crypto/sha256"
	"math"
	"math/big"
	"qin/internal/app/base"
)

const (
	MAXNonce  = math.MaxInt64
	TargetBit = 26
)

type POW struct {
	block  *Block
	target *big.Int
}

func NewPOW(b *Block) *POW {
	target := big.NewInt(1)
	target.Lsh(target, uint(256-TargetBit))

	return &POW{b, target}
}

func (p *POW) Run() (int64, []byte) {
	var hash [32]byte
	var hashInt big.Int
	nonce := int64(0)
	for nonce < MAXNonce {
		data := p.prepareData(nonce)
		hash = sha256.Sum256(data)
		hashInt.SetBytes(hash[:])

		if hashInt.Cmp(p.target) == -1 {
			break
		}
		nonce++
	}
	return nonce, hash[:]
}
func (p *POW) prepareData(nonce int64) []byte {

	data := bytes.Join([][]byte{
		base.Int2Hex(p.block.Timestamp),
		p.block.PreBlockHash,
		p.block.Data,
		base.Int2Hex(nonce),
	}, []byte{})

	return data
}
func (p *POW) Validate() bool {
	var hashInt big.Int
	data := p.prepareData(p.block.Nonce)
	hash := sha256.Sum256(data)
	hashInt.SetBytes(hash[:])
	return hashInt.Cmp(p.target) == -1
}
