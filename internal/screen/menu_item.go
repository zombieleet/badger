package screen

type MenuItem struct{}

func (mItem MenuItem) Draw() {
	x, y := resetDisplay()
	println(x, y)
	display.Display()
}
