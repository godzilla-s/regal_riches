package service

import (
	"errors"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/godzilla-s/regal-riches/pkg/model"
)

var (
	ErrUserNotFound = errors.New("user not found")
)

type LoginRequest struct {
	UserID int
}

type LoginReply struct {
}

func (s *Service) Login(ctx *gin.Context) {
	var req LoginRequest
	err := ctx.Bind(&req)
	if err != nil {
		return
	}

	_, err = s.db.QueryByUserId(ctx, req.UserID)
	if err != nil {
		ctx.JSON(500, ErrUserNotFound)
		return
	}
	ctx.JSON(200, &LoginReply{})
}

type RegistryRequest struct {
	Name    string `json:"name"`
	State   string `json:"state"`
	TonAddr string `json:"ton_addr"`
}

type RegistryReply struct {
	UserId int `json:"user_id"`
}

func (s *Service) Registry(ctx *gin.Context) {
	req := &RegistryRequest{}
	err := ctx.Bind(req)
	if err != nil {
		ctx.JSON(500, err)
		return
	}

	userInfo := &model.UserInfo{
		Name:      req.Name,
		State:     "",
		TonAddr:   req.TonAddr,
		Active:    true,
		CreatedAt: time.Now(),
	}
	err = s.db.SaveUserInfo(userInfo)
	if err != nil {
		ctx.JSON(500, err)
		return
	}

	ctx.JSON(200, &RegistryReply{UserId: int(userInfo.ID)})
}
