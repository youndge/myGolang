package main

import "fmt"

func main() {
	/*1.数组的声明&赋值*/
	/*方法1:*/
	var arr [3]int
	arr[0] = 1
	arr[1] = 2
	arr[2] = 3
	/*方法2:*/
	var arr2 [3]int = [3]int{1, 2, 3}
	fmt.Println(arr2)
	/*方法3:*/
	arr3 := [3]int{1, 2, 3}
	fmt.Println(arr3)
	arr33 := [4]int{3: 1}
	fmt.Println(arr33) /*打印：[0,0,0,1]*/
	/*方法4:*/
	arr4 := [...]int{1, 2, 3}
	fmt.Println(arr4)

	/*2.数组的遍历*/
	/*方法1.常规遍历*/
	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}
	/*方法2.for-range*/
	for index, value := range arr {
		fmt.Printf("index=%v,value=%v", index, value)
	}
}
