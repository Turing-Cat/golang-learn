package main

import "fmt"

func main() {
	var number int
	secret := 42
	for secret != number {
		fmt.Print("请输入一个数字：")
		fmt.Scanln(&number)
		if number > secret {
			fmt.Println("太大了")
		} else if number < secret {
			fmt.Println("太小了")
		} else {
			fmt.Println("恭喜你猜对了")
		}
	}
}
