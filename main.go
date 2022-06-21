package main

import "github.com/fzdwx/rcp/ro"

func main() {
	//fmt.Println(robotgo.GetMainDPI())
	//
	//fmt.Println(screenshot.NumActiveDisplays())
	//
	//bitMap := ro.Cap(0, 0, 1920, 1080)
	//bitMap.SaveToPNG("qwe.png")
	all, err := ro.CapMain()
	if err != nil {
		panic(err)
	}

	all.SaveToPNG("hello.png")
}
