package main

import (
	"fmt"

	"../utils"
)

//1.包的使用
func packageExample() {
	/*utils包的使用*/
	var operator byte = '+'
	n1, n2 := 1.2, 1.3
	fmt.Println(utils.Cal(n1, n2, operator))
}

//2.return语句
/*Go支持返回多个值
1.如果返回多个值，在接收时希望忽略某个返回值，则使用_符合表示占位忽略；
2.如果返回值只有一个，（返回值类型列表）可以不写（）；
*/
func multiReturn(n1, n2 int) (int, int) {
	sum := n1 + n2
	sub := n1 - n2
	return sum, sub
}
func returnExample() {
	res1, res2 := multiReturn(1, 2)
	fmt.Printf("res1=%d,res2=%d\n", res1, res2)
}
func main() {
	//packageExample()
	returnExample()
}
