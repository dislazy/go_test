package main

import "fmt"

//2 学习如何创建变量及简单方法

//全局变量
var name string = "hello world"

func main() {
	//使用标准创建类型
	var name string = "hello world"
	//使用变量类型推断创建变量
	var nameOne = "hello world"
	//给对象赋值，但是只适用于方法内，不适用于全局变量
	nameTwo := "hello world"
	//创建一个匿名对象，初始值为空串，直接输出是对应对象的地址;*输出是地址对应的值
	nameThree := new(string)
	//查看效果
	fmt.Println(name, nameOne, nameTwo)
	//对比对象的地址和对应的值
	fmt.Println(nameThree, "----", *nameThree, "----")

	//同时获取对应的数据及地址的值
	var age int = 28
	var ptr = &age // &后面接变量名，表示取出该变量的内存地址
	fmt.Println("age: ", age)
	fmt.Println("ptr: ", ptr)
}

// 使用 new 获取对应字段的地址
func newInt() *int {
	return new(int)
}

// 使用正常返回对应字段的地址
func newIntOne() *int {
	var dummy int
	return &dummy
}
