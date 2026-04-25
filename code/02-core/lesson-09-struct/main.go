package main

import "fmt"

type Book struct {
	Title  string
	Author string
	Price  float64
}

func main() {
	b1 := Book{
		Title:  "《Go语言实战》",
		Author: "张三",
		Price:  99.89,
	}

	fmt.Printf("%s的作者是%s，价格是%.2f", b1.Title, b1.Author, b1.Price)
}
