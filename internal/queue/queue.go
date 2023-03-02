package queue

import (
	"context"
	"net"
	"os"

	"github.com/ennwy/blog/internal/app"
	"github.com/redis/go-redis/v9"
)

type Q struct {
	storage *redis.Client
}

var _ app.Queue = (*Q)(nil)

func New() *Q {
	opt := &redis.Options{
		Addr:     net.JoinHostPort(os.Getenv("Q_HOST"), os.Getenv("Q_PORT")),
		Password: os.Getenv("Q_PASSWORD"),
		DB:       0,
	}

	queue := &Q{
		storage: redis.NewClient(opt),
	}

	return queue
}

func (q *Q) Push(ctx context.Context, key string, value []byte) error {
	return q.storage.RPush(ctx, key, value).Err()
}

func (q *Q) Pop(ctx context.Context, key string) error {
	return q.storage.LPop(ctx, key).Err()
}
