package model

import (
	"context"
	"time"
)

type RRTxnDetail struct {
	UserId    int
	TxnId     string
	Amount    int64
	SourceId  int
	Type      string // 1 - deposit, 2 - withdraw
	CreatedAt time.Time
}

type RRSourceType struct {
	Id   int
	Name string
}

func (db *Handler) SaveRRTxnDetail(txnDetail *RRTxnDetail) error {
	return db.db.Create(txnDetail).Error
}

func (db *Handler) GetRRBalance(ctx context.Context, userId int) (int64, error) {
	return 0, nil
}

type WithhdrawProposal struct {
	Id     int
	Type   string
	Amount int64
}

func (db *Handler) SaveWithdrawProposal(proposal *WithhdrawProposal) error {
	return db.db.Create(proposal).Error
}

type TonTxnDetail struct {
	TxnId    int32
	UserId   int
	Amount   int32
	CreateAt time.Time
}

func (db *Handler) SaveTonAccount(account *TonTxnDetail) error {
	return db.db.Create(account).Error
}

type TonAccount struct {
	UserId         int
	DepositAmount  int32
	WithdrawAmount int32
}

func (db *Handler) GetTonAccountByUserId(ctx context.Context, userId int) (*TonAccount, error) {
	acc := new(TonAccount)
	return acc, nil
}
