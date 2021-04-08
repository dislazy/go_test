package lesson3

import "fmt"

//可以外部访问
func Test(a int, b int) int {
	return a * b
}

//不可以外部访问
func test() {
	fmt.Println("test")
}

//相当于Java中的{}
func init() {
	fmt.Println("hello ")
}

func init() {
	fmt.Println("world")
}

func init() {
	fmt.Println("bigonelab")
}

//Go的返回参数可以多个
func GetData() (int, int, int) {
	return 100, 200, 100
}

func TestOne(t int) int {
	fmt.Println("TestOne获取到的参数地址", &t)
	return t * 10
}

func TestTwo(t *int) int {
	fmt.Println("TestTwo获取到的参数地址", t)
	return *t * 100
}
