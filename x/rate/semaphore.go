package rate

import (
	"context"

	"github.com/hibiken/asynq"
	"github.com/redis/go-redis/v9"
)

func NewSemaphore(rco asynq.RedisConnOpt, scope string, maxTokens int) *Semaphore {
	_ = "STUB: not implemented"
	return nil
}

type Semaphore struct {
	rc        redis.UniversalClient
	maxTokens int
	scope     string
}

var acquireCmd = redis.NewScript(`
redis.call("ZREMRANGEBYSCORE", KEYS[1], "-inf", tonumber(ARGV[2])-1)
local count = redis.call("ZCARD", KEYS[1])

if (count < tonumber(ARGV[1])) then
     redis.call("ZADD", KEYS[1], ARGV[3], ARGV[4])
     return 'true'
else
     return 'false'
end
`)

func (s *Semaphore) Acquire(ctx context.Context) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (s *Semaphore) Release(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

func (s *Semaphore) Close() error { _ = "STUB: not implemented"; return nil }

func semaphoreKey(scope string) string { _ = "STUB: not implemented"; return "" }
