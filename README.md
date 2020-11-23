# go_study
study go

### 基本类型


### 数组、切片

切片的本质就是对底层数组的封装。包含了：底层数组的指针、长度
容量

切片扩容不能通过下边方式赋值，需要用append方法


### Map

Map 是一种无序的键值对的集合。Map 最重要的一点是通过 key 来快速检索数据，key 类似于索引，指向数据的值。

Map 是一种集合，所以我们可以像迭代数组和切片那样迭代它。不过，Map 是无序的，我们无法决定它的返回顺序，这是因为 Map 是使用 hash 表来实现的。



### 函数

函数是组织好的、可重复使用的、用于执行指定任务的代码块。

函数类型有：函数、匿名函数、闭包

函数有的东西

```
func main() {

}

```

它由以下几部分构成：

1. 任何一个函数的定义，都有一个 func 关键字，用于声明一个函数，就像使用 var 关键字声明一个变量一样。
2. 然后紧跟的 main 是函数的名字，命名符合 Go 语言的规范即可，比如不能以数字开头。
3. main 函数名字后面的一对括号 () 是不能省略的，括号里可以定义函数使用的参数，这里的 main 函数没有参数，所以是空括号 () 。
4. 括号 () 后还可以有函数的返回值，因为 main 函数没有返回值，所以这里没有定义。
5. 最后就是大括号 {} 函数体了，你可以在函数体里书写代码，写该函数自己的业务逻辑。



有返回值

```go
func sum(a int,b int) int{
    return a+b
}
```

这是一个计算两数之和的函数，函数的名字是 sum，它有两个参数 a、b，参数的类型都是 int。sum 函数的返回值也是 int 类型，函数体部分就是把 a 和 b 相加，然后通过 return 关键字返回，如果函数没有返回值，可以不用使用 return 关键字



```go
func sum(a, b int) int {
    return a + b
}
```

a 和 b 形参的类型是一样的，这个时候我们可以省略其中一个类型的声明



**多值返回**



```go
func sum(a, b int) (int,error){
    if a<0 || b<0 {
        return 0,errors.New("a或者b不能是负数")
    }
    return a + b,nil
}
```

(int,error)。这代表函数 sum 有两个返回值，第一个是 int 类型，第二个是 error 类型，函数体中使用 return 返回结果的时候，也要符合这个类型顺序



```go
func main() {
    result,err := sum(1, 2)
    if err!=nil {
        fmt.Println(err)
    }else {
        fmt.Println(result)
    }
}
```

```go
result,_ := sum(1, 2)
```

可以使用下划线 _ 丢弃它，这种方式我在 for range 循环那节课里也使用过，这样即可忽略函数 sum 返回的错误信息，也不用再做判断

> 这里使用的 error 是 Go 语言内置的一个接口，用于表示程序的错误信息。





包级函数

同一个包中的函数哪怕是私有的（函数名称首字母小写）也可以被调用。如果不同包的函数要被调用，那么函数的作用域必须是公有的，也就是函数名称的首字母要大写，比如 Println



1. 函数名称首字母小写代表私有函数，只有在同一个包中才可以被调用。
2. 函数名称首字母大写代表公有函数，不同的包也可以调用。
3. 任何一个函数都会从属于一个包。



注： 引用于《**22 讲通关 Go 语言**》



**匿名函数**

匿名函数就是没有名字的函数，这是它和正常函数的主要区别

```go
func main() {
    sum2 := func(a, b int) int {
        return a + b
    }
    fmt.Println(sum2(1, 2))
}
```



有了匿名函数，就可以在函数中再定义函数（函数嵌套），定义的这个匿名函数，也可以称为内部函数。更重要的是，在函数内定义的内部函数，可以使用外部函数的变量等，这种方式称为闭包



**方法**

方法和函数是两个概念，又非常相似，不同点在于方法必须要有一个接收者，这个接收者是一个类型，这样方法就和这个类型绑定在一起，称为这个类型的方法。




**结构体**

结构体是一种聚合类型，里面可以包含任意类型的值，这些值就是我们定义的结构体的成员，也称为字段

****






