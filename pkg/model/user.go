package model

import (
	"context"
	"time"
)

type UserInfo struct {
	ID        string
	Name      string
	TonAddr   string
	State     string
	Active    bool
	CreatedAt time.Time
}

type Account struct {
	UserID  string
	Addr    string
	Balance int64
	Status  string
}

type AccountFlow struct {
	Addr    string
	TxnDate time.Time
	Amount  int64
}

func (db *Handler) SaveUserInfo(info *UserInfo) error {
	return db.db.Create(info).Error
}

func (db *Handler) QueryByUserId(ctx context.Context, userId int) (*UserInfo, error) {
	var info UserInfo
	err := db.db.WithContext(ctx).Where("id=?", userId).First(&info).Error
	if err != nil {
		return nil, err
	}
	return &info, nil
}
