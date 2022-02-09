package main

import "fmt"

func main() {
	var age int // 变量声明
	fmt.Println("未初始化 自动赋0值 my age is", age)
	age = 29 // 赋值
	fmt.Println("my age is", age)
	age = 54 // 赋值
	fmt.Println("my new age is", age)

	var age1 int = 29
	fmt.Println("声明变量并初始化 age1 is",age1)
}
