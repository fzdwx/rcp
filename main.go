package main

import (
	"fmt"
	"time"

	"github.com/axgle/mahonia"
	"github.com/go-vgo/robotgo"
	_ "github.com/robotn/gohook"
)

func main() {
	var enc = mahonia.NewDecoder("gbk")
	for {
		fmt.Println(enc.ConvertString(robotgo.GetTitle()))
		time.Sleep(5 * time.Second)
	}
}
