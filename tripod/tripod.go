package tripod

import (
	. "yu/blockchain"
	. "yu/env"
	. "yu/txn"
)

type Tripod interface {
	TripodMeta() *TripodMeta

	CheckTxn(*SignedTxn) error

	ValidateBlock(IBlockChain, IBlock) bool

	InitChain(env *Env, land *Land) error

	StartBlock(env *Env, land *Land) (needBroadcast bool, err error)

	EndBlock(env *Env, land *Land) error

	FinalizeBlock(env *Env, land *Land) error
}
