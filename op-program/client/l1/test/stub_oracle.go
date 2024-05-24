package test

import (
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/tokamak-network/tokamak-thanos/op-service/eth"
)

type StubOracle struct {
	t *testing.T

	// Blocks maps block hash to eth.BlockInfo
	Blocks map[common.Hash]eth.BlockInfo

	// Txs maps block hash to transactions
	Txs map[common.Hash]types.Transactions

	// Rcpts maps Block hash to receipts
	Rcpts map[common.Hash]types.Receipts

	// Blobs maps indexed blob hash to l1 block ref to blob
	Blobs map[eth.L1BlockRef]map[eth.IndexedBlobHash]*eth.Blob

	// PtEvals maps hashed input to whether the KZG point evaluation was successful
	PtEvals map[common.Hash]bool
}

func NewStubOracle(t *testing.T) *StubOracle {
	return &StubOracle{
		t:       t,
		Blocks:  make(map[common.Hash]eth.BlockInfo),
		Txs:     make(map[common.Hash]types.Transactions),
		Rcpts:   make(map[common.Hash]types.Receipts),
		Blobs:   make(map[eth.L1BlockRef]map[eth.IndexedBlobHash]*eth.Blob),
		PtEvals: make(map[common.Hash]bool),
	}
}

func (o StubOracle) HeaderByBlockHash(blockHash common.Hash) eth.BlockInfo {
	info, ok := o.Blocks[blockHash]
	if !ok {
		o.t.Fatalf("unknown block %s", blockHash)
	}
	return info
}

func (o StubOracle) TransactionsByBlockHash(blockHash common.Hash) (eth.BlockInfo, types.Transactions) {
	txs, ok := o.Txs[blockHash]
	if !ok {
		o.t.Fatalf("unknown txs %s", blockHash)
	}
	return o.HeaderByBlockHash(blockHash), txs
}

func (o StubOracle) ReceiptsByBlockHash(blockHash common.Hash) (eth.BlockInfo, types.Receipts) {
	rcpts, ok := o.Rcpts[blockHash]
	if !ok {
		o.t.Fatalf("unknown rcpts %s", blockHash)
	}
	return o.HeaderByBlockHash(blockHash), rcpts
}

func (o StubOracle) GetBlob(ref eth.L1BlockRef, blobHash eth.IndexedBlobHash) *eth.Blob {
	blobMap, ok := o.Blobs[ref]
	if !ok {
		o.t.Fatalf("unknown blob ref %s", ref)
	}
	blob, ok := blobMap[blobHash]
	if !ok {
		o.t.Fatalf("unknown blob hash %s %d", blobHash.Hash, blobHash.Index)
	}
	return blob
}

func (o StubOracle) KZGPointEvaluation(input []byte) bool {
	result, ok := o.PtEvals[crypto.Keccak256Hash(input)]
	if !ok {
		o.t.Fatalf("unknown kzg point evaluation %x", input)
	}
	return result
}
