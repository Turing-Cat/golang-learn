# Slice 与 Map 详解

## 1. 本节目标

掌握 Slice 和 Map 的区别、初始化方式、数据操作、安全取值、遍历方式。

## 2. 核心概念

### Slice（切片）

- **是什么**：一个可变长的序列，内部包含三个部分：指针（指向底层数组）、长度（len）、容量（cap）。
- **为什么要有它**：数组长度固定不方便，Slice 解决了动态大小的问题。
- **什么时候用**：需要存储一组可变长的同类型数据时。

### Map（映射）

- **是什么**：键值对的无序集合，键唯一。
- **为什么要有它**：通过键快速查找值（O(1) 复杂度）。
- **什么时候用**：需要用键快速定位、统计、去重时。

## 3. 初始化对比

### Slice 初始化

```go
// 1. 字面量
s1 := []int{1, 2, 3}

// 2. make - 指定长度（长度固定，元素为零值）
s2 := make([]int, 3)      // len=3, cap=3, 值全为 0

// 3. make - 指定长度和容量（可追加）
s3 := make([]int, 0, 5)   // len=0, cap=5, 可追加元素

// 4. 切片操作（从已有切片/数组截取）
s4 := s1[0:2]
```

**区别**：
- `make([]int, 3)` → 长度为 3，已有三个零值元素，可以直接通过下标访问 `s2[0] = 10`
- `make([]int, 0, 5)` → 长度为 0，没有任何元素，必须用 `append` 添加

### Map 初始化

```go
// 1. 字面量
m1 := map[string]int{"张三": 98, "李四": 80}

// 2. make - 推荐用法
m2 := make(map[string]int)

// 3. 声明 nil map（不能直接使用！）
var m3 map[string]int  // nil map，不能赋值，会 panic
```

**重要**：声明后必须用 `make()` 初始化才能使用，否则 `m3["key"] = value` 会 panic。

## 4. 数据操作

### Slice 操作

```go
s := []int{1, 2, 3}

// 追加元素
s = append(s, 4)

// 插入元素到中间
s = append(s[:1], append([]int{9}, s[1:]...)...)  // 在位置1插入9

// 删除元素（删除下标为 i 的元素）
i := 1
s = append(s[:i], s[i+1:]...)  // 删除 s[1]

// 修改
s[0] = 100

// 复制
copy(s2, s1)
```

### Map 操作

```go
m := make(map[string]int)

// 增/改
m["张三"] = 98

// 查 - 安全取值
value, ok := m["张三"]  // ok 为 false 表示 key 不存在
if ok {
    fmt.Println(value)
}

// 删（key 不存在不会 panic）
delete(m, "张三")

// 查长度
len(m)
```

## 5. 安全取值

### Map 的二值写法（最重要）

```go
value, ok := m["key"]
// ok == true  → key 存在，value 有效
// ok == false → key 不存在，value 为零值（int 为 0）
```

### Slice 的越界检查

```go
s := []int{1, 2, 3}
if i < len(s) {
    fmt.Println(s[i])
}
```

Go 运行时会检测下标越界，如果越界会 panic，所以访问前先检查长度。

## 6. 遍历方式

### Slice 遍历

```go
// 方式1：index + value
for i, v := range s {
    fmt.Printf("index=%d, value=%d\n", i, v)
}

// 方式2：只 value
for _, v := range s {
    fmt.Println(v)
}

// 方式3：只 index
for i := range s {
    fmt.Println(i)
}
```

### Map 遍历

```go
// 遍历 key-value
for k, v := range m {
    fmt.Printf("key=%s, value=%d\n", k, v)
}

// 只遍历 key
for k := range m {
    fmt.Println(k)
}

// 只遍历 value（不推荐，无法确定是 0 还是不存在）
for _, v := range m {
    fmt.Println(v)
}
```

**注意**：Map 的遍历顺序是随机的，每次运行顺序可能不同。

## 7. 关键区别总结

| 特性 | Slice | Map |
|------|-------|-----|
| 引用类型 | 是 | 是 |
| 初始化方式 | `[]int{}` 或 `make([]int, n)` | `map[string]int{}` 或 `make(map[string]int)` |
| nil 值 | `var s []int` 是 nil | `var m map[string]int` 是 nil |
| nil 时赋值 | 读取 ok 为 false | 写入 panic |
| 长度 | `len(s)` | `len(m)` |
| 容量 | `cap(s)` | 无 cap |
| key 要求 | 无 key，索引是 int | key 必须可比较（不能用 slice、map、func） |
| 并发安全 | 否 | 否（需要 sync 包） |

## 8. 易错点

1. **nil map 不能赋值**：声明 `var m map[string]int` 后必须 `make()` 才能写入。
2. **切片截取是引用**：`s2 := s1[0:2]` 修改 s2 会影响 s1 的底层数组。
3. **make([]int, 0) vs make([]int, 3)`**：前者是空切片，后者是三个零值元素。
4. **Map 遍历无序**：如果需要有序遍历，需要手动排序 key。

## 9. 一句话总结

Slice 是**可生长的序列**（用 `append`），Map 是**键值查找表**（用 `key` 操作），两者都是引用类型。

## 10. 我的疑问

（留空，学习后填写）

## 11. 复习时间

（建议 3 天后复习）