# L01 Hello World、package、main、import

## 1. 本节目标
理解 Go 程序的基本结构，能独立写出一个可运行的 Hello World 程序。

## 2. 核心概念

- **package main**：声明包名，`main` 是特殊包，表示程序入口，只有 `main` 包才能生成可执行文件
- **import "fmt"**：导入标准库包，Go 不允许导入未使用的包（编译报错）
- **func main()**：程序入口函数，名字必须是 `main`，无参数无返回值
- **fmt.Println**：打印一行，自动加换行（ln = line）
- **fmt.Printf**：格式化打印，不自动换行，需要手动加 `\n`
- **格式化动词**：`%s` 字符串，`%d` 十进制整数

## 3. 最小示例

```go
package main

import "fmt"

func main() {
    fmt.Println("Hello, Go!")
}
```

## 4. 易错点

- 包名不是 `main` → 无法生成可执行文件
- 导入了包但没有使用 → 编译报错
- `Println` 写成小写 `println` → 能跑但不是标准做法
- `Printf` 忘记加 `\n` → 输出不换行

## 5. 一句话总结

一个 Go 程序 = `package main` + `import` + `func main()`，三者缺一不可。

## 6. 我的疑问

（留空，复习时补充自己的疑问）

## 7. 复习时间

学习日期：2026-04-02
建议复习：1 天后、3 天后各看一遍
