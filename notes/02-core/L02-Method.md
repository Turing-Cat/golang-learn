# 2.2 method 与 receiver

## 1. 本节目标

学会用 receiver 给 struct 定义方法，理解值 receiver 和指针 receiver 的区别。

## 2. 核心概念

- **method**：绑定到某个类型的函数，通过 `实例.方法()` 调用
- **receiver（接收者）**：定义在 func 和方法名之间的部分
  - **值 receiver** `(s Student)`：拿到的是原值的副本，修改不影响原值
  - **指针 receiver** `(s *Student)`：拿到的是原值的地址，修改会影响原值
- **什么时候用指针 receiver**：需要修改原值时

## 3. 最小示例

```go
type Book struct {
    Title  string
    Author string
    Price  float64
}

// 值 receiver，不修改原值
func (b Book) PrintInfo() {
    fmt.Printf("%s 作者：%s 价格：%.2f\n", b.Title, b.Author, b.Price)
}

// 指针 receiver，修改原值
func (b *Book) Discount(pct float64) {
    b.Price = b.Price * pct
}

b := Book{Title: "《Go语言实战》", Price: 99.89}
b.PrintInfo()    // 打印原价
b.Discount(0.9)  // 打9折
b.PrintInfo()    // 打印折后价
```

## 4. 易错点

- 忘记用指针 receiver 来修改原值（用了值 receiver 改了就白改了）
- 普通函数和 method 混淆：method 定义有 receiver，调用用 `实例.方法()`
- receiver 语法容易忘记写括号：`(b Book)` 不是 `b Book`

## 5. 一句话总结

method 通过 receiver 绑定到类型，需要修改原值时必须用指针 receiver。

## 6. 我的疑问

## 7. 复习时间
