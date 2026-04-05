package main

import "fmt"

func main0() {
	//声明变量
	var name string = "Go学习者"
	var age = 18
	score := 90
	isLearning := true

	fmt.Println("姓名：", name)
	fmt.Println("年龄：", age)
	fmt.Println("成绩：", score)
	fmt.Println("学习中：", isLearning)

	//零值
	var n int
	var s string
	var b bool
	fmt.Println("---零值---")
	fmt.Println("int零值：", n)
	fmt.Println("string零值：", s)
	fmt.Println("bool零值：", b)

	//var unused int //未使用的变量
}
