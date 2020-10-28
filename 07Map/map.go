package main

import "fmt"

func createMap() {
	/*3种声明并初始化map的方法*/
	// 第一种方法
	var scores map[string]int = map[string]int{"english": 80, "chinese": 85}
	fmt.Println(scores)

	// 第二种方法
	scores2 := map[string]int{"english": 80, "chinese": 85}
	fmt.Println(scores2)

	// 第三种方法 最常用 如果没有{}初始化，就必须make初始化
	scores3 := make(map[string]int)
	scores3["english"] = 80
	scores3["chinese"] = 85
	fmt.Println(scores3)
}
func traverseMap() {
	/*3种遍历map的方法*/
	//1.获取key和value
	scores := map[string]int{"english": 80, "chinese": 85}

	for subject, score := range scores {
		fmt.Printf("key: %s, value: %d\n", subject, score)
	}
	//2.只获取key，value不需要"_"占位
	scores2 := map[string]int{"english": 80, "chinese": 85}

	for subject := range scores2 {
		fmt.Printf("key: %s\n", subject)
	}
	//3.只获取value，需要"_"占位
	scores3 := map[string]int{"english": 80, "chinese": 85}

	for _, score := range scores3 {
		fmt.Printf("value: %d\n", score)
	}
}
func addMap() {
	scores := make(map[string]int)
	//添加key和value
	scores["english"] = 80
	scores["chinese"] = 85
	fmt.Println(scores)

}
func delMap() {
	scores := make(map[string]int)
	scores["english"] = 80
	scores["chinese"] = 85
	//删除元素，使用 delete 函数，如果 key 不存在，delete 函数会静默处理，不会报错。
	delete(scores, "english")
	fmt.Println(scores)
}
func modifyMap() {
	scores := make(map[string]int)
	//添加key和value
	scores["english"] = 80
	//修改value
	scores["english"] = 85
	fmt.Println(scores)
}
func readMap() {
	scores := make(map[string]int)
	scores["english"] = 80
	//访问元素，直接使用 [key] 即可 ，如果 key 不存在，也不报错，会返回其value-type 的零值。
	fmt.Println(scores["english"])
	fmt.Println(scores["math"])
}
func judgeKeyExist() {
	/**
	当key不存在，会返回value-type的零值 ，所以你不能通过返回的结果是否是零值来判断对应的 key 是否存在，
	因为 key 对应的 value 值可能恰好就是零值。

	其实字典的下标读取可以返回两个值，使用第二个返回值都表示对应的 key 是否存在，若存在ok为true，
	若不存在，则ok为false
	*/
	scores := map[string]int{"english": 80, "chinese": 85}
	math, ok := scores["math"]
	if ok {
		fmt.Printf("math 的值是: %d", math)
	} else {
		fmt.Println("math 不存在")
	}
}
func judgeKeyExist2() {
	scores := map[string]int{"english": 80, "chinese": 85}
	if math, ok := scores["math"]; ok {
		fmt.Printf("math 的值是: %d", math)
	} else {
		fmt.Println("math 不存在")
	}
}
func main() {
	// createMap()
	// traverseMap()
	// addMap()
	// delMap()
	// modifyMap()
	//readMap()
	judgeKeyExist()
	judgeKeyExist2()
}
