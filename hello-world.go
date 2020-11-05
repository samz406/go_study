package main

import (
	"fmt"
	"strconv"
	"strings"
	"unsafe"
)

const IP = "127.0.0.1"

func main() {

	//打印类型
	var a = 223
	fmt.Printf("the var %d type is %T\n", a, a)

	//int
	var i8 int8 = 44
	var i16 int16 = 3
	var i32 int32 = 3
	var i64 int64 = 3444555555554444455

	fmt.Println(i8, i16, i32, i64, unsafe.Sizeof(i64))

	var f32 float32 = 3.14

	fmt.Printf(" float is %v-- %f \n", f32, f32)

	var f64 = 3.14158733222
	//输出指定小数位
	fmt.Printf(" float is %v 输出小数点后四位 %.4f \n", f64, f64)

	var userName, age = getUserInfo()

	fmt.Println("userName", userName, "age", age)
	//匿名，忽略参数
	var user, _ = getUserInfo()
	fmt.Println("userName", user)

	fmt.Println("the ip is ", IP)

	var isTrue bool

	isTrue = true

	if isTrue {
		fmt.Println("boolean set value")
	}

	//string
	var muitlStr = `a
	b
	c`

	fmt.Println("多行", muitlStr)

	fmt.Println("muitlStr 长度", len(muitlStr))

	var str = "123-456-678-49393"
	splitStr := strings.Split(str, "-")

	fmt.Println("切片 ", str, splitStr)
	//join
	joinStr := strings.Join(splitStr, "*")

	fmt.Println("join ", splitStr, joinStr)

	fmt.Println("contains is ", strings.Contains(joinStr, "134"))

	//string 转换
	age = 20
	agetStr := strconv.FormatInt(int64(age), 10)
	fmt.Println("字符串参数", agetStr)

}

func getUserInfo() (string, int) {
	return "samz", 20
}
