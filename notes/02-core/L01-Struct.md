# 2.1 struct 与字段

## 1. 本节目标

理解什么是 struct，为什么要使用 struct，学会定义和使用 struct。

## 2. 核心概念

- **struct**：Go 的自定义复合类型，把多个相关的字段打包在一起
- **字段**：struct 里的变量，每个字段有名字和类型
- **定义语法**：
  ```go
  type Student struct {
      Name  string
      Age   int
      Score float64
  }
  ```
- **实例化方式**：
  - 先声明再赋值：`var s Student; s.Name = "张三"`
  - 字面量初始化（推荐）：`s := Student{Name: "张三", Age: 20}`
  - 按字段顺序（不推荐）：`s := Student{"张三", 20, 95.5}`
- **访问字段**：`s.Name`

## 3. 最小示例

```go
type Book struct {
    Title  string
    Author string
    Price  float64
}

b1 := Book{
    Title:  "《Go语言实战》",
    Author: "张三",
    Price:  99.89,
}

fmt.Println(b1.Title)  // 输出：《Go语言实战》
```

## 4. 易错点

- 初始化时推荐用**命名字段**（`Book{Title: "..."}`），不推荐按顺序
- 字段首字母大写 = 可导出（跨包访问），小写 = 仅包内可见
- struct 是值类型，不是引用类型

## 5. 一句话总结

struct 是 Go 用来组合不同类型数据的基本结构，把相关的变量打包成一个整体。

## 6. 我的疑问

## 7. 复习时间
