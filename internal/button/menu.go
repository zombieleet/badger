package button

import (
	"machine"
)

var (
	menuButton = machine.BUTTON_A
	menuChan   = make(chan struct{}, 1)

	menuItemSelected = 0

	menuItems = []string{
		"Badge",
	}
)

type MenuObject interface {
	Draw(menuItems []string, selected int)
}

func InitMenuButton(menu MenuObject) {

	menuButton.Configure(machine.PinConfig{
		Mode: machine.PinInputPulldown,
	})

	menuButton.SetInterrupt(machine.PinRising, func(pin machine.Pin) {
		select {
		case menuChan <- struct{}{}:
		default:
		}
	})

	go func() {
		btnPressed := false
		for range menuChan {
			if btnPressed {
				menuItemSelected += 1
				if len(menuItems) <= menuItemSelected {
					menuItemSelected = 0
				}
				menu.Draw(menuItems, menuItemSelected)
				continue
			}
			btnPressed = true
			menu.Draw(menuItems, menuItemSelected)
		}
	}()

}
