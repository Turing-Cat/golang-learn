# 变量、常量、基本类型、零值

## 1. 本节目标

学会在 Go 中声明变量、理解基本类型、掌握零值机制、使用常量。

## 2. 核心概念

### 变量声明的 3 种方式

```go
var name string = "Go"   // 方式1：显式指定类型
var age = 25              // 方式2：类型推导
score := 99.5             // 方式3：短变量声明（最常用，只能在函数内）
```

### 基本类型

| 类型    | 含义   | 零值    |
| ------- | ------ | ------- |
| `int`    | 整数   | `0`     |
| `float64`| 小数   | `0`     |
| `string` | 字符串 | `""`    |
| `bool`   | 布尔   | `false` |

### 零值

变量声明但不赋值时，自动拥有零值。Go 追求"没有未定义行为"，每个变量声明那一刻就有确定的值。

### 常量

```go
const pi = 3.14  // 不可修改
```

### 格式化输出

```go
fmt.Printf("名字: %s, 年龄: %d, 分数: %.2f\n", name, age, score)
// %s 字符串, %d 整数, %f 浮点数, %v 通用
// Printf 不会自动换行，需要手动加 \n
```

## 3. 最小示例

```go
package main

import "fmt"

func main() {
	var name string = "Go学习者"
	var age = 25
	score := 99.5
	isLearning := true

	fmt.Println("姓名:", name)
	fmt.Println("年龄:", age)
	fmt.Println("分数:", score)
	fmt.Println("学习中:", isLearning)

	// 零值
	var n int
	var s string
	var b bool
	fmt.Println("int零值:", n)    // 0
	fmt.Println("string零值:", s) // ""
	fmt.Println("bool零值:", b)   // false
}
```

## 4. 易错点

- **未使用变量 = 编译错误**（不是警告，直接报错）
- `:=` 只能在函数内使用，包级别必须用 `var`
- `:=` 不能对已声明的变量重复使用，修改变量用 `=`
- `Printf` 不自动换行，别忘了 `\n`
- Go 是静态类型，变量类型声明后不能改变（和 Python 不同）

## 5. 一句话总结

Go 变量声明灵活但有严格规则：三种声明方式、零值保安全、未使用必报错、`:=` 最常用。

## 6. 我的疑问

（暂无）

## 7. 复习时间

- 1 天后快速回顾：三种声明方式的区别、零值表
- 3 天后手写：不用看资料，写一个包含所有基本类型的程序
