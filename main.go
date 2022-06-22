package main

import (
	"fmt"
	"github.com/fzdwx/rcp/common"
	"github.com/fzdwx/rcp/display"
	"time"
)

func main() {

	ticker := time.NewTicker(time.Second)
	// todo
	//  1. 一个携程一直截屏（这个速率可以进行配置） 然后发送到一个channel
	// 	2. 客户订阅这个channel
	for {
		select {
		case <-ticker.C:
			img, err := display.CapMain()
			if err != nil {
				panic(err)
			}
			img.Resize()
			err = img.Save(common.RandString(4) + ".png")
			if err != nil {
				fmt.Println(err)
			}

			fmt.Println("save")
		}
	}

	//time.Sleep(3 * time.Second)
	//all, err := display.CapMain()
	//if err != nil {
	//	panic(err)
	//}
	//
	//all.Resize().Save("hello.png")
	//robotgo.Click("left")
	//key.Right.Do()
}
