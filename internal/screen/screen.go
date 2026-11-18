package screen

import (
	"image/color"
	"machine"
	"tinygo.org/x/drivers"
	"tinygo.org/x/drivers/uc8151"
)

const logo = `
#   #  ####  #      ###   ###  #   #  ####
#   #  #     #     #     #   # ## ##  #
# # #  ###   #     #     #   # # # #  ###
## ##  #     #     #     #   # #   #  #
#   #  ####  ####   ###   ###  #   #  ####
`

var (
	spi     = machine.SPI0
	display uc8151.Device
	black   = color.RGBA{R: 1, G: 1, B: 1, A: 255}
	white   = color.RGBA{R: 0, G: 0, B: 0, A: 255}
)

func InitSPIDisplay() uc8151.Device {
	spi.Configure(machine.SPIConfig{
		Frequency: 12 * machine.MHz,
		SCK:       machine.EPD_SCK_PIN,
		SDO:       machine.EPD_SDO_PIN,
	})

	display = uc8151.New(spi, machine.EPD_CS_PIN, machine.EPD_DC_PIN, machine.EPD_RESET_PIN, machine.EPD_BUSY_PIN)

	display.Configure(
		uc8151.Config(uc8151.Config{
			Rotation:    drivers.Rotation270,
			Speed:       uc8151.TURBO,
			FlickerFree: true,
			Blocking:    true,
		}),
	)

	go welcome()

	return display
}

func showRect(x int16, y int16, w int16, h int16, c color.RGBA) {
	for i := x; i < x+w; i++ {
		for j := y; j < y+h; j++ {
			display.SetPixel(i, j, c)
		}
	}
}

func resetDisplay() (int16, int16) {
	display.ClearBuffer()
	return display.Size()
}

func welcome() {
	w, h := resetDisplay()
	showRect(0, 0, w, h, white)

	const cols, rows, scale = 42, 5, 6
	x := (w - cols*scale) / 2
	y := (h-rows*scale)/2 - scale

	drawPixelArt(logo, x, y, scale, black)

	display.Display()
}

func drawPixelArt(art string, x0, y0, scale int16, c color.RGBA) {
	x, y := x0, y0

	for _, r := range art {
		switch r {
		case '\n':
			x = x0
			y += scale
			continue
		case '#':
			showRect(x, y, scale, scale, c)
		}
		x += scale
	}
}
