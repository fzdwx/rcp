package main

import "github.com/disintegration/imaging"

func main() {

	image, _ := imaging.Open("hello.png")

	resize := imaging.Resize(image, image.Bounds().Dx()*4, 0, imaging.Lanczos)
	imaging.Save(resize, "hello_resize.png")

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
