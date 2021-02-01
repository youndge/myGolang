package main

import "fmt"

/*
1.类型断言(Type Assertion）
	基本介绍：类型断言，由于接口是一般类型，不知道具体类型，
		如果要转换成具体类型，就需要使用类型断言。
	在进行类型断言时，如果类型不匹配，就会报panic，因此使用时，要确保原来的
		空接口指向的就是断言类型；否则就带上检测机制，成功就返回ok，反之也不
		报panic，见TypeAssertionShow03();
*/
type Point struct {
	x int
	y int
}

/*1.1 b := a.(Point)就是类型断言，表示判断a是否是指向Point类型的变量，
如果是就转彻骨Point并赋值给b，否则报错*/
func TypeAssertionShow01() {
	var a interface{}
	point := Point{1, 2}
	a = point      //interface{}可以为左值
	b := a.(Point) //interfa{}为右值需要断言
	fmt.Printf("b's type is %T, b = %v\n", b, b)

}
func TypeAssertionShow02() {
	var x interface{}
	b := 3.14
	x = b
	y := x.(float64) //x.(float32就会panic)
	fmt.Printf("y's type is %T, y = %v\n", y, y)
}
func TypeAssertionShow03() {
	var x interface{}
	b := 3.14
	x = b
	/*带检测的类型断言*/
	if y, ok := x.(float64); ok {
		fmt.Printf("Convert Success!-> y's type is %T, y = %v\n", y, y)
	} else {
		fmt.Println("Convert Fail!")
	}
}

/**
2.类型断言最佳实践
	1）写一个函数，判断传入参数类型；typeAsApp01()
*/

/*2.1 TypeJudge()*/
func TypeJudge(items ...interface{}) {
	for idx, x := range items {
		switch x.(type) {
		case bool:
			fmt.Printf("第%v个参数是bool类型，值为%v\n", idx, x)
		case float32:
			fmt.Printf("第%v个参数是float32类型，值为%v\n", idx, x)
		case float64:
			fmt.Printf("第%v个参数是float64类型，值为%v\n", idx, x)
		case int, int32, int64:
			fmt.Printf("第%v个参数是int类型，值为%v\n", idx, x)
		case string:
			fmt.Printf("第%v个参数是string类型，值为%v\n", idx, x)
		case Point:
			fmt.Printf("第%v个参数是Point类型，值为%v\n", idx, x)
		case *Point:
			fmt.Printf("第%v个参数是*Point类型，值为%v\n", idx, x)
		default:
			fmt.Printf("第%v个参数类型异常，值为%v\n", idx, x)
		}
	}
}
func TypeAsApp01() {
	p := Point{1, 2}
	TypeJudge(true, float32(3.14), 9.90, 250, "youndge", p, &p)
}
func main() {
	TypeAssertionShow01()
	TypeAssertionShow02()
	TypeAsApp01()
}
