package button

import (
	"machine"
)

var (
	selectButton = machine.BUTTON_B
	selectChan   = make(chan struct{}, 1)
)

type MenuItem interface {
	DrawBadge()
	DrawWeather()
	DrawNews()
	Draw()
}

func InitSelectButton(menuItem MenuItem) {
	selectButton.Configure(machine.PinConfig{
		Mode: machine.PinInputPulldown,
	})

	selectButton.SetInterrupt(machine.PinFalling, func(pin machine.Pin) {
		select {
		case selectChan <- struct{}{}:
		default:
		}
	})

	go func() {
		for range selectChan {
			selected := menuItems[menuItemSelected]
			println(selected)
			switch selected {
			case "Badge":
				menuItem.DrawBadge()
			default:
				menuItem.Draw()
			}
		}
	}()
}
