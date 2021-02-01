package main

import "fmt"

/**
1.Demo
*/
//声明一个USB接口
type USB interface {
	//声明没有实现的方法
	Start()
	Stop()
}
type Phone struct {
}

//让Phone实现USB接口的方法
func (p Phone) Start() {
	fmt.Println("Phone Start!")
}
func (p Phone) Stop() {
	fmt.Println("Phone Stop!")
}

type Camera struct {
}

//让Phone实现USB接口的方法
func (c Camera) Start() {
	fmt.Println("Camera Start!")
}
func (c Camera) Stop() {
	fmt.Println("Camera Stop!")
}

type Computer struct {
}

func (c Computer) Working(usb USB) {
	usb.Start()
	usb.Stop()
}
func Demo() {
	computer := Computer{}
	phone := Phone{}
	camera := Camera{}

	computer.Working(phone)
	computer.Working((camera))
}

/**
2.语法即细节
    语法：type 接口名 interface{
                    Method1(参数列表) 返回值列表
                    Method2(参数列表) 返回值列表
                    ...
                }
    interface不能包含任何变量；
    interface类型可以定义一组方法，但是不需要实现,即不包括方法体；
    interface本身不能创建实例，但可以指向实现了该接口的自定义类型的实例；
    一个自定义类型只有实现了某个接口，才能将该自定义类型的实例赋值给该接口类型；
    一个自定义类型实现了某个接口所有的方法才称之为实现了该接口；
    自定义数据类型都可以实现接口，包括int等基础数据类型；
    一个自定义数据类型可以实现多个接口；
    一个接口（比如A接口）可以继承多个别的接口（比如B，C接口），这时如果要实现A接口，也必须将B，C的方法全部实现；
    interface类型默认是一个指针（引用类型）如果没有对interface初始化就使用，会输出nil；
    空接口interface{}没有任何方法，即所有类型都实现了空接口，即任何一个变量都可以赋值给空接口；
*/

// func main() {
// 	Demo()
// }
