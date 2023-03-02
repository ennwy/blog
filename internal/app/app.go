package app

import "context"

type Server interface {
	Start() error
	Stop() error
}

type Queue interface {
	Push(ctx context.Context, key string, value []byte) error
	Pop(ctx context.Context, key string) error
}
