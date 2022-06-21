package key

import "github.com/go-vgo/robotgo"

type (
	// Event is a keyboard or mouse event.
	Event interface {

		// Key source the key to be pressed.
		Key() string

		// Do is a method to press the key.
		Do()
	}

	Mouse string
)

const (
	Left       = Mouse("left")
	Center     = Mouse("center")
	Right      = Mouse("right")
	WheelUp    = Mouse("wheelUp")
	WheelDown  = Mouse("wheelDown")
	WheelLeft  = Mouse("wheelLeft")
	WheelRight = Mouse("wheelRight")
)

func (m Mouse) Key() string {
	return string(m)
}

func (m Mouse) Do() {
	robotgo.Click(m.Key())
}
