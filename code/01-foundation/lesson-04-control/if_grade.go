package main

import "fmt"

func main0() {
	fmt.Print("请输入你的成绩：")
	var score int
	fmt.Scanln(&score)
	switch {
	case score >= 90:
		fmt.Println("优秀")
	case score >= 80:
		fmt.Println("良好")
	case score >= 60:
		fmt.Println("及格")
	default:
		fmt.Println("不及格")
	}
}
