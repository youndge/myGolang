package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

/**
1. os.File封装所有文件相关操作，File是一个结构体；
*/

/**
2.打开和关闭文件
	1）func Open(name string) (file *File, err error)
		Open打开一个文件用于读取。如果操作成功，返回的文件对象方法可用于读取数据，
		对应的文件描述符具有O_RDONLY模式。如果失败，错误底层类型是*PathError；
	2）func (f *File) Close() error 关闭文件f，使文件不能用于读写，返回可能出现的错误；
*/
func openClose() {
	file, err := os.Open("test.c")
	if err != nil {
		fmt.Println("open file err=", err)
	}
	fmt.Printf("file:%v\n", file)
	err = file.Close()
	if err != nil {
		fmt.Println("Close file err=", err)
	}
}

/**
3.读文件并显示在终端
	1）带缓冲区的方式，os。Open，file.Close,bufio.NewReader(),reader.ReadString;
	2)实用ioutil一次性将整个文件读入到内存中，适用于文件不大的情况，ioutil.ReadFile;
*/
func Read1() {
	file, err := os.Open("test.c")
	if err != nil {
		fmt.Println("Open file err:", err)
	}
	defer file.Close()
	//创建一个reader
	reader := bufio.NewReader(file)
	for {
		str, err := reader.ReadString('\n')
		if err == io.EOF {
			break //io.EOF表示文件结尾
		}
		//输出内容
		fmt.Print(str)
	}
	fmt.Println("read file over!")
}
func Read2() {
	//使用ioutil.ReadFile()一次性将文件读取到位
	file := "test.c"
	content, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Printf("read file err:\n", err)
	}
	//将读取内容显示到终端
	fmt.Printf("%v", string(content))
	//没有显示的Open文件，因此也无需显示的Close
	//实际上Open和Close封装到了ReadFile函数内部
}

func Write() {
	filePath := "output.txt"
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("open file err:", err)
		return
	}
	defer file.Close()
	str := "hello world!\n"
	writer := bufio.NewWriter(file)
	writer.WriteString(str)
	writer.Flush()
}
func main() {
	// openClose()
	// Read1()
	// Read2()
	Write()
}
