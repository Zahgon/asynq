package dash

import (
	"github.com/gdamore/tcell/v2"
)

type columnAlignment int

const (
	alignRight columnAlignment = iota
	alignLeft
)

type columnConfig[V any] struct {
	name      string
	alignment columnAlignment
	displayFn func(v V) string
}

type column[V any] struct {
	*columnConfig[V]
	width int
}

func drawTable[V any](d *ScreenDrawer, style tcell.Style, configs []*columnConfig[V], data []V, highlightRowIdx int) {
	_ = "STUB: not implemented"
	return
}
