## 接口

Go 中的接口是用于表示其他类型的行为的数据类型. 即规定对象应当满足的蓝图或协定.
使用接口可以让代码变得更加灵活, 适应性更强, 因为代码未绑定到特定的实现.

与其他语言不同, Go 的接口是隐式的, 即无需指定实现了哪个接口.

### 声明接口

```go
type Shape interface {
    Perimeter() float64
    Area() float64
}
```

即符合此接口的类型必须同时具有以上定义的两种方法并有对应的返回类型.
但对于实现细节不做任何规定.

### 实现接口

接口的实现是隐式的, 无需显式指定实现了哪个接口. 如新建一个 `Square` 结构:

```go
type Square struct {
    size float64
}

func (s Square) Area() float64 {
    return s.size * s.size
}

func (s Square) Perimeter() float64 {
    return s.size * 4
}
```

在运行时, Go 会知道 Square 实现了 Shape 接口.

```go
func main() {
    var s Shape = Square{3}
    fmt.Printf("%T\n", s)
    fmt.Println("Area: ", s.Area())
    fmt.Println("Perimeter:", s.Perimeter())
}
```


再如, 新建一个结构 `Circle`:

```go
type Circle struct {
    radius float64
}

func (c Circle) Area() float64 {
    return math.Pi * c.radius * c.radius
}

func (c Circle) Perimeter() float64 {
    return 2 * math.Pi * c.radius
}
```

创建一个函数, 打印其类型, 面积和周长:

```go
func printInformation(s Shape) {
    fmt.Printf("%T\n", s)
    fmt.Println("Area: ", s.Area())
    fmt.Println("Perimeter:", s.Perimeter())
    fmt.Println()
}
```

则 main 改为:
```go
func main() {
    var s Shape = Square{3}
    printInformation(s)

    c := Circle{6}
    printInformation(c)
}
```

运行后可以看到正常的输出. 即使用接口可以让代码变得更灵活, 更容易拓展.


### 实现字符串接口

```go
type Stringer interface {
    String() string
}
```

如上定义了一个接口 Stringer, 包含一个 String 方法, 返回 string.

而 `fmt.Printf` 函数使用此接口来输出值. 即可以自定义 `String()` 方法来打印自定义字符串.

```go
package main

import "fmt"

type Person struct {
    Name, Country string
}

func (p Person) String() string {
    return fmt.Sprintf("%v is from %v", p.Name, p.Country)
}
func main() {
    rs := Person{"John Doe", "USA"}
    ab := Person{"Mark Collins", "United Kingdom"}
    fmt.Printf("%s\n%s\n", rs, ab)
}
```

### 扩展现有实现

`io.Copy` 的定义是 `func Copy(dst Writer, src Reader) (written int64, err error)`

即目标实现 Writer 接口, 源实现 Reader 接口即可. 而 Writer 接口定义为:

```go
type Writer interface {
    Write(p []byte) (n int, err error)
}
```

对于以下用例:
```go
package main

import (
    "fmt"
    "io"
    "net/http"
    "os"
)

func main() {
    resp, err := http.Get("https://api.github.com/users/microsoft/repos?page=15&per_page=5")
    if err != nil {
        fmt.Println("Error:", err)
        os.Exit(1)
    }

    io.Copy(os.Stdout, resp.Body)
}
```

我们可以替换 os.Stdout 自定义打印到终端的内容.

为了实现 `Writer` 接口我们可以先创建一个空类型:
`type customWriter struct{}`

然后编写自定义 Writer 函数. 另外还需要另一个结构用于 JSON 格式的响应解析为 Go 对象.

```go
type GitHubResponse []struct {
    FullName string `json:"full_name"`
}

func (w customWriter) Write(p []byte) (n int, err error) {
    var resp GitHubResponse
    json.Unmarshal(p, &resp)
    for _, r := range resp {
        fmt.Println(r.FullName)
    }
    return len(p), nil
}
```

然后将 `main` 中的 `io.Copy(os.Stdout, resp.Body)`
中的 os.Stdout 改为 customWriter 的实例即可.


### 编写自定义服务器 API

接口可用于创建服务器API. 一般使用 `net/http` 程序包中的 `http.Handler` 接口.

```go
package http

type Handler interface {
    ServeHTTP(w ResponseWriter, r *Request)
}

func ListenAndServe(address string, h Handler) error
```

编写以下程序:

```go
package main

import (
    "fmt"
    "log"
    "net/http"
)

type dollars float32

func (d dollars) String() string {
    return fmt.Sprintf("$%.2f", d)
}

type database map[string]dollars

func (db database) ServeHTTP(w http.ResponseWriter, req *http.Request) {
    for item, price := range db {
        fmt.Fprintf(w, "%s: %.2f\n", item, price)
    }
}

func main() {
    db := database{"Go T-Shirt": 25, "Go Jacket": 55}
    log.Fatal(http.ListenAndServe("localhost:8000", db))
}
```

更多用法参见 [Writing Web Applications](https://golang.org/doc/articles/wiki/)