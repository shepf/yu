package blockchain

import (
	"gorm.io/gorm"
	. "yu/common"
	. "yu/result"
	ysql "yu/storage/sql"
	"yu/txn"
)

type BlockBase struct {
	db ysql.SqlDB
}

func NewBlockBase(db ysql.SqlDB) *BlockBase {
	return &BlockBase{
		db: db,
	}
}

func (bb *BlockBase) GetTxn(txnHash Hash) (txn.IsignedTxn, error) {
	var stxn txn.IsignedTxn
	bb.db.Db().Where(&TxnScheme{TxnHash: txnHash.String()}).First(&stxn)
	return stxn, nil
}

func (bb *BlockBase) SetTxn(stxn txn.IsignedTxn) error {
	txnSm, err := toTxnScheme(stxn)
	if err != nil {
		return err
	}
	bb.db.Db().Create(&txnSm)
	return nil
}

func (bb *BlockBase) GetTxns(blockHash Hash) ([]txn.IsignedTxn, error) {
	var txns []txn.SignedTxn
	bb.db.Db().Where(&TxnScheme{BlockHash: blockHash.String()}).Find(&txns)
	var itxns []txn.IsignedTxn
	for _, signedTxn := range txns {
		txns = append(txns, signedTxn)
	}
	return itxns, nil
}

func (bb *BlockBase) SetTxns(blockHash Hash, txns []txn.IsignedTxn) error {
	txnSms := make([]TxnScheme, 0)
	for _, stxn := range txns {
		txnSm, err := newTxnScheme(blockHash, stxn)
		if err != nil {
			return err
		}
		txnSms = append(txnSms, txnSm)
	}
	bb.db.Db().Create(&txnSms)
	return nil
}

func (bb *BlockBase) GetEvents(blockHash Hash) ([]*Event, error) {
	var events []*Event
	bb.db.Db().Where(&EventScheme{BlockHash: blockHash.String()}).Find(&events)
	return events, nil
}

func (bb *BlockBase) SetEvents(events []*Event) error {
	eventSms := make([]EventScheme, 0)
	for _, event := range events {
		eventSm, err := toEventScheme(event)
		if err != nil {
			return err
		}
		eventSms = append(eventSms, eventSm)
	}
	bb.db.Db().Create(&eventSms)
	return nil
}

func (bb *BlockBase) GetErrors(blockHash Hash) ([]*Error, error) {
	var errs []*Error
	bb.db.Db().Where(&ErrorScheme{BlockHash: blockHash.String()}).Find(&errs)
	return errs, nil
}

func (bb *BlockBase) SetErrors(errs []*Error) error {
	errSms := make([]ErrorScheme, 0)
	for _, err := range errs {
		errSms = append(errSms, toErrorScheme(err))
	}
	bb.db.Db().Create(&errSms)
	return nil
}

type TxnScheme struct {
	TxnHash   string `gorm:"primaryKey"`
	Pubkey    string
	Signature string
	RawTxn    string

	BlockHash string
}

func (TxnScheme) TableName() string {
	return "txns"
}

func newTxnScheme(blockHash Hash, stxn txn.IsignedTxn) (TxnScheme, error) {
	txnSm, err := toTxnScheme(stxn)
	if err != nil {
		return TxnScheme{}, err
	}
	txnSm.BlockHash = blockHash.String()
	return txnSm, nil
}

func toTxnScheme(stxn txn.IsignedTxn) (TxnScheme, error) {
	rawTxnByt, err := stxn.GetRaw().Encode()
	if err != nil {
		return TxnScheme{}, err
	}
	return TxnScheme{
		TxnHash:   stxn.GetTxnHash().String(),
		Pubkey:    stxn.GetPubkey().String(),
		Signature: ToHex(stxn.GetSignature()),
		RawTxn:    ToHex(rawTxnByt),
		BlockHash: "",
	}, nil
}

func (t TxnScheme) toTxn() (txn.IsignedTxn, error) {
	ut := &txn.UnsignedTxn{}
	rawTxn, err := ut.Decode(FromHex(t.RawTxn))
	if err != nil {
		return nil, err
	}
	return &txn.SignedTxn{
		Raw:       rawTxn,
		TxnHash:   HexToHash(t.TxnHash),
		Pubkey:    nil,
		Signature: FromHex(t.Signature),
	}, nil
}

type EventScheme struct {
	gorm.Model
	Caller     string
	BlockStage string
	BlockHash  string
	Height     BlockNum
	TripodName string
	ExecName   string
	Value      string
}

func (EventScheme) TableName() string {
	return "events"
}

func toEventScheme(event *Event) (EventScheme, error) {
	value, err := event.ValueStr()
	if err != nil {
		return EventScheme{}, err
	}
	return EventScheme{
		Caller:     event.Caller.String(),
		BlockStage: event.BlockStage,
		BlockHash:  event.BlockHash.String(),
		Height:     event.Height,
		TripodName: event.TripodName,
		ExecName:   event.ExecName,
		Value:      value,
	}, nil
}

func (e EventScheme) toEvent() (*Event, error) {

	return &Event{
		Caller:     HexToAddress(e.Caller),
		BlockStage: e.BlockStage,
		BlockHash:  HexToHash(e.BlockHash),
		Height:     e.Height,
		TripodName: e.TripodName,
		ExecName:   e.ExecName,
		Value:      nil,
	}, nil
}

type ErrorScheme struct {
	gorm.Model
	Caller     string
	BlockStage string
	BlockHash  string
	Height     BlockNum
	TripodName string
	ExecName   string
	Error      string
}

func (ErrorScheme) TableName() string {
	return "errors"
}

func toErrorScheme(err *Error) ErrorScheme {
	return ErrorScheme{
		Caller:     err.Caller.String(),
		BlockStage: err.BlockStage,
		BlockHash:  err.BlockHash.String(),
		Height:     err.Height,
		TripodName: err.TripodName,
		ExecName:   err.ExecName,
		Error:      err.Err.Error(),
	}
}
