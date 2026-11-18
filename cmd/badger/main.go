package main

import (
	"time"

	"github.com/zombieleet/badger/internal/blinky"
	"github.com/zombieleet/badger/internal/button"
	"github.com/zombieleet/badger/internal/screen"
)

func main() {

	time.Sleep(500 * time.Millisecond)

	display := screen.InitSPIDisplay()

	button.InitMenuButton(screen.MenuObject{})
	button.InitPowerOnButton(display)
	button.InitSelectButton(screen.MenuItem{})
	blinky.Blink(3)

	select {}
}
