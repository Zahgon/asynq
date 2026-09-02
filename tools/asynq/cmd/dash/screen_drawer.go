package dash

import (
	"github.com/gdamore/tcell/v2"
)

type ScreenDrawer struct {
	l *LineDrawer
}

func NewScreenDrawer(s tcell.Screen) *ScreenDrawer { _ = "STUB: not implemented"; return nil }

func (d *ScreenDrawer) Print(s string, style tcell.Style) { _ = "STUB: not implemented"; return }

func (d *ScreenDrawer) Println(s string, style tcell.Style) { _ = "STUB: not implemented"; return }

func (d *ScreenDrawer) FillLine(r rune, style tcell.Style) { _ = "STUB: not implemented"; return }

func (d *ScreenDrawer) FillUntil(r rune, style tcell.Style, limit int) {
	_ = "STUB: not implemented"
	return
}

func (d *ScreenDrawer) NL() { _ = "STUB: not implemented"; return }

func (d *ScreenDrawer) Screen() tcell.Screen { _ = "STUB: not implemented"; return *new(tcell.Screen) }

func (d *ScreenDrawer) Goto(x, y int) { _ = "STUB: not implemented"; return }

func (d *ScreenDrawer) GoToBottom() { _ = "STUB: not implemented"; return }

type LineDrawer struct {
	s   tcell.Screen
	row int
	col int
}

func NewLineDrawer(row int, s tcell.Screen) *LineDrawer { _ = "STUB: not implemented"; return nil }

func (d *LineDrawer) Draw(s string, style tcell.Style) { _ = "STUB: not implemented"; return }
