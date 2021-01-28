package main

import (
	"fmt"
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

func main() {
	openClose()
}
