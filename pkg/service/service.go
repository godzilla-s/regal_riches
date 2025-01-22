package service

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/godzilla-s/regal-riches/pkg/model"
	"github.com/godzilla-s/regal-riches/pkg/tonclient"
)

type Service struct {
	db  *model.Handler
	ton *tonclient.Client
}

func NewService(c *Config) *Service {
	db, err := model.New(c.DBConfig)
	if err != nil {
		panic(err)
	}

	ton, err := tonclient.New(c.Url)
	if err != nil {
		panic(err)
	}
	return &Service{db: db, ton: ton}
}

type RecieveRequest struct {
	UserId int
	Amount int
}

type RecieveReply struct {
}

func (s *Service) ReciveRR(ctx *gin.Context) {
	req := RecieveRequest{}
	err := ctx.Bind(req)
	if err != nil {
		ctx.JSON(500, err)
		return
	}

	userInfo, err := s.db.QueryByUserId(ctx, req.UserId)
	if err != nil {
		ctx.JSON(500, ErrUserNotFound)
		return
	}

	if !userInfo.Active {
		return
	}

	err = s.db.SaveRRTxnDetail(&model.RRTxnDetail{
		UserId:    req.UserId,
		Amount:    int64(req.Amount),
		CreatedAt: time.Now(),
	})
	if err != nil {
		ctx.JSON(500, err)
		return
	}
}

type PayRequest struct {
	UserId int
	Amount int64
}

type PayReply struct {
	TxnId string
}

func (s *Service) PayRR(ctx *gin.Context) {
	req := RecieveRequest{}
	err := ctx.Bind(req)
	if err != nil {
		ctx.JSON(500, err)
		return
	}

	userInfo, err := s.db.QueryByUserId(ctx, req.UserId)
	if err != nil {
		ctx.JSON(500, ErrUserNotFound)
		return
	}
	if !userInfo.Active {
		return
	}
	err = s.db.SaveRRTxnDetail(&model.RRTxnDetail{})
	if err != nil {
		ctx.JSON(500, err)
		return
	}
	ctx.JSON(200, &PayReply{})
}

type GetRRBalanceRequest struct {
	UserId int
}

type GetRRBalanceReply struct {
	Amount int64
}

func (s *Service) GetRRBalance(ctx *gin.Context) {
	req := &GetRRBalanceRequest{}
	err := ctx.Bind(req)
	if err != nil {
		ctx.JSON(500, err)
		return
	}

	userInfo, err := s.db.QueryByUserId(ctx, req.UserId)
	if err != nil {
		ctx.JSON(500, ErrUserNotFound)
		return
	}

	if !userInfo.Active {
		return
	}
	bal, err := s.db.GetRRBalance(ctx, req.UserId)
	if err != nil {
		ctx.JSON(500, err)
		return
	}

	ctx.JSON(200, &GetRRBalanceReply{Amount: bal})
}

type WithdrawProposalRequest struct {
	UserId int
	Amount int32
}

type WithdrawProposalReply struct {
	Success bool
	Balance int32
}

func (s *Service) WithdrawProposal(ctx *gin.Context) {
	req := &WithdrawProposalRequest{}
	err := ctx.Bind(req)
	if err != nil {
		ctx.JSON(500, err)
		return
	}
	userInfo, err := s.db.QueryByUserId(ctx, req.UserId)
	if err != nil {
		ctx.JSON(500, err)
		return
	}

	if !userInfo.Active {
		return
	}

	s.db.SaveWithdrawProposal(&model.WithhdrawProposal{})
	ctx.JSON(200, &WithdrawProposalReply{})
}

func (s *Service) WithdrawConfirm() {

}

type DepositTonRequest struct {
	UserId int
	Amount int32
}

type DepositTonReply struct {
	TxnId int
}

func (s *Service) DepositTON(ctx *gin.Context) {
	req := new(DepositTonRequest)
	err := ctx.Bind(req)
	if err != nil {
		ctx.JSON(500, err)
		return
	}

	userInfo, err := s.db.QueryByUserId(ctx, req.UserId)
	if err != nil {
		ctx.JSON(500, err)
		return
	}

	if !userInfo.Active {
		return
	}

	txnDetail := &model.TonTxnDetail{
		UserId:   req.UserId,
		Amount:   req.Amount,
		Type:     model.TxnTypeDeposit,
		CreateAt: time.Now(),
	}
	err = s.db.SaveTonAccount(txnDetail)
	if err != nil {
		ctx.JSON(500, err)
		return
	}
	ctx.JSON(200, &DepositTonReply{TxnId: int(txnDetail.TxnId)})
}

type TonBalanceRequest struct {
	UserId int
}

type TonBalanceReply struct {
	Balance int32
}

func (s *Service) GetTonBalance(ctx *gin.Context) {
	req := new(TonBalanceRequest)
	err := ctx.Bind(req)
	if err != nil {
		ctx.JSON(500, err)
		return
	}

	acc, err := s.db.GetTonAccountByUserId(ctx, req.UserId)
	if err != nil {
		ctx.JSON(500, err)
		return
	}

	ctx.JSON(200, &TonBalanceReply{Balance: acc.DepositAmount - acc.WithdrawAmount})
}
