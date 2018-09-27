package main

import (
	"container/list"
	"fmt"
)

func main() {

	// 变量重声明  1.类型必须一致 ，
	//           2.重声明时必须是多个变量，里面呢必须有一个新变量
	//           3.在同一个代码块，使用短变量声明方式
	var a int
	a, b := 1, 2
	fmt.Println(a, b)
	// c := []int{1, 2, 3, 4, 5}
	// list.Element
	d := list.List{}
	// list.Element
	// d.push
	

}
