package main // 声明包名，main 表示这是可执行程序的入口

import "fmt" // 引入 fmt 包，用于格式化输入输出

func main() { // func 定义函数，main 是程序入口，Go 会自动调用它
	fmt.Println("Hello, Go!") // 调用 fmt 包的 Println 函数打印一行文字
}
