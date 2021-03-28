package blockchain

import (
	. "yu/common"
	. "yu/result"
	"yu/txn"
)

type IBlock interface {
	BlockId() BlockId
	TxnsHashes() []Hash
	SetTxnsHashes(hashes []Hash)

	SetHash(hash Hash)
	SetPreHash(hash Hash)
	SetStateRoot(hash Hash)
	SetHeight(BlockNum)

	Header() IHeader
	Extra() interface{}
	SetExtra(extra interface{})

	Encode() ([]byte, error)
	Decode(data []byte) (IBlock, error)
}

type IHeader interface {
	Height() BlockNum
	Hash() Hash
	PrevHash() Hash
	TxnRoot() Hash
	StateRoot() Hash
	Timestamp() int64
	Extra() interface{}
}

type IBlockChain interface {
	NewEmptyBlock() IBlock
	// just generate a block with timestamp
	NewDefaultBlock() IBlock
	// pending a block from other blockchain-node for validating and operating
	InsertBlockFromP2P(ib IBlock) error
	// get a pending block
	GetBlocksFromP2P(height BlockNum) ([]IBlock, error)
	// remove a block when finished validating and operate the block
	FlushBlocksFromP2P(height BlockNum) error

	AppendBlock(b IBlock) error
	GetBlock(blockHash Hash) (IBlock, error)
	Children(prevBlockHash Hash) ([]IBlock, error)
	Finalize(blockHash Hash) error
	LastFinalized() (IBlock, error)
	Leaves() ([]IBlock, error)

	// return the longest children chains
	Longest() ([]IChainStruct, error)
	// return the heaviest children chains
	Heaviest() ([]IChainStruct, error)
}

type IChainStruct interface {
	Append(block IBlock)
	InsertPrev(block IBlock)
	First() IBlock
	Last() IBlock
}

type IBlockBase interface {
	GetTxn(txnHash Hash) (txn.IsignedTxn, error)
	SetTxn(stxn txn.IsignedTxn) error

	GetTxns(blockHash Hash) ([]txn.IsignedTxn, error)
	SetTxns(blockHash Hash, txns []txn.IsignedTxn) error

	GetEvents(blockHash Hash) ([]*Event, error)
	SetEvents(events []*Event) error

	GetErrors(blockHash Hash) ([]*Error, error)
	SetErrors(errs []*Error) error
}
