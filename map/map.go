package main

import (
	"fmt"
)

func main() {
	fmt.Println("this is map study")
	//make 创建
	var user = make(map[string]string)
	user["name"] = "sanmuzi"
	user["age"] = "20"
	fmt.Println(user)

	var user1 = map[string]string{
		"name": "sanmuzi",
		"age":  "20",
	}
	fmt.Println(user1)
	//遍历
	for k, v := range user1 {
		fmt.Println(k, v)
	}
	//修改
	user1["age"] = "30"

	fmt.Println(user1)

	//删除
	delete(user1, "age")
	fmt.Println(user1)

	//切片值是map
	var map1 = make([]map[string]string, 3, 3)

	fmt.Println(map1)
	//赋值
	if map1[0] == nil {
		map1[0] = make(map[string]string)
		map1[0]["name"] = "sanmuzi"
		map1[0]["age"] = "20"
	}
	fmt.Println(map1)
	for _, v := range map1 {
		fmt.Println(v)
	}
	//map value是切片
	var map2 = make(map[string][]string)
	map2["skill"] = []string{"java", "go", "c#"}
	map2["hobby"] = []string{"sing", "movie", "swimming"}
	fmt.Println(map2)
	for k, v := range map2 {
		fmt.Printf(k)
		fmt.Println(v)
	}

	var map3 = map[int]int{
		1: 1,
		2: 2,
		3: 3,
		4: 4,
		5: 5,
	}

	fmt.Println(map3)

}
