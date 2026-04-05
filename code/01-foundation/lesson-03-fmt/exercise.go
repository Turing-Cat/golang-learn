package main

import "fmt"

func main() {
	product := "Go语言圣经"
	price := 89.9
	count := 3
	fmt.Printf("%s x %d，单价 %d元， 总价 %.2f元", product, count, price, price*float64(count))
}
