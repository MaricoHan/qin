package chain

import (
	"github.com/boltdb/bolt"
	"qin/internal/app/block"
	"qin/internal/pkg/log"
)

const (
	dbFile       = "blockchain.db"
	blocksBucket = "blocks"
)

// Blockchain 区块链：指针的切片
type Blockchain struct {
	latestHash []byte
	db         *bolt.DB
}

// NewBlockchain 初始化链，只有创世区块
func NewBlockchain() *Blockchain {
	var latest []byte
	//打开数据库文件
	db, _ := bolt.Open(dbFile, 0600, nil)
	//更新数据库
	db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(blocksBucket))
		if bucket == nil {
			log.Debug.Println("init a new blockchain ...")
			bucket, _ = tx.CreateBucket([]byte(blocksBucket))
			genesis := block.NewGenesisBlock()

			bucket.Put(genesis.Hash, genesis.Serialize())
			bucket.Put([]byte("latest"), genesis.Hash)
			latest = genesis.Hash
		} else {
			latest = bucket.Get([]byte("latest"))
		}

		return nil
	})

	return &Blockchain{latest, db}
}

// AddBlock 后续添加区块
func (b *Blockchain) AddBlock(data string) {
	var latest []byte
	//查出最新hash 和block
	b.db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(blocksBucket))
		latest = bucket.Get([]byte("latest"))
		return nil
	})

	//生成block
	newBlock := block.NewBlock(data, latest)

	//插入
	b.db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(blocksBucket))
		bucket.Put(newBlock.Hash, newBlock.Serialize())

		//更新最新hash
		bucket.Put([]byte("latest"), newBlock.Hash)
		b.latestHash = newBlock.Hash

		return nil
	})

}
