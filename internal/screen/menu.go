package screen

import (
	"tinygo.org/x/tinyfont"
	"tinygo.org/x/tinyfont/freesans"
)

type MenuObject struct{}

var (
	menuFont = &freesans.Bold9pt7b
)

const (
	rowH       = 21
	capH       = 13
	padLeft    = 10
	scrollBarW = 4
)

func (m MenuObject) Draw(items []string, selected int) {
	w, h := resetDisplay()

	if len(items) == 0 {
		display.Display()
		return
	}

	if selected < 0 {
		selected = 0
	}

	if selected >= len(items) {
		selected = 0
	}

	//how many rows fit, and which item is at the top (scrolling)

	visible := int(h / rowH)
	first := 0

	if selected >= visible {
		first = selected - visible + 1
	}

	for row := 0; row < visible && first+row < len(items); row++ {
		idx := first + row
		top := int16(row) * rowH
		baseline := top + (rowH+capH)/2

		fg := black

		if idx == selected {
			showRect(0, top, w, rowH, black)
			fg = white
		}
		tinyfont.WriteLine(&display, menuFont, padLeft, baseline, items[idx], fg)
	}

	if len(items) > visible {
		barH := h * int16(visible) / int16(len(items))
		barY := h * int16(first) / int16(len(items))
		showRect(w-scrollBarW, barY, scrollBarW, barH, black)
	}

	display.Display()
}
