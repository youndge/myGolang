package main

import "fmt"

/**
1.方法的声明和调用
func (recevier type) methodName(parameter list) (return value list){
	func body
	return value
}
1)recevier type:表示这个方法和type类型绑定，或者说该方法作用于type类型；
2）type可以是结构体，亦或是其他自定义类型，包括int，bool等基础数据类型；
3）reciver：就是type类型的一个变量（实例）；
4）返回值列表：返回值可以是多个；
5）return语句不是必须的；

调用：recevier.methodname()
*/
type Person struct {
	Name string
}

//Declare
func (p Person) test() {
	fmt.Println("test() name=", p.Name)
}
func (p Person) getSum(x, y int) int {
	fmt.Println("x+y=", x+y)
	return x + y
}
func DeclareAndCall() {
	p := &Person{"tom"}
	p.test()
	p.getSum(1, 2)
}

/**
2.方法的注意事项和使用细节
	1）方法的访问范围控制的规则和函数一样，即方法名首字母小写，只能在本包访问，
		方法首字母大写，可以在本包及其他包访问。
	2）如果一个类型实现了String()方法，那么fmt.Println()默认会调用这个变量的String()进行输出。
	3）...
*/

/**
3.方法和函数的区别
	1）调用方式不一样： 函数名（参数列表） /  变量。方法名（参数列表）；
	2）对于普通函数，入参为值类型时，不能将指针类型的数据直接传递，反之亦然；testFunc()
	   对于方法，接受者为值类型时，可以直接用指针类型的变量调用方法，反之也可以；
	   不管调用形式如何，真正决定值拷贝还是地址拷贝的是看这个方法和哪个类型绑定；testMethod()
*/
func testFunc() {
	test01 := func(p Person) {
		fmt.Println(p.Name)
	}
	test02 := func(p *Person) {
		fmt.Println(p.Name)
	}
	person := Person{"tom"}
	test01(person)  //test01(*person)错误
	test02(&person) //test02(person)错误
}

func (p Person) test03() {
	p.Name = "jack"
	fmt.Println("test03=", p.Name)
}

func (p *Person) test04() {
	p.Name = "mary"
	fmt.Println("test04=", p.Name)
}
func testMethod() {
	p := Person{"tom"}
	p.test03()
	fmt.Println("testmethod() p.Name=", p.Name)
	(&p).test03() //形式上是传地址，单本质还是值拷贝
	fmt.Println("testmethod() p.Name=", p.Name)

	(&p).test04() //这个就是地址拷贝
	fmt.Println("testmethod() p.Name=", p.Name)
	p.test04() //等价于(&p).test04()，底层自动转换为地址拷贝，方便程序员写代码
	fmt.Println("testmethod() p.Name=", p.Name)
}
func main() {
	// DeclareAndCall()
	// testFunc()
	testMethod()
}
