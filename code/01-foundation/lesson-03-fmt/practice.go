package main

import (
	"fmt"
)

func main3() {
	var name string
	fmt.Print("请输入名字：")
	fmt.Scanln(&name)
	fmt.Print("请输入年龄：")
	var age int
	fmt.Scanln(&age)
	fmt.Printf("你好，%s！你今年%d岁了，明年就%d岁了\n", name, age, age+1)

}
