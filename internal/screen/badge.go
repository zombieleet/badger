package screen

import (
	"image/color"

	"tinygo.org/x/tinyfont"
	"tinygo.org/x/tinyfont/freesans"
)

const (
	badgeFirstName = "Victory"
	badgeLastName  = "Osikwemhe"
	badgeLinkedIn  = "victoryosikwemhe"
	badgeGithub    = "zombieleet"
)

const (
	gopherGap   = 4 // space between gopher and right edge
	leftMargin  = 8
	nameBase1   = 22 // baseline of first name (Bold12pt cap height ~17px)
	nameBase2   = 44 // baseline of last name
	dividerY    = 52
	row1Top     = 62 // top of the Linkedin row
	row2Top     = 92 // top of the Github row
	iconTextGap = 6  // space between icon and text
	contactCapH = 13 // approx. cap height of Regular9pt, for vertical centering
)

func (mItem MenuItem) DrawBadge() {
	w, h := resetDisplay()

	// Right side: gopher, vertically centered

	gx := w - gopherW - gopherGap
	gy := (h - gopherH) / 2

	drawBitmap(gx, gy, gopherW, gopherH, gopherBmp, black)

	// Left side: name on two lines + divider + contacts
	// The full name on one line is ~280px in Bold18pt, too wide next to the gopher

	tinyfont.WriteLine(&display, &freesans.Bold12pt7b, leftMargin, nameBase1, badgeFirstName, black)
	tinyfont.WriteLine(&display, &freesans.Bold12pt7b, leftMargin, nameBase2, badgeLastName, black)

	showRect(leftMargin, dividerY, gx-leftMargin-8, 2, black)

	drawContact(leftMargin, row1Top, linkedinW, linkedinH, linkedinBmp, badgeLinkedIn)
	drawContact(leftMargin, row2Top, githubW, githubH, githubBmp, badgeGithub)

	display.Display()
}

func drawContact(x, top, iw, ih int16, icon []byte, text string) {
	drawBitmap(x, top, iw, ih, icon, black)
	baseline := top + (ih+contactCapH)/2
	tinyfont.WriteLine(&display, &freesans.Regular9pt7b, x+iw+iconTextGap, baseline, text, black)

}

// drawBitmap draws a 1-bit bitmpa: row-major, MSB first, rows padded to whole bytes
// Set bits are drawn in color c; clear bits are left untouched(transparent)
func drawBitmap(x0, y0, w, h int16, data []byte, c color.RGBA) {
	stride := (w + 7) / 8

	for y := int16(0); y < h; y++ {
		for x := int16(0); x < w; x++ {
			if data[y*stride+x/8]&(0x80>>uint(x%8)) != 0 {
				display.SetPixel(x0+x, y0+y, c)
			}
		}
	}
}
