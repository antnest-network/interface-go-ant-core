package iface

import (
	"context"
	"math/big"
)

type ChainAPI interface {
	GetBnbBalanceOf(ctx context.Context, address string) (*big.Int, error)

	GetAntzBalanceOf(ctx context.Context, address string) (*big.Int, error)
}
