# L05 - 函数：参数、返回值、多返回值

## 1. 本节目标

- 掌握函数的基本定义语法 `func name(params) (returns) { body }`
- 理解 Go 的参数传递是**值传递**（拷贝，不是传引用）
- 学会使用**多返回值**和**命名返回值**
- 掌握**可变参数** `...` 的用法
- 理解什么时候用 `:=` 什么时候用 `var`

## 2. 核心概念

### 2.1 函数的基本结构

```go
func 函数名(参数列表) (返回值列表) {
    函数体
}
```

**例子：**
```go
func add(a int, b int) int {
    return a + b
}
```

- `func` — 关键字，表示定义函数
- `add` — 函数名
- `a int, b int` — 参数列表（参数名 + 类型）
- `int` — 返回值类型（只有一个时括号可省略）
- `return a + b` — 返回值

### 2.2 值传递（重要！）

Go 的参数传递**永远是值传递**，即拷贝。

```
调用者                   函数内部
   a = 10  ──拷贝──>  副本 a = 10
```

**直白理解：**
- 把变量想象成一张纸条
- 调用函数时，你抄了一份纸条递给函数
- 函数在自己那份上怎么改，跟你原来的纸条没关系

```go
func double(n int) {
    n = n * 2 // 只改了副本
}

func main() {
    x := 5
    double(x)
    fmt.Println(x) // 仍然是 5，不是 10
}
```

**为什么要这样设计？**
- 避免函数意外修改调用者的数据
- 让程序更容易推理（没有副作用）
- C/C++ 的指针、Java 的引用，本质都是绕过这个保护，Go 不默认这么做

### 2.3 多返回值（Go 特色）

Go 函数可以一次返回多个值，用逗号分隔。

```go
func divide(a, b int) (int, int) {
    quotient := a / b
    remainder := a % b
    return quotient, remainder
}

func main() {
    q, r := divide(10, 3) // 同时接收两个返回值
    fmt.Println(q, r)     // 3 1
}
```

**常见用法：错误处理**
```go
func fetch() (string, error) {
    return "", errors.New("出错了")
}

result, err := fetch()
if err != nil {
    fmt.Println("错误:", err)
    return
}
fmt.Println(result)
```

### 2.4 忽略某个返回值

如果不关心某个返回值，用 `_` 占位：
```go
q, _ := divide(10, 3) // 只取商，忽略余数
```

### 2.5 命名返回值

在返回值列表里给变量起名字，函数体内直接用：

```go
func split(sum int) (x, y int) {
    x = sum * 4 / 9
    y = sum - x
    return // 裸 return，返回 x 和 y
}
```

**注意：** 裸 `return` 直接返回命名变量，简洁但不建议在长函数里滥用（可读性差）。

### 2.6 可变参数

参数数量不固定时用 `...`：

```go
func sum(nums ...int) int {
    total := 0
    for _, n := range nums {
        total += n
    }
    return total
}

func main() {
    fmt.Println(sum(1, 2, 3))       // 6
    fmt.Println(sum(1, 2, 3, 4, 5)) // 15
}
```

**类比：** 就像 `fmt.Println` 可以接受任意数量的参数一样。

### 2.7 `:=` 和 `var` 什么时候用

| 场景 | 用什么 | 例子 |
|------|--------|------|
| 函数内部，初始化时同时声明 | `:=` | `x := 10` |
| 函数外部（包级别） | `var` | `var x int = 10` |
| 声明但不想马上赋值 | `var` | `var x int`（零值 0） |
| 声明的同时赋值，且不需在函数内 | `var` | `var x int = 10` |

```go
func main() {
    // 函数内，用 :=
    a := 1
    b := 2

    // 函数内，已声明过的变量，用 =
    a = 10
}
```

**Go 编译器会根据你写的值自动推断类型，所以很少需要 `var x int = 10` 这种写法。**

## 3. 最小示例

```go
package main

import "fmt"

// 多返回值
func divide(a, b int) (int, int) {
    return a / b, a % b
}

// 可变参数
func sum(nums ...int) int {
    total := 0
    for _, n := range nums {
        total += n
    }
    return total
}

func main() {
    q, r := divide(10, 3)
    fmt.Println(q, r)       // 3 1

    fmt.Println(sum(1, 2))  // 3
    fmt.Println(sum(1, 2, 3, 4)) // 10
}
```

## 4. 易错点

1. **`return` 和 `fmt.Println` 混淆**
   - `return` 是把值返回给调用者
   - `fmt.Println` 是把值打印到终端
   - 写了 `return fmt.Println(...)` 会把打印的行数作为返回值返回，完全不是预期行为

2. **函数内修改了参数，以为调用者能看到**
   - Go 是值传递，不会修改调用者的变量
   - 如果真的需要修改，用**指针**（阶段 2 会学）

3. **可变参数只能放在参数列表最后**
   ```go
   // 合法
   func f(a int, nums ...int)

   // 非法
   func f(nums ...int, a int)
   ```

4. **`_` 只能用在接收多返回值时**
   - `_` 是"忽略这个返回值"，不是"省略类型声明"

5. **裸 `return`（Bare return）在长函数里慎用**
   - 只有在短函数里才容易读懂
   - IDE 高亮通常会显示返回值名称，但调试时仍容易混淆

## 5. 一句话总结

**Go 的函数通过值传递复制参数，支持多返回值和可变参数，函数内部永远看不到调用者的原始变量。**

## 6. 我的疑问

（由学习者填写）

-

## 7. 复习时间

- 3 天后：重写本节所有代码示例，不看提示
- 7 天后：用自己的话解释"值传递 vs 引用传递"
- 30 天后：完成阶段 2 指针前，再回顾一次

## 8. 附录：完整函数语法

```go
// 普通函数
func name(params) returnType { }

// 多返回值
func name(params) (ret1, ret2 type) { }

// 可变参数
func name(params ...int) { }

// 完整形式（命名返回值）
func name(params) (ret1 type1, ret2 type2) {
    ret1 = ...
    ret2 = ...
    return
}
```
