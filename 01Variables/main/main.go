package main

import (
	"fmt"
	"strconv"
)

// func main() {
// 	var year float64 = 20.19
// 	fmt.Printf("year's Type is %T, Size:%v bytes\n", year, unsafe.Sizeof(year))
// }

// func main() {
// 	var num int64 = 4567
// 	str := strconv.Itoa(int(num))
// 	fmt.Printf("str type %T str = %q\n", str, str)
// }

func main() {
	var boolVar bool
	var strBool string = "true"
	boolVar, _ = strconv.ParseBool(strBool)
	fmt.Printf("boolVar's type is %T, boolVar=%v\n", boolVar, boolVar)

	var int32Var int
	var int64Var int64
	var strInt string = "20191001"
	int64Var, _ = strconv.ParseInt(strInt, 10, 64)
	int32Var = int(int64Var)
	fmt.Printf("int64Var's type is %T, int64Var=%v\n", int64Var, int64Var)
	fmt.Printf("int32Var's type is %T, int32Var=%v\n", int32Var, int32Var)

	var floatVar float64
	var strFloat string = "2019.1001"
	floatVar,_=strconv.ParseFloat(strFloat,64)
	fmt.Printf("floatVar's type is %T, floatVar=%v\n", floatVar, floatVar)
}
