## 在 Go 中使用方法

Go 不完全支持 OOP(面向对象编程)

Go 中的方法是特殊的函数, 区别是必须在函数名称前加入一个额外的参数, 此参数为接收方.

方法可用于对函数进行分组并绑定到自定义类型. 上述操作类似于其他语言中创建类.
以此可以实现嵌入, 重载和封装.

### 声明方法

语法如下:

```go
func (variable type) MethodName(parameters ...) {
    // do something
}
```

但在声明方法前必须先创建结构作为接收方. 如一个计算等边三角形周长的方法:

```go
type triangle struct {
    size int
}

func (t triangle) perimeter() int {
    return t.size * 3
}
```

使用如下:
```go
func main() {
    t := triangle{3}
    fmt.Println(t.perimeter())
}
```

不同的结构可以有相同名称的方法, 其实现可以不同.


### 方法中的指针

当方法需要更新变量或者参数太大避免复制时, 可以使用指针传递变量的地址.

```go
func (t *triangle) doubleSize() {
    t.size *= 2
}
```

使用如下:
```go
func main() {
    t := triangle{3}
    t.doubleSize()
    fmt.Println(t.perimeter())
}
```


### 声明其他类型的方法

方法的接收方可以为任意类型, 不只是自定义类型(如结构).

虽然不能在基本类型上创建方法, 但可以基于基本类型创建自定义类型, 然后创建方法.

```go

import (
    "fmt"
    "strings"
)

type upperstring string

func (s upperstring) Upper() string {
    return strings.ToUpper(string(s))
}

func main() {
    s := upperstring("Learing Go!")
    fmt.Println(s)
    fmt.Println(s.Upper())
}
```

### 嵌入方法

结构体可以进行嵌入, 方法可以沿用

```go
type coloredTriangle struct {
    triangle
    color string

}

t := coloredTriangle{triangle{3}, "blue"}
fmt.Println("Size:", t.size)
fmt.Println("Perimeter", t.perimeter())
```

以上行为看起来像子类继承了基类的方法, 实际上不完全如此.
Go 编译器创建了如下的包装器方法来推广原结构体的方法:
```go
func (t coloredTriangle) perimeter() int {
    return t.triangle.perimeter()
}
```

### 重载方法

如上, 编译器默认对原方法进行了推广, 但我们可以对其进行覆盖, 也就是重载其实现.
设彩色三角形周长为普通三角形的两倍, 则:
```go
func (t coloredTriangle) perimeter() int {
    return t.size * 3 * 2
}
```
若仍需使用原来的方法则可显式调用:

```go
func (t coloredTriangle) perimeter() int {
    return t.triangle.perimeter() * 2
}
```

### 方法中的封装

Go 中无 private 或 public 关键字, 使用大写标识符则可以设为公开.

Go 中的封装仅在程序包之间有效, 只能隐藏来自其他程序包的实现详细信息.

创建一个程序包 `geometry`:

```go
package geometry

type Triangle struct {
    size int
}

func (t *Triangle) doubleSize() {
    t.size *= 2
}

func (t *Triangle) SetSize(size int) {
    t.size = size
}

func (t *Triangle) Perimeter() int {
    t.doubleSize()
    return t.size * 3
}
```

则可以使用上述程序包:

```go
func main() {
    t := geometry.Triangle{}
    t.SetSize(3)
    fmt.Println("Perimeter", t.Perimeter())
}
```

但无法使用 `t.size` 或 `t.doubleSize()` 方法.
