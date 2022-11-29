package main

import "fmt"

// 定义一个接口
type USB interface {
	oInput()
	oOutput()
}

type Mouse struct {
	name string
}

// 结构体实现了接口中的所有方法，就会自动实现该接口，代码解耦
func (mouse Mouse) oInput() {
	fmt.Println(mouse.name, "鼠标输入")
}
func (mouse Mouse) oOutput() {
	fmt.Println(mouse.name, "鼠标输出")
}

// 结构体实现了接口中的所有方法，就会自动实现该接口，代码解耦
type Keyboard struct {
	name string
}

func (keyboard Keyboard) oInput() {
	fmt.Println(keyboard.name, "键盘输入")
}
func (keyboard Keyboard) oOutput() {
	fmt.Println(keyboard.name, "键盘输出")
}

func test(u USB) {
	u.oInput()
	u.oOutput()
}

func main() {

	mouse := Mouse{name: "罗技"}

	keyboard := Keyboard{
		name: "雷蛇",
	}

	test(mouse)
	test(keyboard)

	var usb USB
	usb = mouse
	//fmt.Println(usb.name)
	usb.oInput()
	usb.oOutput()

}
