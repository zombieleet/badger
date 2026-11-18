package button

import (
	"machine"

	"tinygo.org/x/drivers/uc8151"
)

var (
	powerOnBtn     = machine.BUTTON_UP
	powerStateChan = make(chan struct{}, 1)
)

func InitPowerOnButton(display uc8151.Device) {
	powerOnBtn.Configure(machine.PinConfig{
		Mode: machine.PinInputPulldown,
	})

	powerOnBtn.SetInterrupt(machine.PinFalling, func(pin machine.Pin) {
		select {
		case powerStateChan <- struct{}{}:
		default:
		}
	})

	go func() {
		powerState := false
		for range powerStateChan {
			if powerState {
				display.PowerOff()
				continue
			}
			powerState = true
			display.PowerOn()
		}
	}()
}
