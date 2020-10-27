package main

import "fmt"

func createSlice() {
	/*创建切片的4种方法*/
	/*1.引用数组*/
	myarr := [5]int{1, 2, 3, 4, 5}
	fmt.Printf("myarr 的长度为：%d，容量为：%d\n", len(myarr), cap(myarr))

	mysli := myarr[1:3]
	fmt.Printf("mysli 的长度为：%d，容量为：%d\n", len(mysli), cap(mysli))

	fmt.Println(mysli)
	/*2.常规声明*/
	// 声明字符串切片
	var strList []string
	fmt.Printf("mysli 的长度为：%d，容量为：%d\n", len(strList), cap(strList))
	// 声明整型切片同时赋值
	var numList []int = []int{1, 2, 3}
	fmt.Printf("mysli 的长度为：%d，容量为：%d\n", len(numList), cap(numList))
	//简化声明
	numList2 := []int{1, 2, 3, 4}
	fmt.Printf("mysli 的长度为：%d，容量为：%d\n", len(numList2), cap(numList2))
	// 声明一个空切片
	var numListEmpty = []int{}
	fmt.Printf("mysli 的长度为：%d，容量为：%d\n", len(numListEmpty), cap(numListEmpty))

	/*3.make函数  make 函数的格式：make( []Type, size, cap )*/
	a := make([]int, 2)
	b := make([]int, 2, 10)
	fmt.Println(a, b)
	fmt.Println(len(a), len(b))
	fmt.Println(cap(a), cap(b))
	/*4.同数组一样偷懒的方法*/
	c := []int{4: 2}
	fmt.Println(c)
	fmt.Println(len(c), cap(c))
}

/*切片的遍历*/
func traverseSlice() {
	s := []int{1, 2, 3}
	/*方法1.常规遍历*/
	for i := 0; i < len(s); i++ {
		fmt.Println(s[i])
	}
	/*方法2.for-range*/
	for index, value := range s {
		fmt.Printf("index=%v,value=%v", index, value)
	}
}

/*切片的添加*/
func addSlice() {
	//内置函数append可以动态添加切片
	slice1 := []int{100, 200, 300}
	fmt.Printf("len:%v,cap:%v\n", len(slice1), cap(slice1))
	//追加元素
	slice1 = append(slice1, 400, 500, 600)
	fmt.Printf("len:%v,cap:%v\n", len(slice1), cap(slice1))
	//切片追加切片
	slice1 = append(slice1, slice1...)
	fmt.Printf("len:%v,cap:%v\n", len(slice1), cap(slice1))

}

/*切片的删除*/
func delSlice() {
	slice2 := []int{1, 2, 3, 4, 5}
	//删除切片中第3个元素
	slice2 = append(slice2[:2], slice2[3:]...)
	fmt.Printf("del idx2 elem:%v\n", slice2)
}

/*切片的拷贝*/
func copySlice() {
	slice3 := []int{1, 2, 3, 4}
	slice4 := make([]int, 2)
	copy(slice4, slice3)
	fmt.Printf("slice3:%v,slice4:%v\n", slice3, slice4)
}

/*二维切片*/
func twoDimensionalSlice() {
	tSlice := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	fmt.Println(tSlice)
}
func main() {
	//createSlice()
	//traverseSlice()
	//addSlice()
	//delSlice()
	//copySlice()
	twoDimensionalSlice()
}
