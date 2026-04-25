package main

import "fmt"

type Book struct {
	Title  string
	Author string
	Price  float64
}

func (b Book) PrintInfo() {
	fmt.Printf("%s的作者是%s，价格是%.2f\n", b.Title, b.Author, b.Price)
}

func (b *Book) Discount(pct float64) {
	b.Price = b.Price * pct
}
func main() {
	b1 := Book{
		Title:  "《Go语言实战》",
		Author: "张三",
		Price:  99.89,
	}

	b1.PrintInfo()
	b1.Discount(0.9)
	b1.PrintInfo()
}
