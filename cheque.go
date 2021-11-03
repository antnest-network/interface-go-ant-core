package iface

import "context"

type Cheque struct {
	Chequebook         string
	CumulativeReward   string
	CumulativeReleased string
	CashedOut          string
	CanCashOut         string
}

type ChequeAPI interface {
	Get(ctx context.Context, Chequebook string) (Cheque, error)

	List(ctx context.Context) ([]Cheque, error)

	CashOut(ctx context.Context, chequebook string) error

	CashOutAll(ctx context.Context) error
}
