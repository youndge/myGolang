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

/**
4.channel的遍历和关闭
	1）内置close函数可以关闭channel，关闭的channel不能再写数据，可以读数据；
	2）channel支持for...:= range方式进行遍历，如果遍历时channel没有关闭，
		则会出现deadlock错误；
*/
func TraverseChan() {
	intChan := make(chan int, 10)
	for i := 0; i < 10; i++ {
		intChan <- i * 2
	}
	fmt.Printf("channel length:%v，channel cap:%v\n", len(intChan), cap(intChan))
	close(intChan) //关闭channel
	// for i:=0;i<len(intChan);i++{}普通for循环不能遍历channel
	for v := range intChan {
		fmt.Println("v:", v)
	}
}

/**
5.channel应用实例
	1）goroutine和channel协同工作的案例：ChanAppTest01()
		i.开启一个writeData协裎，向管道intChan写入50个数据；
		ii.开启一个readData协裎，从管道intChan中读取WriteData写入的数据；
		iii.注意：writeData和readData操作的是同一个管道；
		iv.主线程需要等待writeData和readData协裎都完成工作才能退出管道；
*/
func ChanAppTest01() {
	//创建两个管道
	intChan := make(chan int, 10)
	exitChan := make(chan bool, 1)

	//开启线程
	go writeData(intChan)
	go readData(intChan, exitChan)
	//如果exitChan读出true表示协裎运行完，退出
	for {
		_, ok := <-exitChan
		if !ok {
			break
		}
	}

}
func writeData(intChan chan int) {
	for i := 1; i <= 10; i++ { //注释掉go readData，写入数据扩大到50条，会阻塞而deadlock；
		intChan <- i
		fmt.Println("writeData:", i)
	}
	close(intChan)
}
func readData(intChan chan int, exitChan chan bool) {
	for {
		v, ok := <-intChan
		if !ok {
			break
		}
		fmt.Println("readData:", v)
	}
	//读完50个数据，写入exitChan为true
	exitChan <- true
	close(exitChan)
	fmt.Println("***readData complete!")
}
func main() {
	// ChannelDemo()
	// TraverseChan()
	ChanAppTest01()
}
