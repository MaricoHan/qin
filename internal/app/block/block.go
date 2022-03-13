package block

import (
	"bytes"
	"encoding/gob"
	"time"
)

// Block 区块结构
type Block struct {
	Timestamp    int64  //时间戳
	Data         []byte //数据域
	PreBlockHash []byte
	Hash         []byte
	Nonce        int64 //随机值
}

// NewBlock 创建新区块
func NewBlock(data string, preBlockHash []byte) *Block {

	//先构造一个nonce为0的区块
	block := &Block{
		time.Now().Unix(),
		[]byte(data),
		preBlockHash,
		[]byte{},
		0}

	//初始化挖矿的目标值
	p := NewPOW(block)
	//挖该block
	nonce, hash := p.Run()

	block.Nonce = nonce
	block.Hash = hash

	return block
}

// NewGenesisBlock 创建创世区块
func NewGenesisBlock() *Block {
	return NewBlock("Genesis Block", []byte{})
}
func (b *Block) Serialize() []byte {
	var res bytes.Buffer
	//生成编码器
	encoder := gob.NewEncoder(&res)
	encoder.Encode(b)
	return res.Bytes()
}
