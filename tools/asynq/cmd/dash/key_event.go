package dash

import (
	"time"

	"github.com/gdamore/tcell/v2"
)

type keyEventHandler struct {
	s     tcell.Screen
	state *State
	done  chan struct{}

	fetcher fetcher
	drawer  drawer

	ticker       *time.Ticker
	pollInterval time.Duration
}

func (h *keyEventHandler) quit() { _ = "STUB: not implemented"; return }

func (h *keyEventHandler) HandleKeyEvent(ev *tcell.EventKey) { _ = "STUB: not implemented"; return }

func (h *keyEventHandler) goBack() { _ = "STUB: not implemented"; return }

func (h *keyEventHandler) handleDownKey() { _ = "STUB: not implemented"; return }

func (h *keyEventHandler) downKeyQueues() { _ = "STUB: not implemented"; return }

func (h *keyEventHandler) downKeyQueueDetails() { _ = "STUB: not implemented"; return }

func (h *keyEventHandler) handleUpKey() { _ = "STUB: not implemented"; return }

func (h *keyEventHandler) upKeyQueues() { _ = "STUB: not implemented"; return }

func (h *keyEventHandler) upKeyQueueDetails() { _ = "STUB: not implemented"; return }

func (h *keyEventHandler) handleEnterKey() { _ = "STUB: not implemented"; return }

func (h *keyEventHandler) resetTicker() { _ = "STUB: not implemented"; return }

func (h *keyEventHandler) enterKeyQueues() { _ = "STUB: not implemented"; return }

func (h *keyEventHandler) enterKeyQueueDetails() { _ = "STUB: not implemented"; return }

func (h *keyEventHandler) handleLeftKey() { _ = "STUB: not implemented"; return }

func (h *keyEventHandler) handleRightKey() { _ = "STUB: not implemented"; return }

func (h *keyEventHandler) nextPage() { _ = "STUB: not implemented"; return }

func (h *keyEventHandler) prevPage() { _ = "STUB: not implemented"; return }

func (h *keyEventHandler) showHelp() { _ = "STUB: not implemented"; return }
