package main

import (
	"fmt"
	"runtime"
	"strconv"
	"sync"
	"time"
)

/**
1.并发和并行
	1）并发：多线程程序在单核上运行；
	2）并行：多线程程序在多核上运行；
2.Go主线程和协裎
	1）协裎是轻量级线程，一个线程可以起多个协裎；
	2）有独立的栈空间
	3）共享程序堆空间
	4）调度由用户控制
*/
//协裎demo，test()表示协裎，main()表示线程；
func ConvergeDemo() {
	test := func() {
		for i := 1; i <= 10; i++ {
			fmt.Println("test() hello world! ", strconv.Itoa(i))
			time.Sleep(time.Second)
		}
	}
	go test() //开启了一个协裎
	for i := 1; i <= 10; i++ {
		fmt.Println("main() hello world! ", strconv.Itoa(i))
		time.Sleep(time.Second)
	}

}

//获取当前CPU数
func GetCpuNum() {
	cpuNum := runtime.NumCPU()
	runtime.GOMAXPROCS(cpuNum)
	fmt.Println("cpuNum:", cpuNum)
}

/**
3.不同goroutine之间如何通信
	1）全局变量的互斥锁
	2）使用管道channel来解决
*/
var (
	myMap = make(map[int]int, 10)
	/*声明一个全局互斥锁lock
	sync是包：synchornized同步
	Mutex：是互斥
	*/
	lock sync.Mutex
)

//启200个协裎计算1～200的阶乘
func Factorial() {
	test := func(n int) {
		res := 1
		//求阶乘
		// for i := 1; i <= n; i++ {
		// 	res *= i
		// }
		//求平方
		res = n * n
		lock.Lock() //加锁
		myMap[n] = res
		lock.Unlock() //解锁
	}
	for i := 1; i <= 20; i++ {
		go test(i)
	}
	//等10s，估算10秒内20个协裎会全部运行完，主线程才结束，不保险，所以引入channel
	time.Sleep(time.Second * 10)
	lock.Lock()
	for i, v := range myMap {
		fmt.Printf("map[%d]=%d\n", i, v)
	}
	lock.Unlock()
}
func main() {
	// ConvergeDemo()
	// GetCpuNum()
	Factorial()
}
