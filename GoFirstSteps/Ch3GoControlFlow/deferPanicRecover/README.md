## defer, panic, recover

终于要有点和其他语言不一样的了. 首先我们来逐个了解.

### defer

defer 后跟的表达式都会推迟执行, 会在包含此表达式的函数完成后执行. 若有多个 defer,
则从后向前执行. 这一特点使得 defer 用于关闭文件或运行清理等工作.

例如:

```go
package main

import "fmt"

func main() {
    for i := 1; i <= 4; i++ {
        defer fmt.Println("deferred", -i)
        fmt.Println("regular", i)
    }
}
```

会按顺序打印四个 regular, 再反向打印四个 deferred.

一个典型用例是使用文件后将其关闭:

```go
package main

import (
    "io"
    "os"
)

func main() {
    f, err := os.Create("notes.txt")
    if err != nil {
        return
    }
    defer f.Close()

    if _, err = io.WriteString(f, "Learning Go!"); err != nil {
        return
    }

    f.Sync()
}
```

创建或打开某个文件后, 使用 `defer f.Close()`, 以免向文件中写入内容后忘记关闭该文件.


### panic 与 recover

运行时错误可以让 Go 进入紧急状态, 调用 panic 函数也可以.

调用 `panic()` 函数时可以添加任何值作为参数, 此参数作为进入紧急状态的说明.

进入紧急状态后程序不会继续执行, 但会将所有 `defer` 的表达式后进先出执行完毕.

所以我们可以结合 recover, 在第一个定义最后一个执行的 `defer` 语句中使用 `recover`
重新获得程序的控制权. 可用于避免程序崩溃, 处理资源等.

例如:

```go
package main

import "fmt"

func main() {
    defer func() {
        if r := recover(); r != nil {
            fmt.Println("Recovered in main", r)
        }
    }()
    g(0)
    fmt.Println("Program finished successfully!")
}

func g(i int) {
    if i > 3 {
        fmt.Println("Panicking!")
        panic("Panic in g() (major)")
    }
    defer fmt.Println("Defer in g()", i)
    fmt.Println("Printing in g()", i)
    g(i + 1)
}
```

输出为:

```
Printing in g() 0
Printing in g() 1
Printing in g() 2
Printing in g() 3
Panicking!
Defer in g() 3
Defer in g() 2
Defer in g() 1
Defer in g() 0
Recovered in main Panic in g() (major)
```

若不使用在 defer 表达式中的 recover, 程序将崩溃并打印堆栈错误.


Go 无 `try/catch` 块, 使用 `panic/recover` 处理异常. 这是和其他语言非常大的一个不同点.

