package main

import "fmt"

/**
1.channel基本介绍
	1）channel本质就是队列
	2）数据是先进先出 FIFO
	3）线程安全，多goroutine访问时，不需要加锁；即channel本身就是线程安全的；
	4）channel是有类型的，一个string类型的channel只能存放string类型的数据；
	5）示意图

2.定义/声明channel ：var 变量名 chan 数据类型
	例如:var intChan chan int  (intChan用于存放int数据)
		var mapChan chan map[int]string  (mapChan用于存放map[int]string数据)
		var perChan chan Person
		var perPtrChan chan *Person
	说明:channel是引用类型
		channel必须初始化才能写入数据，即make之后才能使用；

3.channel初始化，写入数据到管道，从管道读数据的基本注意事项 ChannelDemo()
	注意事项：1）channel只能存放指定数据
			2）channel写满数据就不能再写了，取出数据后可以再写；
			3）在没有使用协裎的情况下，已经全部取出数据的管道再取数据就会deadlock；
*/
func ChannelDemo() {
	//1.创建一个可以存储3个int的channel
	var intChan chan int
	intChan = make(chan int, 3)
	//2.打印intChan看看
	fmt.Printf("channel value:%v, channel addr:%v\n", intChan, &intChan)
	//3.向管道写数据
	intChan <- 10
	num := 200
	intChan <- num
	intChan <- -50
	// intChan <- 2 //写数据不能超过其容量

	//4.查询channel的长度和容量
	fmt.Printf("channel length:%v，channel cap:%v\n", len(intChan), cap(intChan))

	//5.从管道读取数据
	val := <-intChan
	fmt.Println("val:", val)
	fmt.Printf("channel length:%v，channel cap:%v\n", len(intChan), cap(intChan))

	//6.在没有使用协裎的情况下，已经全部取出数据的管道再取数据就会deadlock
	val2, val3 := <-intChan, <-intChan
	fmt.Printf("val2:%v,val3:%v\n", val2, val3)
	fmt.Printf("channel length:%v，channel cap:%v\n", len(intChan), cap(intChan))

	/*deadlock*/
	// val4 := <-intChan
	// fmt.Println("val:", val4)
}
func main() {
	ChannelDemo()
}
