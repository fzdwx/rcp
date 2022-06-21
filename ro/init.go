package ro

import "github.com/kbinani/screenshot"

var (
	ActiveDisplayNums = 0
)

func init() {
	ActiveDisplayNums = screenshot.NumActiveDisplays()
}
