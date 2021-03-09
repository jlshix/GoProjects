## 变量

### 声明变量

```go

// 使用 `var 变量名1 [, 变量名2, ...] 变量类型` 的格式声明
var firstName string

// 或
var firstName, lastName string
var age int


// 或
var (
    firstName, lastName string
    age int
)
```

### 初始化变量

```go
var (
    firstName string = "John"
    lastName string = "Snow"
    age int = 32
)

// 支持类型推断
var (
    firstName = "John"
    lastName = "Snow"
    age = 32
)

// 或
var (
    firstName, lastName, age = "John", "Snow", 32
)
```

更常见的方式是:

```go
func main() {
    firstName, lastName := "John", "Snow"
    age := 32
    println(firstName, lastName, age)
}
```

定义变量名称后使用 `:=` 赋值. 使用此运算符时要声明的变量必须是新变量.

即: 在函数内使用 `变量名 := 变量值` 的格式声明新的变量并初始化, 可不写 `var` 与变量类型.
这也是最常见的方式.


### 声明常量

使用 `const` 声明, 和 `var` 一样可进行类型推断. 常量一般混合大小写字母或全部大写.

```go
const HTTPStatusOK = 200

// 或
const (
    StatusOK = 0
    StatusConnectionReset = 1
    StatusOtherError = 2
)
```

关于 `iota` 参见 [Go wiki iota](https://github.com/golang/go/wiki/Iota)

### 变量声明但未使用, Go 会在编译时抛出错误