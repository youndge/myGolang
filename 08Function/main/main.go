package main

import (
	"fmt"
	"strconv"

	"../utils"
)

var (
	//使用方式3:全局匿名函数，将匿名函数赋值给一个全局变量
	Func1 = func(n1, n2 int) int {
		return n1 * n2
	}
)

//1.包的使用
func packageExample() {
	/*utils包的使用*/
	var operator byte = '+'
	n1, n2 := 1.2, 1.3
	fmt.Println(utils.Cal(n1, n2, operator))
}

//2.return语句
/*Go支持返回多个值
1.如果返回多个值，在接收时希望忽略某个返回值，则使用_符合表示占位忽略；
2.如果返回值只有一个，（返回值类型列表）可以不写（）；
*/
func multiReturn(n1, n2 int) (int, int) {
	sum := n1 + n2
	sub := n1 - n2
	return sum, sub
}
func returnExample() {
	res1, res2 := multiReturn(1, 2)
	fmt.Printf("res1=%d,res2=%d\n", res1, res2)
}

//3.递归

//4.init函数
/*
每一个源文件都可以包含一个init函数，该函数在main函数执行前被Go运行框架调用；
init函数细节：
1.一个源文件内的执行顺序：全局变量>init()>main();
2.init函数作用，就是完成一些初始化工作；
3.该源文件import的文件的变量和init函数先于本文件的变量和init函数执行；
*/
func init() {
	fmt.Println("Call init func before main func.")
}

//5.匿名函数
/*
Go支持匿名函数，就是没有名字的函数；如果希望某个函数只调用一次，则考虑使用匿名函数；当然也可以多次调用；
*/
func anonymousFunc() {
	//使用方式1:定义即调用，只能调用一次；
	res := func(n1, n2 int) int {
		return n1 + n2
	}(10, 20)
	fmt.Println(res)
	//使用方式2:将匿名函数赋值给一个变量，实现多次调用；
	a := func(n1, n2 int) int {
		return n1 + n2
	}
	res = a(10, 20)
	fmt.Println(res)
	res = a(90, 30)
	fmt.Println(res)
	//使用方式3:全局匿名函数，将匿名函数赋值给一个全局变量
	res = Func1(5, 5)
	fmt.Println(res)
}

//6.闭包
/*
基本介绍：闭包就是一个函数和与其相关的引用环境组合的一个整体(实体)；
1.AddUpper是一个函数，返回的数据类型是func (int) int;
2.说明：AddUpper返回一个匿名函数，匿名函数引用到函数外的n，该匿名函数和n形成的整体就是一个闭包；
3.可以这样理解闭包，闭包是类，函数是操作，n是字段；函数和n构成闭包；
4.
*/
//累加器
func AddUpper() func(int) int {
	n := 10
	return func(x int) int {
		n += x
		return n
	}
}

//使用闭包
func ClosureFunc() {
	f := AddUpper()
	fmt.Println(f(1)) //11
	fmt.Println(f(2)) //13
	fmt.Println(f(3)) //16
}

//7.defer延时机制
/*
在函数中，程序员经常需要创建资源，如数据库连接，文件句柄，锁等，为了在执行完毕后，及时释放资源，Go提供了defer；
*/
func deferFunc(n1, n2 int) {
	//1.当执行到defer时，暂时不执行，会将defer后面的语句压入独立的栈（defer栈）
	//2.当函数执行完毕后，再从defer栈，按照先入后出的方式出栈；
	defer fmt.Println("3. n1=", n1)
	defer fmt.Println("2. n1=", n2)
	res := n1 + n2
	fmt.Println("1. res=", res)
}

//8.值传递&引用传递&指针传递

//9.字符串常用系统函数
func strFuncs() {
	//1.len(str)
	//2.字符串遍历，同时处理有中文的问题 r:=[]rune(str)
	strTraverse := func() {
		str := "hello北京"
		r := []rune(str)
		for i := 0; i < len(r); i++ {
			fmt.Printf("字符=%c\n", r[i])
		}
	}
	strTraverse()
	//3.字符串转整数  num, err := strconv.Atoi("12")
	strConvertToInt := func() {
		num, err := strconv.Atoi("12")
		if nil != err {
			fmt.Println("convert error!")
		} else {
			fmt.Println("convert success! num=%v", num)
		}
	}
	strConvertToInt()
	//4.正数转字符串
	intConvertTo := func() {
		str := strconv.Itoa(12345)
		fmt.Printf("str=%v,str's type is %T\n", str, str)
	}
	intConvertTo()
}
func main() {
	//packageExample()
	//returnExample()
	//anonymousFunc()
	//ClosureFunc()
	//deferFunc(10, 20)
	strFuncs()
}
