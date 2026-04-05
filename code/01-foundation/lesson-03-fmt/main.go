package main

import "fmt"

func main0() {
	name := "小明"
	age := 25
	score := 92.5

	// 1. Println — 自动加空格和换行
	fmt.Println("姓名:", name, "年龄:", age)

	// 2. Printf — 格式化输出，不自动换行
	fmt.Printf("%s 今年 %d 岁，考了 %.1f 分\n", name, age, score)

	// 3. Sprintf — 不打印，返回格式化后的字符串
	msg := fmt.Sprintf("恭喜 %s！", name)
	fmt.Println(msg)

	// 4. 看类型
	fmt.Printf("name 的类型是 %T\n", name)
}
