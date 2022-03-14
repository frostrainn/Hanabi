package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	//fmt.Println(time.Now())
	//s,sep := "",""
	//for _, arg := range os.Args[1:]{
	//	s += sep + arg
	//	sep = " "
	//}
	//fmt.Println(s)
	fmt.Println(strings.Join(os.Args[1:], " "))
	//fmt.Println(time.Now())
}
