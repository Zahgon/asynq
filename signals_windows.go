//go:build windows

package asynq

func (srv *Server) waitForSignals() { _ = "STUB: not implemented"; return }

func (s *Scheduler) waitForSignals() { _ = "STUB: not implemented"; return }
