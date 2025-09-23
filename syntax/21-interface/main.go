package main

import "fmt"

type Usber interface {
	start()
	stop()
}

type Phone struct {
	Name string
}

type Camera struct {
	Pixels int
}

type Computer struct {
}

// 手机要实现usb接口的话，就必须实现usb里面所有的方法
func (p Phone) start() {
	fmt.Println(p.Name, "启动")
}

func (p Phone) stop() {
	fmt.Println(p.Name, "关闭")
}

// 相机
func (c Camera) start() {
	fmt.Println("启动")
}

func (c Camera) stop() {
	fmt.Println("关闭")
}

func (c Camera) Photograph() {
	fmt.Println("相机开始拍照")
}

// 电脑
func (c Computer) work(usb Usber) {
	usb.start()
	usb.stop()
}

func main() {
	phone := Phone{Name: "华为手机"}
	phone.start()

	var p1 Usber = phone
	p1.start()

	var camera = Camera{Pixels: 90}
	var Uc Usber = camera
	Uc.start()
	camera.Photograph()

	var cc = Computer{}
	cc.work(camera)
	cc.work(phone)

	fmt.Println("bbbbbb")

}
