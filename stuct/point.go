package main

import "fmt"

type User struct {
	name    string
	age     int
	address string
}

type Person struct {
	name    string
	age     int
	address string
}

type Animal struct {
	name string
}

type Dog struct {
	age    string
	Animal //继承
}

func (d Dog) haha() {

	fmt.Printf("%v 在 haha", d.name) //name 父struct 的属性
}

//结构体 方法，打印结构体信息
func (p Person) getInfo() {

	fmt.Printf("姓名 %v 年龄:%v", p.name, p.age)
}

//结构体方法指针类型
func (p *Person) getInfo2() {

	fmt.Printf("姓名 %v 年龄:%v", p.name, p.age)
}

func main() {

	//值类型 ： 改变变量副本值时，不会改变变量本身值 (基本类型、数组、结构体)
	//引用类型： 改变变量副本值时，会改变变量本身值 (切片、map)

	fmt.Println("point struct")
	//point
	var a = 10
	var b = &a
	fmt.Println(b)

	fun2(a)
	fmt.Println(a)
	//修改形参
	fun1(&a)
	fmt.Println(a)

	//引用数据类型需要先分配存储空间，可以用make分配，返回的引用类型
	var userMap = make(map[string]string)
	userMap["userName"] = "samz"
	fmt.Println(userMap)

	var slice = make([]int, 4, 4)
	slice[0] = 1
	fmt.Println(slice)
	//new 分配存储空间，返回的是指针类型
	var d = new(int)
	fmt.Println(d)
	*d = 100
	fmt.Println(*d)

	//struct 可以封装 多种类型
	//type 类型名 struct{
	// 字段名 类型
	// 字段名 类型
	// }

	var user1 User

	user1.age = 20
	user1.address = "chengdu"
	user1.name = "samz"

	fmt.Println(user1)

	var user2 = new(User)
	user2.age = 30
	user2.address = "beijing"
	user2.name = "samz"
	//直接输出为结构体指针
	fmt.Println(*user2)
	fmt.Println(user2)

	var user3 = &User{}
	user3.age = 20
	user3.address = "shanghai"
	user3.name = "sam2"
	fmt.Println(user3)

	var user4 = User{
		age:     20,
		address: "shenzhen",
		name:    "sam3",
	}
	fmt.Println(user4)

	var p1 Person
	p1.name = "haha"
	p1.age = 10
	p1.getInfo()

	p1.getInfo2()

}

func fun1(a *int) {
	*a = 20
}

func fun2(a int) {
	a = 5
}
