## 函数

函数的定义与其他语言相同. 目前我们一直在使用 `main` 和 `println`.

### main 函数

Go 中所有可执行程序都具有此函数, 因为它是程序的起点, 且只能有一个.
如果创建的是 Go 包, 则可无需编写 `main` 函数.

`main` 函数无任何参数, 不返回任何值.

若要读取命令行参数, 可使用 `os.Args`

若使用自定义函数求和, 则可以:

```go
func sum(number1 string, number2 string) int {
    int1, _ = strconv.Atoi(number1)
    int2, _ = strconv.Atoi(number2)
    return int1 + int2
}
```

或者为返回值设置名称:

```go
func sum(number1 string, number2 string) (result int) {
    int1, _ = strconv.Atoi(number1)
    int2, _ = strconv.Atoi(number2)
    result = int1 + int2
    return
}
```

不过这种方式不清晰, 不建议使用.


### 返回多个值

比如定义一个函数计算和与积:
```go
func calc(number1 string, number2 string) (sum int, mul int) {
    int1, _ = strconv.Atoi(number1)
    int2, _ = strconv.Atoi(number2)
    sum = int1 + int2
    mul = int1 * int2
    return
}
```

则 `main` 为:

```go
func main() {
    sum, mul = calc(os.Args[1], os.Args[2])
    println(sum, mul)
}
```

### 更改函数参数值(使用指针)

Go 作为按值传递的编程语言, 函数中的每个更改都不会影响调用方.

对于以下示例:

```go
package main

func main() {
    firstName := "John"
    updateName(firstName)
    println(firstName)
}

func updateName(name string) {
    name = "David"
}
```

输出依然是 `John`, 因为 Go 传递的是 name 的副本, 一个新的变量.

若要进行原地修改, 则需要使用指针.

> 指针是包含另一个变量的内存地址的变量.

当向函数发送指针作为变量时, 不会传递值, 而是传递地址内存, 因此对变量所做的每个更改都会影响调用方.

Go 同 C 一样, 有两个运算符可以用于处理指针:

- `&` 运算符使用其后对象的地址
- `*` 运算符取消引用指针, 即使用指针所指向的对象.

因此 `updateName` 函数可以这样实现:

```go
func updateName(name *string) {
    *name = "David"
}
```
其中 `*string` 表示类型为指向字符串的指针; `*name` 表示 `name` 指向的字符串对象.

调用则变为 `updateName(&name)`. 其中 `&name` 表示指向 `name` 的字符串指针.

