## 数据类型

Go 作为强类型语言, 有四类数据类型:

- 基本类型: 数字, 字符串, 布尔值
- 聚合类型: 数组和结构
- 引用类型: 指针, 切片, 映射, 函数, 通道
- 接口类型: 接口


### 整数数字

除 `int` 外, Go 还提供了 `int8`, `int16`, `int32`, `int64` 等.

```go
var integer8 int8 = 127
var integer16 int16 = 32767
var integer32 int32 = 2147483647
var integer64 int64 = 9223372036854775807
println(integer8, integer16, integer32, integer64)

// 以下编译会出错: mismatched types int16 and int32
// println(integer16 + integer32)
```

需要注意的是, 需要强制转换时需要显式进行. 如在不同类型之间执行数学运算时.

另外, `runes` 为 `int32` 的别名, 用于表示 Unicode 字符.


### 浮点数字

浮点数包含 `float32` 和 `float64`

```go
package main

import "math"

func main() {
    var float32 float32 = 2147483647
    var float64 float64 = 9223372036854775807
    println(float32, float64)
    println(math.MaxFloat32, math.MaxFloat64)
}
```


### 布尔型

只可能为 `true`, `false`, 使用 `bool` 声明.

不可隐式转换为 0 或 1, 必须显式操作.


### 字符串

最常用, 使用 string 表示, 使用双引号.
单引号表示单个字符和 runes.

```go
var firstName string = "John"
lastName := "Doe"
println(firstName, lastName)
```

### 默认值

在 Go 中, 如果不进行初始化, 所有数据类型都有默认值, 无需检查是否已经初始化.

```go
var defaultInt int
var defaultFloat32 float32
var defaultFloat64 float64
var defaultBool bool
var defaultString string
println(defaultInt, defaultBool, defaultFloat32, defaultFloat64, defaultString)
```

### 类型转换

前面多次提到类型转换需要显式进行.

一种方式是使用内置函数:

```go
var integer16 int16 = 127
var integer32 int32 = 32767
println(int32(integer16) + integer32)
```

另一种方式是使用 [`strconv`](https://golang.org/pkg/strconv/).

```go
package main

import "strconv"

func main() {
	i, _ := strconv.Atoi("-42")
	s := strconv.Itoa(-42)
	println(i, s)
}
```

其中 `_` 作为变量名称时, 意味着我们不会使用该变量的值. 编译器会忽略未使用的问题.
