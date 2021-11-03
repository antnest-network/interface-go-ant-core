package iface

import (
	"context"
	"crypto/ecdsa"
)

type WalletAPI interface {
	NewAddress(ctx context.Context) (*ecdsa.PrivateKey, error)

	Import(ctx context.Context, privateKeyHex string) (*ecdsa.PrivateKey, error)

	Get(ctx context.Context, address string) (*ecdsa.PrivateKey, error)

	SetDefaultAddress(ctx context.Context, address string) error

	GetDefaultAddress(context.Context) (*ecdsa.PrivateKey, error)

	List(ctx context.Context) ([]*ecdsa.PrivateKey, error)

	Delete(ctx context.Context, address string) error
}
