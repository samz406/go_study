package main

import "fmt"

func main() {

	var array1 [3]int
	fmt.Println(array1)
	//声明
	var array2 [3]int
	array2[0] = 1
	array2[1] = 2
	array2[2] = 3
	fmt.Println(array2)

	//声明初始化
	var array3 = [3]string{"java", "go", "js"}
	fmt.Println(array3)

	//声明初始化(...)表示可变长度
	var array4 = [...]string{"java", "go", "js", "python"}
	fmt.Println(array4, "长度为", len(array4))

	//循环range
	for k, v := range array4 {
		fmt.Printf("index is %v, value is %v\n", k, v)
	}

	//循环for
	for i := 0; i < len(array4); i++ {
		fmt.Printf("i is %v, value is %v\n", i, array4[i])

	}

	var array5 = []int{4, 67, 89, 2, 3, 56, 10, 43}
	var max = 0
	var avg = 0

	for i := 0; i < len(array5); i++ {

		if max < array5[i] {
			max = array5[i]
		}
	}
	avg = max / len(array5)

	fmt.Printf("max is %v, avg is %v\n", max, avg)

	//数组是值类型
	var array6 = [3]int{1, 2, 3}
	array7 := array6

	array6[0] = 3
	fmt.Println(array6)

	fmt.Println(array7)

	//切片是引用类型
	var array8 = [3]int{1, 2, 3}
	array9 := array8
	//下标为0的数变化
	array8[0] = 3
	fmt.Println(array8)
	fmt.Println(array9)

	//二维数组

	var array10 = [3][2]string{
		{"A1", "A2"},
		{"B1", "B2"},
		{"C1", "C2"}}

	for i := 0; i < len(array10); i++ {
		var column = array10[i]
		for j := 0; j < len(column); j++ {
			fmt.Println(column[j])
		}
	}

	//切片,不需要指定长度
	var array11 = []int{1, 2, 3}
	fmt.Println(array11)
	//初始化值
	var array12 []int
	fmt.Println(array12 == nil)

	var array13 = []string{"java", "go", "js", "python"}

	for i := 0; i < len(array13); i++ {
		fmt.Printf(array13[i])
	}
	//基于切片定义切片
	array14 := array13[2:]
	fmt.Println(array14)

	//切片
	var slice1 = make([]int, 4, 8)
	fmt.Printf("长度 %d，容量 %d", len(slice1), cap(slice1))

}
