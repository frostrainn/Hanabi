package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var a int = 89
	b := 95
	var x int64 = 10231234124
	fmt.Printf("type of x is %T, size of x is %d \n",x, unsafe.Sizeof(x))
	fmt.Println("value of a is", a, "and b is", b)
	fmt.Printf("type of a is %T, size of a is %d", a, unsafe.Sizeof(a)) // a 的类型和大小
	fmt.Printf("\ntype of b is %T, size of b is %d", b, unsafe.Sizeof(b)) // b 的类型和大小
}