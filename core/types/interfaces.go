package types

import (
	. "github.com/yu-org/yu/common"
)

//type IBlock interface {
//	IHeader
//	GetHeader() IHeader
//
//	GetBlockId() BlockId
//	GetTxnsHashes() []Hash
//	SetTxnsHashes(hashes []Hash)
//
//	SetHash(hash Hash)
//	SetPreHash(hash Hash)
//	SetTxnRoot(hash Hash)
//	SetStateRoot(hash Hash)
//	SetHeight(BlockNum)
//	SetTimestamp(ts uint64)
//	SetPeerID(peer.ID)
//
//	SetLeiLimit(e uint64)
//	UseLei(e uint64)
//
//	SetSignature([]byte)
//
//	Encode() ([]byte, error)
//	Decode(data []byte) (IBlock, error)
//
//	CopyFrom(other IBlock)
//}
//
//type IHeader interface {
//	GetHeight() BlockNum
//	GetHash() Hash
//	GetPrevHash() Hash
//	GetTxnRoot() Hash
//	GetStateRoot() Hash
//	GetTimestamp() uint64
//	GetPeerID() peer.ID
//	GetLeiLimit() uint64
//	GetLeiUsed() uint64
//
//	GetSignature() []byte
//}

// --------------- blockchain interface ----------------

type ConvergeType int

const (
	Longest ConvergeType = iota
	Heaviest
	Finalize
)

type IBlockChain interface {
	ItxDB
	ConvergeType() ConvergeType

	NewEmptyBlock() *Block

	GetGenesis() (*Block, error)
	SetGenesis(b *Block) error

	AppendBlock(b *Block) error
	GetBlock(blockHash Hash) (*CompactBlock, error)
	GetBlockByHeight(height BlockNum) (*CompactBlock, error)
	GetAllBlocksByHeight(height BlockNum) ([]*CompactBlock, error)
	ExistsBlock(blockHash Hash) (bool, error)
	UpdateBlock(b *CompactBlock) error
	UpdateBlockByHeight(b *CompactBlock) error

	Children(prevBlockHash Hash) ([]*CompactBlock, error)
	Finalize(blockHash Hash) error
	LastFinalized() (*CompactBlock, error)
	GetEndBlock() (*CompactBlock, error)
	GetAllBlocks() ([]*CompactBlock, error)

	GetRangeBlocks(startHeight, endHeight BlockNum) ([]*Block, error)
}

type ItxDB interface {
	GetTxn(txnHash Hash) (*SignedTxn, error)
	ExistTxn(txnHash Hash) bool
	SetTxns(txns []*SignedTxn) error

	SetReceipts(receipts map[Hash]*Receipt) error
	GetReceipt(txHash Hash) (*Receipt, error)
	SetReceipt(txHash Hash, receipt *Receipt) error
}
