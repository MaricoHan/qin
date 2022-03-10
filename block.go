package main

import (
	"bytes"
	"crypto/sha256"
	"strconv"
)

// Block 区块结构
type Block struct {
	Timestamp    int64  //时间戳
	Data         []byte //数据域
	PreBlockHash []byte
	Hash         []byte
	Nonce        int64 //随机值
}

// SetHash 区块设置内部hash
func (b *Block) SetHash() {
	timestamp := []byte(strconv.FormatInt(b.Timestamp, 10))
	headers := bytes.Join([][]byte{b.PreBlockHash, b.Data, timestamp}, []byte{})
	//计算本块的hash
	hash := sha256.Sum256(headers)
	//[32]byte->[]byte
	b.Hash = hash[:]
}

//// NewBlock 创建新区块
//func NewBlock(data string, preBlockHash []byte) *Block {
//	//构造区块
//	block := &Block{time.Now().Unix(), []byte(data), preBlockHash, []byte{}}
//	block.SetHash()
//	return block
//}

//// NewGenesisBlock 创建创世区块
//func NewGenesisBlock() *Block {
//	return NewBlock("Genesis Block", []byte{})
//}
