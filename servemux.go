package asynq

import (
	"context"
	"errors"
	"sync"
)

var ErrHandlerNotFound = errors.New("handler not found for task")

type ServeMux struct {
	mu  sync.RWMutex
	m   map[string]muxEntry
	es  []muxEntry
	mws []MiddlewareFunc
}

type muxEntry struct {
	h       Handler
	pattern string
}

type MiddlewareFunc func(Handler) Handler

func NewServeMux() *ServeMux { _ = "STUB: not implemented"; return nil }

func (mux *ServeMux) ProcessTask(ctx context.Context, task *Task) error {
	_ = "STUB: not implemented"
	return nil
}

func (mux *ServeMux) Handler(t *Task) (h Handler, pattern string) {
	_ = "STUB: not implemented"
	return *new(Handler), ""
}

func (mux *ServeMux) match(typename string) (h Handler, pattern string) {
	_ = "STUB: not implemented"
	return *new(Handler), ""
}

func (mux *ServeMux) Handle(pattern string, handler Handler) { _ = "STUB: not implemented"; return }

func appendSorted(es []muxEntry, e muxEntry) []muxEntry { _ = "STUB: not implemented"; return nil }

func (mux *ServeMux) HandleFunc(pattern string, handler func(context.Context, *Task) error) {
	_ = "STUB: not implemented"
	return
}

func (mux *ServeMux) Use(mws ...MiddlewareFunc) { _ = "STUB: not implemented"; return }

func NotFound(ctx context.Context, task *Task) error { _ = "STUB: not implemented"; return nil }

func NotFoundHandler() Handler { _ = "STUB: not implemented"; return *new(Handler) }
