package model

import (
	"context"
	"time"

	"gorm.io/gorm"
)

type UserInfo struct {
	Name      string
	TonAddr   string
	State     string
	Active    bool
	CreatedAt time.Time
	*gorm.Model
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
