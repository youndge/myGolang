package main

import "fmt"

/**
1)指针变量存的是地址，地址指向的空间存的才是值；
*/
func showPointer()  {
	var i = 10
	var ptr *int = &i
	fmt.Printf("i = %v, addr = %v, ptr = %v\n", i, &i, ptr)
	fmt.Printf("ptr = %v, ptr addr = %v, ptr value = %v\n",ptr, &ptr,*ptr)
}

func main()  {
	showPointer()
}