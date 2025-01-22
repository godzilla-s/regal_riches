package tonclient

import (
	"context"
	"math/big"

	"github.com/xssnick/tonutils-go/address"
	"github.com/xssnick/tonutils-go/liteclient"
	"github.com/xssnick/tonutils-go/tlb"
	"github.com/xssnick/tonutils-go/ton"
	"github.com/xssnick/tonutils-go/ton/wallet"
)

type Client struct {
	client *liteclient.ConnectionPool
	api    *ton.APIClient
}

func New(url string) (*Client, error) {
	client := liteclient.NewConnectionPool()
	cfg, err := liteclient.GetConfigFromUrl(context.Background(), "")
	if err != nil {
		return nil, err
	}

	err = client.AddConnectionsFromConfig(context.Background(), cfg)
	if err != nil {
		return nil, err
	}

	return &Client{client: client, api: ton.NewAPIClient(client)}, nil
}

func (c *Client) GetBalance(ctx context.Context, addr string) (*tlb.Account, error) {
	b, err := c.api.CurrentMasterchainInfo(c.client.StickyContext(ctx))
	if err != nil {
		return nil, err
	}

	res, err := c.api.WaitForBlock(b.SeqNo).GetAccount(ctx, b, address.MustParseAddr(addr))
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (c *Client) Payment(ctx context.Context, fromAddr string, amount int64) error {
	return c.transfer(ctx, fromAddr, "", amount)
}

func (c *Client) Accept(ctx context.Context, toAddr string, amount int64) error {
	return c.transfer(ctx, "", toAddr, amount)
}

func (c *Client) transfer(ctx context.Context, fromAddr, toAddr string, amount int64) error {
	from := address.MustParseAddr(fromAddr)
	to := address.MustParseAddr(toAddr)

	_, _ = from, to
	w, err := wallet.FromSeed(c.api, []string{}, wallet.V4R2)
	if err != nil {
		return err
	}

	message, err := w.BuildTransfer(to, tlb.FromNanoTON(big.NewInt(amount)), false, "")
	if err != nil {
		return err
	}

	w.SendWaitTransaction(ctx, message)

	return nil
}
