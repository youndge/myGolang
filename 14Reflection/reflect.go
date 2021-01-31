package main

import (
	"fmt"
	"reflect"
)

/**
1.基本介绍
	1）使用反射，需要import “reflect”；
	2）反射可以在运行时动态获取变量的各种信息，比如变量的类型type，类别kind；
	3）如果是结构体变量，还可以获取到结构体本身的信息（包括结构体的字段，方法）；
	4）通过反射可以修改变量的值，可以调用关联的方法；
*/

/**
2.反射的应用场景
	1）函数适配器，不知道接口调用哪个函数，需要根据入参在运行时确定调用哪个接口；
		func bridge(funcPtr Interface{}, args ...interface{})
	2)对结构体序列化时，如果结构体有指定Tag，也会使用到反射生成对应的字符串；
*/

/**
3.反射的重要函数
	1）reflect.TypeOf(变量名)，获取变量的类型，返回reflect.Type类型；
	2）reflect.ValueOf(变量名)，获取变量的值，返回reflect.Value类型，
		reflect.Value是一个结构体类型；
	3）变量，interface{}，reflect.Value是可以相互转换的；
*/
//Demo01:演示对基本数据类型，interface{}，reflect.Value进行反射的基本操作；
func Demo01(b interface{}) {
	//通过反射获取传入的变量的type，kind，值
	//1.先获取到refflect.Type
	rType := reflect.TypeOf(b)
	fmt.Println("rType:", rType)
	//2.获取到reflect.Value
	rVal := reflect.ValueOf(b)
	fmt.Println("rVal:", rVal.Int())
	//将rVal转换成interface{}
	iV := rVal.Interface()
	//将interface{}通过断言转换成需要的类型
	num := iV.(int)
	fmt.Println("num:", num)
}

//Demo01:演示对结构体类型，interface{}，reflect.Value进行反射的基本操作；
func Demo02(b interface{}) {
	//通过反射获取传入的变量的type，kind，值
	//1.先获取到refflect.Type
	rType := reflect.TypeOf(b)
	fmt.Println("rType:", rType)
	//2.获取到reflect.Value
	rVal := reflect.ValueOf(b)

	//将rVal转换成interface{}
	iV := rVal.Interface()
	//将interface{}通过断言转换成需要的类型
	fmt.Printf("iv:%v,iv type:%T\n", iV, iV)
	stu, ok := iV.(Student)
	if ok {
		fmt.Printf("stu.Name:%v\n", stu.Name)
	}
}

type Student struct {
	Name string
	Age  int
}

func Demo() {
	num := 100
	Demo01(num)
	stu := Student{"tom", 20}
	Demo02(stu)
}

/**
4.注意事项和细节
	1）reflect包；
	2）reflect.Value.Kind获取变量的类别，返回的是一个常量；
		Type和Kind的区别，可能相同也可能不同；
	3）通过反射可以让变量在interface{}和reflect.Value之间转换；
	4）使用反射来获取变量的值并返回对应的类型，要求数据类型匹配，比如var是int，
		就应该使用reflect.Value(var).Int(),否则会报panic；
	5）通过反射来修改变量，注意使用SetXxx方法来设置需要通过对应的指针类型来完成
		这样才能改变传入变量的值，同时需要使用到reflect.Value.Elem()方法；
		ValueElemSet()

*/
func ValueElemSet(b interface{}) {
	val := reflect.ValueOf(b)
	fmt.Printf("val type:%T\n", val)
	val.Elem().SetInt(110)
	fmt.Printf("val:%V\n", val)
}

/**
5.最佳实践
	1）使用反射来遍历结构体的字段，调用结构体的方法，并获取结构体标签的值；
		ReflectStructure()
*/
type Monster struct {
	Name  string  `json:"name"`
	Age   int     `json:"monster_age"`
	Score float32 `json:"成绩"`
	Sex   string
}

//方法
func (m Monster) GetSum(n1, n2 int) int {
	return n1 + n2
}
func (m Monster) Set(name, sex string, age int, score float32) {
	m.Name, m.Age = name, age
	m.Score, m.Sex = score, sex
}
func (m Monster) Print() {
	fmt.Println("------start~------")
	fmt.Println(m)
	fmt.Println("------end~------")
}
func ReflectStruct(b interface{}) {
	typ := reflect.TypeOf(b)
	val := reflect.ValueOf(b)
	kd := val.Kind()
	if kd != reflect.Struct {
		fmt.Println("Expect struct!")
		return
	}
	//获取该struct有几个字段
	fieldNum := val.NumField()
	fmt.Println("struct has fieldnum:", fieldNum)
	//遍历每个字段
	for i := 0; i < fieldNum; i++ {
		fmt.Printf("Field No.%d,value:%v\n", i, val.Field(i))
		//获取struct的标签，需要通过reflect.Type来获取tag的值
		tagVal := typ.Field(i).Tag.Get("json")
		if tagVal != "" {
			fmt.Printf("Field No.%d, tag:%v\n", i, tagVal)
		}
	}
	//获取该struct有几个方法
	methodNum := val.NumMethod()
	fmt.Println("struct has methodnum:", methodNum)
	//方法的排序默认是按照函数名的排序（ASCII码）
	val.Method(1).Call(nil) //获取并调用第二个方法

	//调用结构体的第一个方法Method(0)
	var params []reflect.Value //声明[]reflect.Value
	params = append(params, reflect.ValueOf(10))
	params = append(params, reflect.ValueOf(40))
	res := val.Method(0).Call(params) //传入[]reflect.Value，返回[]reflect.Value
	fmt.Println("res:", res[0].Int())
}
func main() {
	// Demo()

	// val := 10
	// ValueElemSet(&val)
	// fmt.Println(val)

	m := Monster{
		Name:  "tony",
		Age:   23,
		Score: 45.5,
	}
	ReflectStruct(m)
}
