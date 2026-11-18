package blinky

import (
	"machine"
	"time"
)

var (
	led = machine.LED
)

func Blink(count int8) {
	led.Configure(machine.PinConfig{
		Mode: machine.PinOutput,
	})
	for range count {
		led.High()
		time.Sleep(time.Millisecond * 500)
		led.Low()
		time.Sleep(time.Millisecond * 500)
	}
}
