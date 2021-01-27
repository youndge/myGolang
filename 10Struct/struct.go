package main
import (
	"fmt"
	"unsafe"
)
/*
1.结构体声明：
type 结构体名称 struct {
	field1 type
	field2 type
}

举例：
type Struct struct{
	Name string
	Age int
	Score float32
}
*/
type Person struct{
	Name string
	Age int
	Score [5]float64
	ptr *int /*指针*/
	slice []int/*切片*/
	map1 map[string]string /*map*/
}
func Declare()  {
	/*创建结构体，而没有给字段赋值，则字段默认为零值*/
	var p1 Person
	fmt.Printf("%v \n",p1)
	fmt.Printf("Name:%v,Age:%v,Score:%v,ptr:%v,slice:%v,map:%v\n",
		p1.Name,p1.Age,p1.Score,p1.ptr,p1.slice,p1.map1)
	fmt.Printf("Name:%v,Age:%v,Score:%v,ptr:%v,slice:%v,map:%v\n",
		unsafe.Sizeof(p1.Name),
		unsafe.Sizeof(p1.Age),
		unsafe.Sizeof(p1.Score),
		unsafe.Sizeof(p1.ptr),
		unsafe.Sizeof(p1.slice),
		unsafe.Sizeof(p1.map1))
	fmt.Printf("Addr->Name:%p,Age:%p,Score:%p,ptr:%p,slice:%p,map:%p\n",
		&p1.Name,&p1.Age,&p1.Score,&p1.ptr,&p1.slice,&p1.map1)
	p1.Name = "Tony"
	p1.Age = 18
	p1.Score[0] = 100
	p1.ptr = &p1.Age
	
	//使用slice/map，再次说明，一定要make
	p1.slice = make([]int, 2)
	p1.slice[0] = 10

	p1.map1 = make(map[string]string)
	p1.map1["key"] = "value"
	fmt.Printf("Name:%v,Age:%v,Score:%v,ptr:%v,slice:%v,map:%v\n",
		p1.Name,p1.Age,p1.Score,p1.ptr,p1.slice,p1.map1)

}

/*
2.创建结构体
	方式1：直接声明 var person Person  见上面
	方式2：var person Person = Person{}
	方式3：var person *Person = new(Person) 返回结构体指针
	方式4：var person *Person = &Person{} 返回结构体指针
*/
func Create()  {
	//方式2：
	p2 := Person{Name:"marry",Age:20}
	fmt.Println(p2)
	//方式3：
	var p3 *Person = new(Person)
	/*p3是指针，赋值方式应该是(*p3).Name = "smith",
	但是也可以这样写p3.Name = "smith" Go的设计者为了方便，
	底层会自动把p3.Name = "smith" 转换为 (*p3).Name = "smith"
	但是不能这样写：*p3.Name = "smith",因为.比*优先级高*/
	(*p3).Name, (*p3).Age = "smith", 18
	fmt.Println(*p3)
	p3.Name, p3.Age = "john", 20
	fmt.Println(*p3)

	//方式4：var person *Person = &Person{}
	/*&Person{}和new(Person)一样，
	也可以直接赋值 var person *Person = &Person{Name:"tom",Age:20}
	*/
}
/*
3.结构体的内存分配
*/
func AssignMemory()  {
	p1 := Person{Name:"tony",Age:20}
	p2 := &p1
	fmt.Println((*p2).Age)
	fmt.Println(p2.Age)
	p2.Name = "tom"
	fmt.Printf("p2.Name:%v, p1.Name:%v\n",p2.Name,p1.Name)
	fmt.Printf("p2.Name:%v, p1.Name:%v\n",(*p2).Name,p1.Name)
	
	fmt.Printf("p1 addr:%p\n", &p1)
	fmt.Printf("p2 addr:%p, p2 value:%p\n", &p2, p2)

}
/**
4.结构体使用细节和注意事项：
	i)结构体的所有字段在内存中是连续的；FieldMemory()
	ii)结构体之间类型转换需要有完全相同的字段（字段名，字段个数，字段类型）；
	iii)结构体用type重新定义（相当于取别名），Go会认为是全新的类型，但是可以强转；
	iV)struct的每个字段上可以写上一个tag，该tag可以通过反射机制获取，常见的使用场景
	就是序列化和反序列化；
*/
type Point struct{
	x int
	y int
}
type Rect struct{
	leftUp,rightDown Point
}
type Rect2 struct{
	leftUp,rightDown *Point
}
func FieldMemory()  {
	r1 := Rect{Point{1,2},Point{3,4}}
	//r1有4个int，在内存中是连续分布的
	fmt.Printf("r1.leftUp.x addr:%p\n",&r1.leftUp.x)
	fmt.Printf("r1.leftUp.y addr:%p\n",&r1.leftUp.y)
	fmt.Printf("r1.rightDown.x addr:%p\n",&r1.rightDown.x)
	fmt.Printf("r1.rightDown.y addr:%p\n",&r1.rightDown.y)

	//r2有两个*Point类型，这个连个*Point类型的本身地址也是连续的
	//但是它们指向的地址不一定是连续的
	r2 := Rect2{&Point{10,20},&Point{30,40}}
	//本身地址
	fmt.Printf("r2.leftUp:%p,r2.rightDown:%p\n",&r2.leftUp,&r2.rightDown)
	//指向地址
	fmt.Printf("r2.leftUp:%p,r2.rightDown:%p\n",r2.leftUp,r2.rightDown)
}
func main()  {
	//Declare()
	//Create()
	//AssignMemory()
	FieldMemory()
}