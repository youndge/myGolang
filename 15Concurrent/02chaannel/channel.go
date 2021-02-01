package main

import (
	"fmt"
	"time"
)

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

/*应用实例3:统计1～200000的数字中，哪些是素数？*/
func ChanAppTest03() {
	intChan := make(chan int, 1000)
	primeChan := make(chan int, 2000)
	//标识退出的管道
	exitChan := make(chan bool, 4)

	//开启一个协裎，向intChan写入1-8000个数
	go putNum(intChan)
	//开启4个协裎，从intChan取数，是素数就写入primeChan
	for i := 0; i < 4; i++ {
		go primeNum(intChan, primeChan, exitChan)
	}

	//主线程处理
	go func() {
		for i := 0; i < 4; i++ {
			<-exitChan
		}
		//当exitChan中4个结果全部取出，即可关闭primeChan
		close(primeChan)
	}()

	//遍历primeChan，取出结果
	for {
		res, ok := <-primeChan
		if !ok {
			break
		}
		fmt.Printf("素数:%d\n", res)
	}

	fmt.Println("main 线程结束")
}
func putNum(intChan chan int) {
	for i := 1; i <= 8000; i++ {
		intChan <- i
	}
	close(intChan)
}
func primeNum(intChan chan int, primeChan chan int, exitChan chan bool) {
	var flag bool
	for {
		time.Sleep(time.Second * 2)
		num, ok := <-intChan
		if !ok {
			break
		}
		flag = true
		//判断num是否是素数
		for i := 2; i < num; i++ {
			if num%i == 0 {
				//不是素数置false
				flag = false
				break
			}
		}
		if flag {
			primeChan <- num
		}
	}
	fmt.Println("有一个primeNum协裎因为取不到数据，退出！")
	exitChan <- true
}

/**
6.channel使用细节和注意事项
	1）channel默认是双向的，但可以声明为只读，或者只写性质；
		var chanOnlyWrite chan<- int  //声明只写channel
		var chanOnlyRead <-chan int   //声明只读channel
		应用实例：OnlyRWChanApp()
	2)使用select可以解决从管道取数据的阻塞问题；SelectApp()
	3）goroutine中使用recover，解决协裎中出现的panic，导致程序崩溃问题；RecoverApp()
*/
func send(ch chan<- int, exitChan chan struct{}) {
	for i := 0; i < 10; i++ {
		ch <- i
	}
	close(ch)
	var a struct{}
	exitChan <- a
}
func recv(ch <-chan int, exitChan chan struct{}) {
	for {
		v, ok := <-ch
		if !ok {
			break
		}
		fmt.Println(v)
	}
	var a struct{}
	exitChan <- a
}
func OnlyRWChanApp() {
	ch := make(chan int, 10)
	exitChan := make(chan struct{}, 2)
	go send(ch, exitChan)
	go recv(ch, exitChan)

	total := 0
	for _ = range exitChan {
		total++
		if total == 2 {
			break
		}
	}
	fmt.Println("over...")
}

/*6.2 select解决阻塞应用实例*/
func SelectApp() {
	//1.创建一个channel int 10
	intChan := make(chan int, 10)
	for i := 0; i < 10; i++ {
		intChan <- i
	}
	//2.创建一个channel string 5
	strChan := make(chan string, 5)
	for i := 0; i < 5; i++ {
		strChan <- "hello" + fmt.Sprintf("%d", i)
	}

	//传统方法在遍历管道时，如果不关闭会阻塞导致deadlock
	//然而实际开发中我们不好确定什么时候需要关闭管道，所以引用select来解决;
	for {
		select {
		/*注意：如果intChan一直没有关闭，不会一直阻塞而deadlock，
		这里会自动匹配下一个case*/
		case v := <-intChan:
			fmt.Printf("Read data(%d) from intChan\n", v)
			time.Sleep(time.Second)
		case v := <-strChan:
			fmt.Printf("Read data(%d) from strChan\n", v)
			time.Sleep(time.Second)
		default:
			fmt.Printf("Read data fail!\n")
			time.Sleep(time.Second)
			return
		}
	}
}

/*6.3 recover应用实例*/
func sayHello() {
	for i := 0; i < 5; i++ {
		time.Sleep(time.Second)
		fmt.Println("hello world!")
	}
}
func test() {
	//这里使用defer + recover
	defer func() {
		//捕获test抛出的panic
		if err := recover(); err != nil {
			fmt.Println("test() error occurred:", err)
		}
	}()
	//定义了一个map
	myMap := make(map[int]string)
	myMap[0] = "golang" //error
}
func RecoverApp() {
	go sayHello()
	go test()

	for i := 0; i < 5; i++ {
		fmt.Sprintln("main() ok=", i)
		time.Sleep(time.Second)
	}
}
func main() {
	// ChannelDemo()
	// TraverseChan()
	// ChanAppTest01()
	// ChanAppTest03()
	// OnlyRWChanApp()
	// SelectApp()
	RecoverApp()
}
