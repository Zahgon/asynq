package asynq

import (
	"context"
)

func GetTaskID(ctx context.Context) (id string, ok bool) {
	_ = "STUB: not implemented"
	return "", false
}

func GetRetryCount(ctx context.Context) (n int, ok bool) {
	_ = "STUB: not implemented"
	return 0, false
}

func GetMaxRetry(ctx context.Context) (n int, ok bool) { _ = "STUB: not implemented"; return 0, false }

func GetQueueName(ctx context.Context) (queue string, ok bool) {
	_ = "STUB: not implemented"
	return "", false
}
